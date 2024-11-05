package authproxy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/allowed"
	"go.skia.org/infra/go/cleanup"
	"go.skia.org/infra/go/mockhttpclient"
	"go.skia.org/infra/go/roles"
	"go.skia.org/infra/kube/go/authproxy/auth/mocks"
)

const (
	viewerEmail     = "nobody@example.org"
	notAViewerEmail = "notallowed@example.org"
)

var (
	errForTesting = errors.New("my test error")
)

var commonAllowed = map[roles.Role]allowed.Allow{
	roles.Viewer: allowed.NewAllowedFromList([]string{viewerEmail}),
}

func assertValidEmailAndRole(t *testing.T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{viewerEmail}, r.Header.Values(WebAuthHeaderName))
		require.Equal(t, []string{string(roles.Viewer)}, r.Header.Values(WebAuthRoleHeaderName))
	}
}

func setupForTest(t *testing.T, cb http.HandlerFunc) (*url.URL, *bool, *httptest.ResponseRecorder, *http.Request) {
	called := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cb(w, r)
		called = true
	}))
	t.Cleanup(func() {
		ts.Close()
	})
	u, err := url.Parse(ts.URL)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", ts.URL, nil)
	return u, &called, w, r
}

func TestProxyServeHTTP_AllowPostAndNotAuthenticated_WebAuthHeaderValueIsEmptyString(t *testing.T) {
	u, called, w, r := setupForTest(t, func(w http.ResponseWriter, r *http.Request) {
		// Note that if the header webAuthHeaderName hadn't been set then the value would be nil.
		require.Equal(t, []string{""}, r.Header.Values(WebAuthHeaderName))
		require.Equal(t, []string(nil), r.Header.Values("X-SOME-UNSET-HEADER"))
	})
	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return("", errForTesting)

	proxy := newProxy(u, authMock, true, false, false, false, true)
	proxy.allowedRoles = commonAllowed

	proxy.ServeHTTP(w, r)
	require.True(t, *called)
}

func TestProxyServeHTTP_UserIsLoggedIn_HeaderWithUserEmailIsIncludedInRequest(t *testing.T) {
	u, called, w, r := setupForTest(t, assertValidEmailAndRole(t))

	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return(viewerEmail, nil)

	proxy := newProxy(u, authMock, false, false, false, false, true)
	proxy.allowedRoles = commonAllowed

	proxy.ServeHTTP(w, r)
	require.True(t, *called)
}

func TestProxyServeHTTP_UserIsLoggedInAndBelongsToTwoRoles_HeaderWithBothRolesIsIncludedInRequest(t *testing.T) {
	u, called, w, r := setupForTest(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{viewerEmail}, r.Header.Values(WebAuthHeaderName))
		receivedRoles := strings.Split(r.Header.Get(WebAuthRoleHeaderName), ",")
		sort.Strings(receivedRoles)
		expectedRoles := []string{
			string(roles.Editor),
			string(roles.Viewer),
		}
		require.Equal(t, expectedRoles, receivedRoles)
	})

	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return(viewerEmail, nil)

	allowedRoles := map[roles.Role]allowed.Allow{
		roles.Viewer: allowed.NewAllowedFromList([]string{viewerEmail}),
		roles.Editor: allowed.NewAllowedFromList([]string{viewerEmail}),
	}
	proxy := newProxy(u, authMock, false, false, false, false, true)
	proxy.allowedRoles = allowedRoles

	proxy.ServeHTTP(w, r)
	require.True(t, *called)
}

func TestProxyServeHTTP_UserIsNotLoggedIn_HeaderWithUserEmailIsStrippedFromRequest(t *testing.T) {
	u, called, w, r := setupForTest(t, func(w http.ResponseWriter, r *http.Request) {})
	r.Header.Add(WebAuthHeaderName, viewerEmail) // Try to spoof the header.
	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return("", errForTesting)
	authMock.On("LoginURL", w, r).Return("http://example.org/login")

	proxy := newProxy(u, authMock, false, false, false, false, true)
	proxy.allowedRoles = commonAllowed

	proxy.ServeHTTP(w, r)
	require.False(t, *called)
}

func TestProxyServeHTTP_UserIsLoggedInButNotAViewer_ReturnsStatusForbidden(t *testing.T) {
	u, called, w, r := setupForTest(t, func(w http.ResponseWriter, r *http.Request) {})
	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return(notAViewerEmail, nil)

	proxy := newProxy(u, authMock, false, false, false, false, true)
	proxy.allowedRoles = commonAllowed

	proxy.ServeHTTP(w, r)
	require.False(t, *called)
	require.Equal(t, http.StatusForbidden, w.Result().StatusCode)
}

