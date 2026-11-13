package historystore

import (
	"os"
	"path/filepath"
)

const historyDirPermissions = 0o700

// writeAtomic writes to a temp file and renames it into place, so a crash or
// a concurrent shell never leaves a half-written history behind.
func writeAtomic(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, historyDirPermissions); err != nil {
		return err
	}

	tmp, err := os.CreateTemp(dir, ".history-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}
