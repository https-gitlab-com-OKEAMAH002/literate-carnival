/*
	Frontend server for interacting with the AutoRoller.
*/

package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"go.skia.org/infra/autoroll/go/config"
	"go.skia.org/infra/autoroll/go/config/db"
	"go.skia.org/infra/autoroll/go/manual"
	"go.skia.org/infra/autoroll/go/recent_rolls"
	"go.skia.org/infra/autoroll/go/roller_cleanup"
	"go.skia.org/infra/autoroll/go/rpc"
	"go.skia.org/infra/autoroll/go/status"
	"go.skia.org/infra/autoroll/go/unthrottle"
	"go.skia.org/infra/go/alogin"
	"go.skia.org/infra/go/alogin/proxylogin"
	"go.skia.org/infra/go/auth"
	"go.skia.org/infra/go/common"
	"go.skia.org/infra/go/ds"
	"go.skia.org/infra/go/firestore"
	"go.skia.org/infra/go/gerrit"
	"go.skia.org/infra/go/git"
	"go.skia.org/infra/go/gitiles"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/login"
	"go.skia.org/infra/go/roles"
	"go.skia.org/infra/go/sklog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
)

const (
	// The gerrit OAuth 2 3-legged flow redirect handler.
	gerritOAuth2Redirect = "/gerritRedirect/"
)

var (
	// flags.
	configRefreshInterval = flag.Duration("config_refresh", 5*time.Minute, "How often to refresh the roller configs from the database.")
	configGerritProject   = flag.String("config_gerrit_project", "", "Gerrit project used for editing configs.")
	configRepo            = flag.String("config_repo", "", "Repo URL where configs are stored.")
	configRepoPath        = flag.String("config_repo_path", "", "Path within the config repo where configs are stored.")
	firestoreInstance     = flag.String("firestore_instance", "", "Firestore instance to use, eg. \"production\"")
	host                  = flag.String("host", "localhost", "HTTP service host")
	internal              = flag.Bool("internal", false, "If true, display the internal rollers.")
	local                 = flag.Bool("local", false, "Running locally if true. As opposed to in production.")
	port                  = flag.String("port", ":8000", "HTTP service port (e.g., ':8000')")
	project               = flag.String("project", "", "GCP project that this service runs in.")
	promPort              = flag.String("prom_port", ":20000", "Metrics service address (e.g., ':10110')")
	resourcesDir          = flag.String("resources_dir", "", "The directory to find templates, JS, and CSS files. If blank the current directory will be used.")
	hang                  = flag.Bool("hang", false, "If true, don't spin up the server, just hang without doing anything.")

	allowedViewers = []string{
		"prober@skia-public.iam.gserviceaccount.com",
		"skia-status@skia-public.iam.gserviceaccount.com",
		"skia-status-internal@skia-corp.google.com.iam.gserviceaccount.com",
		"status@skia-buildbots.google.com.iam.gserviceaccount.com",
		"status-internal@skia-buildbots.google.com.iam.gserviceaccount.com",
		"showy-dashboards@prod.google.com",
	}

	mainTemplate            *template.Template = nil
	rollerTemplate          *template.Template = nil
	configTemplate          *template.Template = nil
	modeHistoryTemplate     *template.Template = nil
	rollHistoryTemplate     *template.Template = nil
	strategyHistoryTemplate *template.Template = nil

	srv *rpc.AutoRollServer

	configEditsInProgress               = map[string]*config.Config{}
	configGitiles         *gitiles.Repo = nil

	// gerritOauthConfig is the OAuth 2.0 client configuration used for
	// interacting with Gerrit.
	gerritOauthConfig = &oauth2.Config{
		ClientID:     "not-a-valid-client-id",
		ClientSecret: "not-a-valid-client-secret",
		Scopes:       []string{gerrit.AuthScope},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8000/oauth2callback/",
	}
)

