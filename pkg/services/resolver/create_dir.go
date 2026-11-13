package resolver

import (
	"context"

	"github.com/0x6B30/pls/pkg/infra/dirs"
)

type createDir struct {
	dirs dirs.Creator
}

func (s createDir) resolve(_ context.Context, req Request) (string, bool, error) {
	path := absolute(req.Cwd, req.Target)
	if err := s.dirs.Ensure(path); err != nil {
		return "", true, err
	}
	return path, true, nil
}
