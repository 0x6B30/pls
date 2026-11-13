package gitcli

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

type gitCloner struct {
	progress io.Writer
}

func (g *gitCloner) Clone(ctx context.Context, url, dest string) error {
	cmd := exec.CommandContext(ctx, "git", "clone", "--", url, dest)
	cmd.Stdout = g.progress
	cmd.Stderr = g.progress

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %w", ErrCloneFailed, err)
	}
	return nil
}
