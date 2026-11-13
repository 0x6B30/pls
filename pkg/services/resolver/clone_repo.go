package resolver

import (
	"context"

	"github.com/0x6B30/pls/pkg/infra/dirs"
	"github.com/0x6B30/pls/pkg/infra/gitcli"
	"github.com/0x6B30/pls/pkg/utils/giturl"
)

type cloneRepo struct {
	cloner gitcli.Cloner
	dirs   dirs.Checker
}

func (s cloneRepo) resolve(ctx context.Context, req Request) (string, bool, error) {
	if !giturl.IsRemote(req.Target) {
		return "", false, nil
	}

	folder, err := cloneFolder(req)
	if err != nil {
		return "", true, err
	}

	dest := absolute(req.Cwd, folder)
	if s.dirs.Exists(dest) {
		return dest, true, nil
	}
	if err := s.cloner.Clone(ctx, req.Target, dest); err != nil {
		return "", true, err
	}
	return dest, true, nil
}

func cloneFolder(req Request) (string, error) {
	if req.CloneDir != "" {
		return req.CloneDir, nil
	}
	return giturl.RepoName(req.Target)
}
