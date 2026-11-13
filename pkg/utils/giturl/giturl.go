package giturl

import (
	"errors"
	"strings"
)

var ErrNoRepoName = errors.New("cannot derive a folder name from the repository url")

var remotePrefixes = []string{"git@", "ssh://", "https://", "http://", "git://"}

func IsRemote(target string) bool {
	for _, prefix := range remotePrefixes {
		if strings.HasPrefix(target, prefix) {
			return true
		}
	}
	return false
}

func RepoName(url string) (string, error) {
	name := strings.TrimRight(url, "/")
	name = name[strings.LastIndexAny(name, "/:")+1:]
	name = strings.TrimSuffix(name, ".git")

	if name == "" || name == "." || name == ".." {
		return "", ErrNoRepoName
	}
	return name, nil
}
