package dirs

import "os"

const newDirPermissions = 0o755

type osDirs struct{}

func (osDirs) Exists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func (osDirs) Ensure(path string) error {
	return os.MkdirAll(path, newDirPermissions)
}
