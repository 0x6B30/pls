package gitcli

import (
	"context"
	"errors"
	"io"
)

var ErrCloneFailed = errors.New("git clone failed")

type Cloner interface {
	Clone(ctx context.Context, url, dest string) error
}

// New sends git's own output to progress, keeping stdout free for the
// resolved path that the shell wrapper reads.
func New(progress io.Writer) Cloner {
	return &gitCloner{progress: progress}
}
