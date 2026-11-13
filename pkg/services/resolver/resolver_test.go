package resolver

import (
	"context"
	"errors"
	"testing"

	"github.com/0x6B30/pls/pkg/infra/dirs"
	"github.com/0x6B30/pls/pkg/infra/gitcli"
)

const cwd = "/work"

type fakes struct {
	existing map[string]bool
	created  []string
	cloned   []string
}

func (f *fakes) resolver() Resolver {
	return New(Dependencies{
		Cloner: &gitcli.Mock{CloneFunc: func(_ context.Context, url, dest string) error {
			f.cloned = append(f.cloned, url+" -> "+dest)
			return nil
		}},
		Dirs: &dirs.Mock{
			ExistsFunc: func(path string) bool { return f.existing[path] },
			EnsureFunc: func(path string) error { f.created = append(f.created, path); return nil },
		},
	})
}

func resolve(t *testing.T, f *fakes, req Request) string {
	t.Helper()
	req.Cwd = cwd
	path, err := f.resolver().Resolve(context.Background(), req)
	if err != nil {
		t.Fatalf("resolve %+v: %v", req, err)
	}
	return path
}

func TestClonesNewRepo(t *testing.T) {
	f := &fakes{}

	got := resolve(t, f, Request{Target: "git@github.com:me/tool.git"})

	if got != "/work/tool" || len(f.cloned) != 1 {
		t.Fatalf("got %q, cloned %v", got, f.cloned)
	}
}

func TestReusesAlreadyClonedRepo(t *testing.T) {
	f := &fakes{existing: map[string]bool{"/work/custom": true}}

	got := resolve(t, f, Request{Target: "https://github.com/me/tool", CloneDir: "custom"})

	if got != "/work/custom" || len(f.cloned) != 0 {
		t.Fatalf("got %q, cloned %v", got, f.cloned)
	}
}

func TestReturnsExistingDirWithoutCreating(t *testing.T) {
	f := &fakes{existing: map[string]bool{"/work/api": true}}

	got := resolve(t, f, Request{Target: "api"})

	if got != "/work/api" || len(f.created) != 0 {
		t.Fatalf("got %q, created %v", got, f.created)
	}
}

func TestCreatesMissingDir(t *testing.T) {
	f := &fakes{}

	got := resolve(t, f, Request{Target: "a/b/c"})

	if got != "/work/a/b/c" || len(f.created) != 1 {
		t.Fatalf("got %q, created %v", got, f.created)
	}
}

func TestRejectsInvalidRequests(t *testing.T) {
	cases := map[error]Request{
		ErrEmptyTarget:         {},
		ErrCloneDirWithoutRepo: {Target: "plain", CloneDir: "x"},
	}
	for want, req := range cases {
		if _, err := (&fakes{}).resolver().Resolve(context.Background(), req); !errors.Is(err, want) {
			t.Errorf("Resolve(%+v) = %v, want %v", req, err, want)
		}
	}
}
