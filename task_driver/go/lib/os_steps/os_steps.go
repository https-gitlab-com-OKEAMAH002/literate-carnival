package os_steps

/*
	Canned steps to be used for performing OS and filesystem interactions in task drivers.
*/

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"go.skia.org/infra/go/util"
	"go.skia.org/infra/task_driver/go/td"
)

// Stat is a wrapper for os.Stat.
func Stat(ctx context.Context, path string) (os.FileInfo, error) {
	var rv os.FileInfo
	return rv, td.Do(ctx, td.Props(fmt.Sprintf("Stat %s", path)).Infra(), func(context.Context) error {
		fi, err := os.Stat(path)
		rv = fi
		return err
	})
}

// MkdirAll is a wrapper for os.MkdirAll.
func MkdirAll(ctx context.Context, path string) (err error) {
	return td.Do(ctx, td.Props(fmt.Sprintf("MkdirAll %s", path)).Infra(), func(context.Context) error {
		return os.MkdirAll(path, os.ModePerm)
	})
}

type tempDirContextKeyType string

// TempDirContextKey can be used to overwrite TempDir in tests.
const TempDirContextKey = tempDirContextKeyType("overwriteTempDir")

// TempDir is a wrapper for os.MkdirTemp. In swarming, this file might not be in /tmp, but in a
// subdirectory of the swarming task work directory.
//
// Use the TempDirContextKey to overwrite os.MkdirTemp with another function with the same
// signature.
func TempDir(ctx context.Context, dir, pattern string) (string, error) {
	var tempDir string
	err := td.Do(ctx, td.Props("Creating TempDir").Infra(), func(context.Context) error {
		var d string
		var err error
		if fn := ctx.Value(TempDirContextKey); fn != nil {
			d, err = fn.(func(string, string) (string, error))(dir, pattern)
		} else {
			d, err = os.MkdirTemp(dir, pattern)
		}

		if err != nil {
			return err
		}
		tempDir = d
		return nil
	})
	return tempDir, err
}

// RemoveAll is a wrapper for os.RemoveAll.
func RemoveAll(ctx context.Context, path string) (err error) {
	return td.Do(ctx, td.Props(fmt.Sprintf("RemoveAll %s", path)).Infra(), func(context.Context) error {
		return os.RemoveAll(path)
	})
}

// Abs is a wrapper for filepath.Abs.
func Abs(ctx context.Context, path string) (string, error) {
	var rv string
	err := td.Do(ctx, td.Props(fmt.Sprintf("Abs %s", path)).Infra(), func(context.Context) error {
		var err error
		rv, err = filepath.Abs(path)
		return err
	})
	return rv, err
}

// ReadFile is a wrapper for os.ReadFile.
func ReadFile(ctx context.Context, path string) ([]byte, error) {
	var rv []byte
	err := td.Do(ctx, td.Props(fmt.Sprintf("Read %s", path)).Infra(), func(context.Context) error {
		var err error
		rv, err = os.ReadFile(path)
		return err
	})
	return rv, err
}

// ReadDir is a wrapper for os.ReadDir.
func ReadDir(ctx context.Context, path string) ([]fs.DirEntry, error) {
	var rv []fs.DirEntry
	err := td.Do(ctx, td.Props(fmt.Sprintf("ReadDir %s", path)).Infra(), func(context.Context) error {
		var err error
		rv, err = os.ReadDir(path)
		return err
	})
	return rv, err
}

// Rename is a wrapper for os.Rename.
func Rename(ctx context.Context, oldpath, newpath string) error {
	return td.Do(ctx, td.Props(fmt.Sprintf("Rename %s %s", oldpath, newpath)).Infra(), func(context.Context) error {
		return os.Rename(oldpath, newpath)
	})
}

// WriteFile is a wrapper for os.WriteFile.
func WriteFile(ctx context.Context, path string, data []byte, perm os.FileMode) error {
	return td.Do(ctx, td.Props(fmt.Sprintf("Write %s", path)).Infra(), func(context.Context) error {
		return os.WriteFile(path, data, perm)
	})
}

// Which returns the result of "which <exe>" (or "where <exe>" on Windows).
func Which(ctx context.Context, exe string) (string, error) {
	var rv string
	err := td.Do(ctx, td.Props(fmt.Sprintf("Which %s", exe)).Infra(), func(context.Context) error {
		var err error
		rv, err = exec.LookPath(exe)
		return err
	})
	return rv, err
}

// CopyFile copies the given file.
func CopyFile(ctx context.Context, src, dst string) error {
	return td.Do(ctx, td.Props(fmt.Sprintf("Copy %s %s", src, dst)).Infra(), func(context.Context) (rvErr error) {
		return util.CopyFile(src, dst)
	})
}

// Chdir is a wrapper for os.Chdir.
func Chdir(ctx context.Context, dir string) error {
	return td.Do(ctx, td.Props(fmt.Sprintf("chdir %s", dir)).Infra(), func(context.Context) error {
		return os.Chdir(dir)
	})
}
