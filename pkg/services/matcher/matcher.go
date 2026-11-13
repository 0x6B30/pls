package matcher

import (
	"errors"
	"time"

	"github.com/0x6B30/pls/pkg/infra/dirs"
	"github.com/0x6B30/pls/pkg/infra/historystore"
)

var ErrNoMatch = errors.New("no visited folder matches")

type Matcher interface {
	Best(query, exclude string) (string, error)
}

func New(history historystore.Reader, dirs dirs.Checker, now func() time.Time) Matcher {
	return &frecencyMatcher{history: history, dirs: dirs, now: now}
}
