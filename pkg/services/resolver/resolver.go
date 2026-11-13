package resolver

import (
	"context"
	"errors"

	"github.com/0x6B30/pls/pkg/infra/dirs"
	"github.com/0x6B30/pls/pkg/infra/gitcli"
)

var (
	ErrEmptyTarget         = errors.New("no target given")
	ErrCloneDirWithoutRepo = errors.New("a destination folder is only allowed when cloning a repository")
	ErrUnresolved          = errors.New("could not resolve target")
)

type Request struct {
	Target   string
	CloneDir string
	Cwd      string
}

type Resolver interface {
	Resolve(ctx context.Context, req Request) (string, error)
}

type Dependencies struct {
	Cloner gitcli.Cloner
	Dirs   dirs.Dirs
}

func New(deps Dependencies) Resolver {
	return &chain{strategies: []strategy{
		cloneRepo{cloner: deps.Cloner, dirs: deps.Dirs},
		existingDir{dirs: deps.Dirs},
		createDir{dirs: deps.Dirs},
	}}
}
