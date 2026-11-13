package resolver

import (
	"context"
	"path/filepath"

	"github.com/0x6B30/pls/pkg/utils/giturl"
)

type strategy interface {
	resolve(ctx context.Context, req Request) (path string, handled bool, err error)
}

type chain struct {
	strategies []strategy
}

func (c *chain) Resolve(ctx context.Context, req Request) (string, error) {
	if err := validate(req); err != nil {
		return "", err
	}

	for _, s := range c.strategies {
		path, handled, err := s.resolve(ctx, req)
		if handled || err != nil {
			return path, err
		}
	}
	return "", ErrUnresolved
}

func validate(req Request) error {
	if req.Target == "" {
		return ErrEmptyTarget
	}
	if req.CloneDir != "" && !giturl.IsRemote(req.Target) {
		return ErrCloneDirWithoutRepo
	}
	return nil
}

func absolute(cwd, path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	return filepath.Join(cwd, path)
}
