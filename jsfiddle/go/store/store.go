// Stores and retrieves jsfiddles in Google Storage.
package store

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"

	"cloud.google.com/go/storage"
	"go.skia.org/infra/go/auth"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/skerr"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const (
	JSFIDDLE_STORAGE_BUCKET = "skia-jsfiddle"
)

// Store is used to read and write user code and media to and from Google
// Storage.
type Store interface {
	// PutCode writes the code to Google Storage.
	// Returns the fiddleHash.
	PutCode(code, fiddleType string) (string, error)

	// GetCode reads the code from Google Storage.
	// Returns the code.
	GetCode(hash, fiddleType string) (string, error)
}

// store implements Store.
type store struct {
	bucket *storage.BucketHandle
}

// New creates a new store.
//
// local - True if running locally.
func New(ctx context.Context, local bool) (*store, error) {
	ts, err := google.DefaultTokenSource(ctx, auth.ScopeReadWrite)
	if err != nil {
		return nil, skerr.Wrapf(err, "problem setting up client OAuth")
	}
	client := httputils.DefaultClientConfig().WithTokenSource(ts).With2xxOnly().Client()
	storageClient, err := storage.NewClient(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, skerr.Wrapf(err, "problem creating storage client")
	}
	return &store{
		bucket: storageClient.Bucket(JSFIDDLE_STORAGE_BUCKET),
	}, nil
}

// PutCode implements Store.
func (s *store) PutCode(code, fiddleType string) (string, error) {
	hash := computeHash(code)

	path := strings.Join([]string{fiddleType, hash, "draw.js"}, "/")
	w := s.bucket.Object(path).NewWriter(context.Background())
	w.ObjectAttrs.ContentEncoding = "text/plain"

	if n, err := w.Write([]byte(code)); err != nil {
		_ = w.Close()
		return "", skerr.Wrapf(err, "There was a problem storing the code. Uploaded %d bytes", n)
	}
	if err := w.Close(); err != nil {
		return "", skerr.Wrap(err)
	}
	return hash, nil
}

// GetCode implements Store.
func (s *store) GetCode(hash, fiddleType string) (string, error) {
	path := strings.Join([]string{fiddleType, hash, "draw.js"}, "/")
	o := s.bucket.Object(path)
	r, err := o.NewReader(context.Background())
	if err != nil {
		return "", skerr.Wrapf(err, "failed to open source file for %s", hash)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return "", skerr.Wrapf(err, "failed to read source file for %s", hash)
	}
	return string(b), nil
}

func computeHash(code string) string {
	sum := sha256.Sum256([]byte(code))
	return fmt.Sprintf("%x", sum)
}

// Confirm the *store implements Store.
var _ Store = (*store)(nil)
