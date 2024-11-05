// Package header supports extracting the email of an authorized user from a
// protobuf in an HTTP Header.
package protoheader

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"go.skia.org/infra/go/secret"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/kube/go/authproxy/auth"
	"google.golang.org/protobuf/proto"
)

const (
	// HeaderSecretName is the name of the GCP secret for login.
	HeaderSecretName = "authproxy-protoheader-name"

	// LoginURNSecretName is the name of the GCP secret for the login URL.
	LoginURNSecretName = "authproxy-loginurl"

	// Project is the project where the above secrets are stored in.
	Project = "skia-infra-public"
)

var (
	errDotInHeaderRequired = errors.New("Failed to find a '.' separated header value.")
)

// ProtoHeader implements auth.Auth.
type ProtoHeader struct {
	headerName string
	loginURL   string
}

// New creates a ProtoHeader.
func New(ctx context.Context, secretClient secret.Client) (ProtoHeader, error) {
	var ret ProtoHeader
	headerName, err := secretClient.Get(ctx, Project, HeaderSecretName, secret.VersionLatest)
	if err != nil {
		return ret, skerr.Wrapf(err, "failed loading secrets from GCP secret manager; failed to retrieve secret %q", HeaderSecretName)
	}
	ret.headerName = headerName

	loginURL, err := secretClient.Get(ctx, Project, LoginURNSecretName, secret.VersionLatest)
	if err != nil {
		return ret, skerr.Wrapf(err, "failed loading secrets from GCP secret manager; failed to retrieve secret %q", LoginURNSecretName)
	}
	ret.loginURL = loginURL

	return ret, nil
}

// Init implements auth.Auth.
func (p ProtoHeader) Init(ctx context.Context) error {
	return nil
}

// LoggedInAs implements auth.Auth.
func (p ProtoHeader) LoggedInAs(r *http.Request) (string, error) {
	// The header value contains a base64 encoded proto and then it's signed
	// which adds a signature, separated by a period.
	headerValue := r.Header.Get(p.headerName)
	parts := strings.Split(headerValue, ".")
	if len(parts) != 2 {
		return "", errDotInHeaderRequired
	}
	// Only use the base64 encoded data before the period.
	b, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", skerr.Wrapf(err, "decoding base64 header: %q", p.headerName)
	}
	var h Header
	err = proto.Unmarshal(b, &h)
	if err != nil {
		return "", skerr.Wrapf(err, "decoding proto %q", p.headerName)
	}
	return h.Email, nil
}

// LoginURL implements auth.Auth.
func (p ProtoHeader) LoginURL(w http.ResponseWriter, r *http.Request) string {
	return p.loginURL
}

// Confirm we implement the interface.
var _ auth.Auth = ProtoHeader{}
