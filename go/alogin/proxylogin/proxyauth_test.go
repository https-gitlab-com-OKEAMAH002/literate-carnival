package proxylogin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/alogin"
	"go.skia.org/infra/go/roles"
	"go.skia.org/infra/kube/go/authproxy"
)

const (
	goodHeaderName                 = authproxy.WebAuthHeaderName
	unknownHeaderName              = "X-SOME-UNKNOWN-HEADER"
	email             alogin.EMail = "someone@example.org"
	emailAsString     string       = string(email)
)

func TestLoggedInAs_HeaderIsMissing_ReturnsEmptyString(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	login, err := New(unknownHeaderName, "")
	require.NoError(t, err)
	require.Equal(t, alogin.NotLoggedIn, login.LoggedInAs(r))
}

func TestLoggedInAs_HeaderPresent_ReturnsUserEmail(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set(goodHeaderName, emailAsString)
	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	require.Equal(t, email, login.LoggedInAs(r))
}

func TestLoggedInAs_RegexProvided_ReturnsUserEmail(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set(goodHeaderName, "accounts.google.com:"+emailAsString)
	login, err := New(goodHeaderName, "accounts.google.com:(.*)")
	require.NoError(t, err)
	require.Equal(t, email, login.LoggedInAs(r))
}

func TestLoggedInAs_RegexHasTooManySubGroups_ReturnsEmptyString(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set(goodHeaderName, emailAsString)
	login, err := New(goodHeaderName, "(too)(many)(subgroups)")
	require.NoError(t, err)
	require.Equal(t, alogin.NotLoggedIn, login.LoggedInAs(r))
}

func TestNeedsAuthentication_EmitsStatusForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	login.NeedsAuthentication(w, r)
	require.Equal(t, http.StatusForbidden, w.Result().StatusCode)
}

func TestStatus_HeaderPresent_ReturnsUserEmail(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set(authproxy.WebAuthHeaderName, emailAsString)
	r.Header.Set(authproxy.WebAuthRoleHeaderName, roles.Roles{roles.Admin}.ToHeader())
	expected := alogin.Status{
		EMail: email,
		Roles: roles.Roles{roles.Admin},
	}
	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	require.Equal(t, expected, login.Status(r))
}

func TestNew_InvalidRegex_ReturnsError(t *testing.T) {
	_, err := New(goodHeaderName, "\\y")
	require.Error(t, err)
}

func TestRoles_HeaderPresent_ReturnAllRoles(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set(authproxy.WebAuthRoleHeaderName, roles.AllValidRoles.ToHeader())
	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	require.Equal(t, roles.AllValidRoles, login.Roles(r))
}

func TestRoles_HeaderMissing_ReturnsEmptyListOfRoles(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	require.Empty(t, login.Roles(r))
}

func TestHasRoles_HeaderPresent_ReturnsTrue(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set(authproxy.WebAuthRoleHeaderName, roles.AllValidRoles.ToHeader())
	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	require.True(t, login.HasRole(r, roles.Admin))
}

func TestHasRoles_HeaderMissingPresent_ReturnsFalse(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	require.False(t, login.HasRole(r, roles.Admin))
}

func testLoginURL(t *testing.T, expected, domain string) {
	t.Helper()

	login, err := New(goodHeaderName, "")
	require.NoError(t, err)
	r := httptest.NewRequest("GET", "/", nil)
	r.Host = domain

	require.Equal(t, expected, login.LoginURL(r))

}

func TestAuthImpl_LoginURL(t *testing.T) {
	testLoginURL(t, "https://skia.org/login/", "foo.skia.org")
	testLoginURL(t, "https://luci.app/login/", "perf.luci.app")
}