func TestProxyServeHTTP_UserIsLoggedIn_HeaderWithUserEmailIsIncludedInRequestAndSpoofedEmailIsRemoved(t *testing.T) {
	u, called, w, r := setupForTest(t, assertValidEmailAndRole(t))
	r.Header.Add(WebAuthHeaderName, "haxor@example.org") // Try to spoof the header.
	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return(viewerEmail, nil)

	proxy := newProxy(u, authMock, false, false, false, false, true)
	proxy.allowedRoles = commonAllowed

	proxy.ServeHTTP(w, r)
	require.True(t, *called)
}

func TestProxyServeHTTP_UserIsNotLoggedInAndPassiveFlagIsSet_RequestIsPassedAlongWithoutEmailHeader(t *testing.T) {
	u, called, w, r := setupForTest(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{""}, r.Header.Values(WebAuthHeaderName))
		require.Equal(t, []string{""}, r.Header.Values(WebAuthRoleHeaderName))
	})

	r.Header.Add(WebAuthHeaderName, "haxor@example.org") // Try to spoof the header.
	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return("", errForTesting)

	proxy := newProxy(u, authMock, false, true, false, false, true)
	proxy.allowedRoles = commonAllowed

	proxy.ServeHTTP(w, r)
	require.True(t, *called)
}

func TestProxyServeHTTP_UserIsLoggedInAndPassiveFlagIsSet_RequestIsPassedAlongWithEmailHeader(t *testing.T) {
	u, called, w, r := setupForTest(t, assertValidEmailAndRole(t))

	r.Header.Add(WebAuthHeaderName, "haxor@example.org") // Try to spoof the header.
	authMock := mocks.NewAuth(t)
	authMock.On("LoggedInAs", r).Return(viewerEmail, nil)

	proxy := newProxy(u, authMock, false, true, false, false, true)
	proxy.allowedRoles = commonAllowed

	proxy.ServeHTTP(w, r)
	require.True(t, *called)
}

func TestValidateFlags_NoRoleFlagsSpecified_ReturnsError(t *testing.T) {
	app := &App{
		roleFlags: nil,
	}

	require.Error(t, app.validateFlags())
}

func TestValidateFlags_MockedUserWithoutAuthTypeMockedFlagsSpecified_ReturnsError(t *testing.T) {
	app := &App{
		mockLoggedInAs: "me@tehgoog.com",
		roleFlags:      []string{"@tehgoog.com"},
	}
	require.Error(t, app.validateFlags())

	app = &App{
		mockLoggedInAs: "me@tehgoog.com",
		authType:       string(OAuth2),
		roleFlags:      []string{"@tehgoog.com"},
	}
	require.Error(t, app.validateFlags())

	app = &App{
		mockLoggedInAs: "me@tehgoog.com",
		authType:       string(Mocked),
		roleFlags:      []string{"@tehgoog.com"},
	}
	require.NoError(t, app.validateFlags())
}

func TestAppRun_ContextIsCancelled_ReturnsNil(t *testing.T) {
	// Construct minimal App.
	target, err := url.Parse("http://my-service")
	require.NoError(t, err)
	app := &App{
		target:   target,
		port:     ":0",
		promPort: ":0",
	}
	app.registerCleanup()

	var w sync.WaitGroup
	w.Add(1)
	go func() {
		err := app.Run(context.Background())
		assert.NoError(t, err)
		w.Done()
	}()

	// Ensure the server has been started.
	for app.server == nil {
		time.Sleep(time.Millisecond)
	}

	// Force a cleanup.
	cleanup.Cleanup()
	w.Wait()

	// Test will fail by timeout if the app.Run() didn't return.
}

const testCriaGroupName = "mytestgroup"

const mockCriaResponse = `{
	"group": {
	  "members": [
		"user:test@example.org",
		"user:*@chromium.org"
	  ],
	  "globs": [
		"user:*@gotham.org"
	  ]
	}
  }`

func mockCriaClient(t *testing.T) *http.Client {
	m := mockhttpclient.NewURLMock()
	m.Mock(fmt.Sprintf(allowed.GROUP_URL_TEMPLATE, testCriaGroupName), mockhttpclient.MockGetDialogue([]byte(mockCriaResponse)))
	return m.Client()
}

