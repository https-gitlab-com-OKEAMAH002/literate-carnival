package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"cloud.google.com/go/datastore"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"go.skia.org/infra/autoroll/go/status"
	"go.skia.org/infra/go/alogin"
	"go.skia.org/infra/go/alogin/proxylogin"
	"go.skia.org/infra/go/baseapp"
	"go.skia.org/infra/go/common"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/metrics2"
	"go.skia.org/infra/go/roles"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/go/util"
)

// defaultSkiaRepo is the repo to default to when no repos have been specified.
// It is the main "skia.git" repo.
const defaultSkiaRepo = "skia"

// Flags
var (
	host               = flag.String("host", "tree-status.skia.org", "HTTP service host")
	modifyGroup        = flag.String("modify_group", "project-skia-committers", "The chrome infra auth group to use for who is allowed to change tree status.")
	chromeInfraAuthJWT = flag.String("chrome_infra_auth_jwt", "/var/secrets/skia-public-auth/key.json", "Path to a local file, or name of a GCP secret, containing the JWT key for the service account that has access to chrome infra auth.")
	namespace          = flag.String("namespace", "tree-status-staging", "The Cloud Datastore namespace.")
	dsProject          = flag.String("ds-project", "skia-public", "Name of the GCP project used for Datastore.")
	repos              = common.NewMultiStringFlag("repo", nil, "These repos will have tree status endpoints.")
	secretProject      = flag.String("secret-project", "skia-infra-public", "Name of the GCP project used for secret management.")
	internalPort       = flag.String("internal_port", "", "HTTP internal service address (eg: ':8001' for unauthenticated in-cluster requests.")
	firestoreInstance  = flag.String("firestore_instance", "production", "Firestore instance to use, eg. \"production\"")
)

var (
	// dsClient is the Cloud Datastore client to access tree statuses.
	dsClient *datastore.Client

	// repoNameRegex matches the format of supported repo names.
	repoNameRegex = "{repo:[0-9a-zA-Z._-]+}"
)

// Server is the state of the server.
type Server struct {
	templates  *template.Template
	autorollDB status.DB

	// skiaRepoSpecified is set to true when the main skia has been specified.
	// This boolean is used because the main skia repo requires support for
	// non-repo specified URLs (for backwards compatibility) and for watching
	// autorollers.
	skiaRepoSpecified bool

	plogin alogin.Login
}

// New implements baseapp.Constructor.
func New() (baseapp.App, error) {
	ctx := context.Background()
	ts, err := google.DefaultTokenSource(ctx, "https://www.googleapis.com/auth/datastore")
	if err != nil {
		return nil, skerr.Wrapf(err, "Problem setting up default token source")
	}

	dsClient, err = datastore.NewClient(context.Background(), *dsProject, option.WithTokenSource(ts))
	if err != nil {
		return nil, skerr.Wrapf(err, "Failed to initialize Cloud Datastore for tree status")
	}

	// Check to see if the main skia repo has been specified. If it has been
	// specified then it will require special handling.
	skiaRepoSpecified := IsRepoSupported(defaultSkiaRepo)

	var autorollDB status.DB
	if skiaRepoSpecified {
		// Start watching for statuses with autorollers specified. Only supported for
		// the default repo (skia).
		autorollDB, err = AutorollersInit(ctx, defaultSkiaRepo, *firestoreInstance, ts)
		if err != nil {
			return nil, skerr.Wrapf(err, "Could not init autorollers")
		}

		// Load the last status and whether autorollers need to be watched.
		s, err := GetLatestStatus(defaultSkiaRepo)
		if err != nil {
			return nil, skerr.Wrapf(err, "Could not find latest status")
		}
		if s.Rollers != "" {
			sklog.Infof("Last status has rollers that need to be watched: %s", s.Rollers)
			StartWatchingAutorollers(s.Rollers)
		}
	}

	plogin := proxylogin.NewWithDefaults()

	srv := &Server{
		autorollDB:        autorollDB,
		skiaRepoSpecified: skiaRepoSpecified,
		plogin:            plogin,
	}
	srv.loadTemplates()
	liveness := metrics2.NewLiveness("alive", map[string]string{})
	fmt.Println(liveness)

	return srv, nil
}

func (srv *Server) loadTemplates() {
	blah := *baseapp.ResourcesDir
	srv.templates = template.Must(template.New("").Delims("{%", "%}").ParseFiles(
		filepath.Join(blah, "index.html"),
	))
}

// IsEditor returns true if the currently logged in user is an editor.
func (srv *Server) IsEditor(r *http.Request) bool {
	ret := true
	if !*baseapp.Local {
		ret = srv.plogin.Roles(r).Has(roles.Editor)
	}
	return ret
}

// AddHandlers implements baseapp.App.
func (srv *Server) AddHandlers(r chi.Router) {
	r.HandleFunc("/_/login/status", alogin.LoginStatusHandler(srv.plogin))

	// All endpoints that require authentication should be added to this router. The
	// rest of endpoints are left unauthenticated because they are accessed from various
	// places like: Skia infra apps, Gerrit plugin, Chrome extensions, presubmits, etc.
	appRouter := chi.NewRouter()

	if srv.skiaRepoSpecified {
		// If the main skia repo has been specified then leave default repo
		// handlers around for backwards compatibility.
		appRouter.Get("/", srv.treeStateDefaultRepoHandler)
		r.Get("/current", httputils.CorsHandler(srv.bannerStatusHandler))
	}
	appRouter.Post("/_/get_autorollers", srv.autorollersHandler)

	// Add repo-specific endpoints.
	appRouter.Get(fmt.Sprintf("/%s", repoNameRegex), srv.treeStateDefaultRepoHandler)
	appRouter.Post(fmt.Sprintf("/%s/_/add_tree_status", repoNameRegex), srv.addStatusHandler)
	appRouter.Post(fmt.Sprintf("/%s/_/recent_statuses", repoNameRegex), srv.recentStatusesHandler)
	r.Get(fmt.Sprintf("/%s/current", repoNameRegex), httputils.CorsHandler(srv.bannerStatusHandler))

	if *internalPort != "" {
		internalRouter := chi.NewRouter()
		internalRouter.Get(fmt.Sprintf("/%s/current", repoNameRegex), httputils.CorsHandler(srv.bannerStatusHandler))
		internalRouter.Get("/current", srv.bannerStatusHandler)

		go func() {
			sklog.Infof("Internal server on %q", *internalPort)
			sklog.Fatal(http.ListenAndServe(*internalPort, internalRouter))
		}()
	}

	// Use the appRouter as a handler and wrap it into middleware that enforces authentication.
	appHandler := http.Handler(appRouter)

	r.Handle("/*", appHandler)
}

// AddMiddleware implements baseapp.App.
func (srv *Server) AddMiddleware() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{}
}

// IsRepoSupported is a utility function that returns true if the specified
// repo is a supported repo (i.e. has been specified in the repos flag).
func IsRepoSupported(repo string) bool {
	return util.In(repo, *repos)
}

func main() {
	// Parse flags to be able to send *host to baseapp.Serve
	flag.Parse()
	baseapp.Serve(New, []string{*host})
}
