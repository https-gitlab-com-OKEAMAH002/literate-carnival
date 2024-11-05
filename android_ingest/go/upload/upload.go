package upload

import (
	"crypto/md5"
	"fmt"
	"path"
	"strings"
	"time"
)

// ObjectPath returns the Google Cloud Storage path where the JSON serialization
// of benchData should be stored.
//
// gcsPath will be the root of the path.
// now is the time which will be encoded in the path.
// b is the source bytes of the incoming file.
func ObjectPath(key map[string]string, gitHash string, gcsPath string, now time.Time, b []byte) string {
	hash := fmt.Sprintf("%x", md5.Sum(b))
	keyparts := []string{}
	for k, v := range key {
		keyparts = append(keyparts, k, v)
	}
	filename := fmt.Sprintf("%s_%s_%s.json", gitHash, strings.Join(keyparts, "_"), hash)
	return path.Join(gcsPath, now.Format("2006/01/02/15/"), filename)
}

// LogPath returns the Google Cloud Storage path where the raw POST data
// should be stored.
//
// gcsPath will be the root of the path.
// now is the time which will be encoded in the path.
// b is the source bytes of the incoming file.
func LogPath(gcsPath string, now time.Time, b []byte) string {
	hash := fmt.Sprintf("%x", md5.Sum(b))
	filename := fmt.Sprintf("%s.json", hash)
	return path.Join(gcsPath, now.Format("2006/01/02/15/"), filename)
}