func TestAppPopulateAllowedRoles_MultipleGroups_Success(t *testing.T) {
	a := newEmptyApp()
	a.criaClient = mockCriaClient(t)
	a.roleFlags = []string{
		"editor=cria_group:" + testCriaGroupName,
		"admin=google.com",
	}

	err := a.populateAllowedRoles()
	require.NoError(t, err)
	require.True(t, a.proxy.allowedRoles[roles.Editor].Member("fred@chromium.org"))
	require.False(t, a.proxy.allowedRoles[roles.Admin].Member("fred@chromium.org"))
	require.True(t, a.proxy.allowedRoles[roles.Admin].Member("barney@google.com"))
}

func TestAppPopulateAllowedRoles_MultipleGroupsSameRoles_RoleContainsUnionOfAllows(t *testing.T) {
	a := newEmptyApp()
	a.criaClient = mockCriaClient(t)
	a.roleFlags = []string{
		"editor=cria_group:" + testCriaGroupName,
		"editor=google.com",
	}

	err := a.populateAllowedRoles()
	require.NoError(t, err)
	require.True(t, a.proxy.allowedRoles[roles.Editor].Member("fred@chromium.org"))
	require.True(t, a.proxy.allowedRoles[roles.Editor].Member("barney@google.com"))
}

func TestAppPopulateAllowedRoles_InvalidCriaGroup_ReturnsError(t *testing.T) {
	a := newEmptyApp()
	a.criaClient = mockCriaClient(t)
	a.roleFlags = []string{
		"editor=cria_group:this-is-not-a-valid-group",
		"admin=google.com",
	}

	err := a.populateAllowedRoles()
	require.Contains(t, err.Error(), "Failed parsing")
}

func TestAppPopulateAllowedRoles_UnknownRole_ReturnsError(t *testing.T) {
	a := newEmptyApp()
	a.criaClient = mockCriaClient(t)
	a.roleFlags = []string{
		"not-a-known-role=cria_group:" + testCriaGroupName,
	}

	err := a.populateAllowedRoles()
	require.Contains(t, err.Error(), "Invalid Role")
}

func TestAppPopulateAllowedRoles_BadFlagFormat_ReturnsError(t *testing.T) {
	a := newEmptyApp()
	a.criaClient = mockCriaClient(t)
	a.roleFlags = []string{
		"too=many=equals",
	}

	err := a.populateAllowedRoles()
	require.Contains(t, err.Error(), "Invalid format")
}

func TestAppPopulateAllowedRoles_CommaSeparatedRoles_RoleContainsUnionOfAllows(t *testing.T) {
	a := newEmptyApp()
	a.criaClient = mockCriaClient(t)
	a.roleFlags = []string{
		"viewer=google.com chromium.org",
	}

	err := a.populateAllowedRoles()
	require.NoError(t, err)
	require.True(t, a.proxy.allowedRoles[roles.Viewer].Member("fred@chromium.org"))
	require.True(t, a.proxy.allowedRoles[roles.Viewer].Member("barney@google.com"))
}

func TestAppPopulateAllowedRoles_TestMultiFlagParsing(t *testing.T) {
	a := newEmptyApp()
	err := a.Flagset().Parse([]string{"--role=viewer=google.com chromium.org", "--role=editor=google.com chromium.org"})
	require.NoError(t, err)
	expected := []string{
		"viewer=google.com chromium.org",
		"editor=google.com chromium.org",
	}
	require.Equal(t, expected, a.roleFlags)
}

func TestToAuthType_ValidType_ReturnsTypeUnchanged(t *testing.T) {
	for _, typ := range AllValidAuthTypes {
		require.Equal(t, typ, ToAuthType(string(typ)))
	}
}

func TestToAuthType_UnknownTypes_ReturnsInvalid(t *testing.T) {
	require.Equal(t, Invalid, ToAuthType("this is not a valid auth type"))
}

func TestParseTargetPort_OnlyPortIsSupplied_LocalhostUsedAsDomain(t *testing.T) {
	got, err := parseTargetPort(":8000")
	require.NoError(t, err)
	require.Equal(t, "http://localhost:8000", got.String())
}

func TestParseTargetPort_FullURLIsSupplied_LocalhostInNotAddedToDomain(t *testing.T) {
	got, err := parseTargetPort("http://foo:8000")
	require.NoError(t, err)
	require.Equal(t, "http://foo:8000", got.String())
}
