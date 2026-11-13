package resolver

import (
	"context"
	"errors"
	"strings"

	"github.com/0x6B30/pls/pkg/services/matcher"
)

type historyMatch struct {
	matcher matcher.Matcher
}

func (s historyMatch) resolve(_ context.Context, req Request) (string, bool, error) {
	if looksLikePath(req.Target) {
		return "", false, nil
	}

	path, err := s.matcher.Best(req.Target, req.Cwd)
	if errors.Is(err, matcher.ErrNoMatch) {
		return "", false, nil
	}
	if err != nil {
		return "", true, err
	}
	return path, true, nil
}

func looksLikePath(target string) bool {
	return strings.ContainsRune(target, '/')
}
