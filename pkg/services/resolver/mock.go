package resolver

import "context"

type Mock struct {
	ResolveFunc func(ctx context.Context, req Request) (string, error)
}

func (m *Mock) Resolve(ctx context.Context, req Request) (string, error) {
	return m.ResolveFunc(ctx, req)
}