func reloadTemplates() {
	if *resourcesDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			sklog.Fatal(err)
		}
		*resourcesDir = filepath.Join(wd, "dist")
	}
	sklog.Infof("Reading resources from %s", *resourcesDir)
	mainTemplate = template.Must(template.New("index.html").Funcs(map[string]interface{}{
		"marshal": func(data interface{}) template.JS {
			b, _ := json.Marshal(data)
			return template.JS(b)
		},
	}).ParseFiles(
		filepath.Join(*resourcesDir, "index.html"),
	))
	rollerTemplate = template.Must(template.ParseFiles(
		filepath.Join(*resourcesDir, "roller.html"),
	))
	configTemplate = template.Must(template.ParseFiles(
		filepath.Join(*resourcesDir, "config.html"),
	))
	modeHistoryTemplate = template.Must(template.ParseFiles(
		filepath.Join(*resourcesDir, "mode-history.html"),
	))
	rollHistoryTemplate = template.Must(template.ParseFiles(
		filepath.Join(*resourcesDir, "roll-history.html"),
	))
	strategyHistoryTemplate = template.Must(template.ParseFiles(
		filepath.Join(*resourcesDir, "strategy-history.html"),
	))
}

func getRoller(w http.ResponseWriter, r *http.Request) *config.Config {
	name := chi.URLParam(r, "roller")
	if name == "" {
		http.Error(w, "Unable to find roller name in request path.", http.StatusBadRequest)
		return nil
	}
	roller, err := srv.GetRoller(name)
	if err != nil {
		http.Error(w, "No such roller", http.StatusNotFound)
		return nil
	}
	return roller.Cfg
}

func rollerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	cfg := getRoller(w, r)
	if cfg == nil {
		return // Errors are handled by getRoller.
	}
	page := struct {
		ChildName  string
		ParentName string
		Roller     string
		Project    string
	}{
		ChildName:  cfg.ChildDisplayName,
		ParentName: cfg.ParentDisplayName,
		Roller:     cfg.RollerName,
		Project:    *project,
	}
	if err := rollerTemplate.Execute(w, page); err != nil {
		httputils.ReportError(w, err, "Failed to expand template.", http.StatusInternalServerError)
	}
}

func modeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	cfg := getRoller(w, r)
	if cfg == nil {
		return // Errors are handled by getRoller.
	}
	page := struct {
		ChildName  string
		ParentName string
		Roller     string
	}{
		ChildName:  cfg.ChildDisplayName,
		ParentName: cfg.ParentDisplayName,
		Roller:     cfg.RollerName,
	}
	if err := modeHistoryTemplate.Execute(w, page); err != nil {
		httputils.ReportError(w, err, "Failed to expand template.", http.StatusInternalServerError)
	}
}

func rollHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	cfg := getRoller(w, r)
	if cfg == nil {
		return // Errors are handled by getRoller.
	}
	page := struct {
		ChildName  string
		ParentName string
		Roller     string
	}{
		ChildName:  cfg.ChildDisplayName,
		ParentName: cfg.ParentDisplayName,
		Roller:     cfg.RollerName,
	}
	if err := rollHistoryTemplate.Execute(w, page); err != nil {
		httputils.ReportError(w, err, "Failed to expand template.", http.StatusInternalServerError)
	}
}

func strategyHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	cfg := getRoller(w, r)
	if cfg == nil {
		return // Errors are handled by getRoller.
	}
	page := struct {
		ChildName  string
		ParentName string
		Roller     string
	}{
		ChildName:  cfg.ChildDisplayName,
		ParentName: cfg.ParentDisplayName,
		Roller:     cfg.RollerName,
	}
	if err := strategyHistoryTemplate.Execute(w, page); err != nil {
		httputils.ReportError(w, err, "Failed to expand template.", http.StatusInternalServerError)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := mainTemplate.Execute(w, nil); err != nil {
		httputils.ReportError(w, errors.New("Failed to expand template."), fmt.Sprintf("Failed to expand template: %s", err), http.StatusInternalServerError)
	}
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse and validate the config.
		configJson := r.FormValue("configJson")
		var cfg config.Config
		if err := protojson.Unmarshal([]byte(configJson), &cfg); err != nil {
			httputils.ReportError(w, err, "Failed to parse config as JSON", http.StatusBadRequest)
			return
		}
		if err := cfg.Validate(); err != nil {
			httputils.ReportError(w, err, err.Error(), http.StatusBadRequest)
			return
		}

		// We're going to redirect for the OAuth2 flow. Store the config in
		// memory.
		// TODO(borenet): What happens if we scale Kubernetes up to multiple
		// frontend pods and the user redirects back to a different instance?
		var sessionID string
		for {
			sessionID = uuid.New().String()
			if _, ok := configEditsInProgress[sessionID]; !ok {
				break
			}
		}
		configEditsInProgress[sessionID] = &cfg
		time.AfterFunc(time.Hour, func() {
			delete(configEditsInProgress, sessionID)
		})

		// Redirect for OAuth2.
		opts := []oauth2.AuthCodeOption{oauth2.AccessTypeOnline, oauth2.SetAuthURLParam("approval_prompt", "auto")}
		redirectURL := gerritOauthConfig.AuthCodeURL(sessionID, opts...)
		http.Redirect(w, r, redirectURL, http.StatusFound)
	} else {
		w.Header().Set("Content-Type", "text/html")
		if err := configTemplate.Execute(w, nil); err != nil {
			httputils.ReportError(w, errors.New("Failed to expand template."), fmt.Sprintf("Failed to expand template: %s", err), http.StatusInternalServerError)
		}
	}
}

func configJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cfg := getRoller(w, r)
	if cfg == nil {
		return // Errors are handled by getRoller.
	}

	b, err := protojson.Marshal(cfg)
	if err != nil {
		httputils.ReportError(w, err, "Failed to encode response.", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		httputils.ReportError(w, err, "Failed to write response.", http.StatusInternalServerError)
		return
	}
}

func submitConfigUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sessionID := r.FormValue("state")
	cfg, ok := configEditsInProgress[sessionID]
	if !ok {
		msg := "Unable to find config"
		httputils.ReportError(w, errors.New(msg), msg, http.StatusBadRequest)
		return
	}
	content, err := prototext.MarshalOptions{
		Indent: "  ",
	}.Marshal(cfg)
	if err != nil {
		httputils.ReportError(w, err, "Failed to encode config to proto.", http.StatusInternalServerError)
		return
	}
	code := r.FormValue("code")
	token, err := gerritOauthConfig.Exchange(ctx, code)
	if err != nil {
		httputils.ReportError(w, err, "Failed to authenticate.", http.StatusInternalServerError)
		return
	}
	ts := gerritOauthConfig.TokenSource(ctx, token)
	client := httputils.DefaultClientConfig().WithTokenSource(ts).Client()
	g, err := gerrit.NewGerrit(gerrit.GerritSkiaURL, client)
	if err != nil {
		httputils.ReportError(w, err, "Failed to initialize Gerrit API.", http.StatusInternalServerError)
		return
	}
	baseCommit, err := configGitiles.ResolveRef(ctx, git.MainBranch)
	if err != nil {
		httputils.ReportError(w, err, "Failed to find base commit.", http.StatusInternalServerError)
		return
	}
	configFile := cfg.RollerName + ".cfg"
	if *configRepoPath != "" {
		configFile = path.Join(*configRepoPath, configFile)
	}
	// TODO(borenet): Handle custom commit messages.
	ci, err := gerrit.CreateAndEditChange(ctx, g, *configGerritProject, git.MainBranch, "Update AutoRoller Config", baseCommit, "", func(ctx context.Context, g gerrit.GerritInterface, ci *gerrit.ChangeInfo) error {
		return g.EditFile(ctx, ci, configFile, string(content))
	})
	if err != nil {
		httputils.ReportError(w, err, "Failed to create change.", http.StatusInternalServerError)
		return
	}
	redirectURL := g.Url(ci.Issue)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// addCorsMiddleware wraps the specified HTTP handler with a handler that applies the
// CORS specification on the request, and adds relevant CORS headers as necessary.
// This is needed for some handlers that do not have this middleware. Eg: the twirp
// handler (https://github.com/twitchtv/twirp/issues/210).
func addCorsMiddleware(handler http.Handler) http.Handler {
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		Debug:          true,
	})
	return corsWrapper.Handler(handler)
}

