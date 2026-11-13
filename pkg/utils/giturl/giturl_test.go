package giturl

import (
	"errors"
	"testing"
)

func TestIsRemote(t *testing.T) {
	cases := map[string]bool{
		"git@github.com:user/repo.git": true,
		"https://github.com/user/repo": true,
		"ssh://git@host/repo.git":      true,
		"a/b/c":                        false,
		"repo.git":                     false,
		"ext::sh -c touch% /tmp/pwned": false,
		"-oProxyCommand=evil":          false,
	}
	for target, want := range cases {
		if got := IsRemote(target); got != want {
			t.Errorf("IsRemote(%q) = %v, want %v", target, got, want)
		}
	}
}

func TestRepoName(t *testing.T) {
	cases := map[string]string{
		"git@github.com:user/repo.git":  "repo",
		"https://github.com/user/repo/": "repo",
		"git@host:repo.git":             "repo",
		"ssh://git@host:22/team/svc":    "svc",
	}
	for url, want := range cases {
		got, err := RepoName(url)
		if err != nil || got != want {
			t.Errorf("RepoName(%q) = %q, %v; want %q", url, got, err, want)
		}
	}
}

func TestRepoNameRejectsUnnamedUrls(t *testing.T) {
	for _, url := range []string{"https://", "git@host:.git", "https://host/.."} {
		if _, err := RepoName(url); !errors.Is(err, ErrNoRepoName) {
			t.Errorf("RepoName(%q) should fail, got %v", url, err)
		}
	}
}
