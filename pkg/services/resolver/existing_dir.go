package resolver

import (
	"context"

	"github.com/0x6B30/pls/pkg/infra/dirs"
)

type existingDir struct {
	dirs dirs.Checker
}

func (s existingDir) resolve(_ context.Context, req Request) (string, bool, error) {
	path := absolute(req.Cwd, req.Target)
	if !s.dirs.Exists(path) {
		return "", false, nil
	}
	return path, true, nil
}