func runServer(ctx context.Context, serverURL string, srv http.Handler, plogin alogin.Login) {
	r := chi.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.Handle("/dist/*", http.StripPrefix("/dist/", http.HandlerFunc(httputils.MakeResourceHandler(*resourcesDir))))
	r.HandleFunc("/config", configHandler)
	r.HandleFunc(gerritOAuth2Redirect, submitConfigUpdate)
	r.HandleFunc("/_/login/status", alogin.LoginStatusHandler(plogin))
	r.Route("/r/{roller}", func(r chi.Router) {
		r.HandleFunc("/", rollerHandler)
		r.HandleFunc("/config", configJSONHandler)
		r.HandleFunc("/mode-history", modeHistoryHandler)
		r.HandleFunc("/roll-history", rollHistoryHandler)
		r.HandleFunc("/strategy-history", strategyHistoryHandler)
	})
	r.Handle(rpc.AutoRollServicePathPrefix+"*", addCorsMiddleware(srv))
	h := httputils.LoggingRequestResponse(r)
	h = httputils.XFrameOptionsDeny(h)
	if !*local {
		if *internal {
			h = alogin.ForceRole(h, plogin, roles.Viewer)
		}
		h = httputils.HealthzAndHTTPS(h)
	}
	http.Handle("/", h)
	sklog.Infof("Ready to serve on %s", serverURL)
	sklog.Fatal(http.ListenAndServe(*port, nil))
}

func main() {
	common.InitWithMust(
		"autoroll-fe",
		common.PrometheusOpt(promPort),
	)
	defer common.Defer()

	reloadTemplates()

	if *hang {
		select {}
	}

	ctx := context.Background()

	ts, err := google.DefaultTokenSource(ctx, auth.ScopeUserinfoEmail, auth.ScopeGerrit, datastore.ScopeDatastore)
	if err != nil {
		sklog.Fatal(err)
	}
	namespace := ds.AUTOROLL_NS
	if *internal {
		namespace = ds.AUTOROLL_INTERNAL_NS
	}
	if err := ds.InitWithOpt(common.PROJECT_ID, namespace, option.WithTokenSource(ts)); err != nil {
		sklog.Fatal(err)
	}

	statusDB, err := status.NewDB(ctx, firestore.FIRESTORE_PROJECT, namespace, *firestoreInstance, ts)
	if err != nil {
		sklog.Fatalf("Failed to create status DB: %s", err)
	}
	configDB, err := db.NewDBWithParams(ctx, firestore.FIRESTORE_PROJECT, namespace, *firestoreInstance, ts)
	if err != nil {
		sklog.Fatal(err)
	}
	rollsDB := recent_rolls.NewDatastoreRollsDB(ctx)
	manualRollDB, err := manual.NewDBWithParams(ctx, firestore.FIRESTORE_PROJECT, *firestoreInstance, ts)
	if err != nil {
		sklog.Fatal(err)
	}
	cleanupDB, err := roller_cleanup.NewDBWithParams(ctx, firestore.FIRESTORE_PROJECT, *firestoreInstance, ts)
	if err != nil {
		sklog.Fatal(err)
	}
	throttleDB := unthrottle.NewDatastore(ctx)

	if *configRepo == "" {
		sklog.Fatal("--config_repo is required.")
	}
	if *configGerritProject == "" {
		sklog.Fatal("--config_gerrit_project is required.")
	}
	client := httputils.DefaultClientConfig().WithTokenSource(ts).Client()
	configGitiles = gitiles.NewRepo(*configRepo, client)

	plogin := proxylogin.NewWithDefaults()
	srv, err = rpc.NewAutoRollServer(ctx, statusDB, configDB, rollsDB, manualRollDB, cleanupDB, throttleDB, *configRefreshInterval, plogin)
	if err != nil {
		sklog.Fatal(err)
	}

	serverURL := "https://" + *host
	if *local {
		serverURL = "http://" + *host + *port
	}

	// Load the OAuth2 config information.

	// TODO(jcgregorio) This should not use
	// the client id and secret used for skia.org login, instead a separate
	// client id and secret should be used.
	_, clientID, clientSecret, err := login.TryLoadingFromAllSources(ctx)
	if err != nil {
		sklog.Fatalf("Failed to load OAuth2 configuration: %s", err)
	}
	gerritOauthConfig.ClientID = clientID
	gerritOauthConfig.ClientSecret = clientSecret
	gerritOauthConfig.RedirectURL = serverURL + gerritOAuth2Redirect

	// Create the server.
	runServer(ctx, serverURL, srv.GetHandler(), plogin)
}
