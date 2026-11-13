package gitcli

import "context"

type Mock struct {
	CloneFunc func(ctx context.Context, url, dest string) error
}

func (m *Mock) Clone(ctx context.Context, url, dest string) error {
	return m.CloneFunc(ctx, url, dest)
}
