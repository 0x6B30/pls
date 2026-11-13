package historystore

import (
	"errors"
	"time"

	"github.com/0x6B30/pls/pkg/domains/history"
)

var ErrInvalidPath = errors.New("only absolute paths without tabs or newlines can be recorded")

type Reader interface {
	Entries() ([]history.Entry, error)
}

type Recorder interface {
	Record(path string, at time.Time) error
}

type Store interface {
	Reader
	Recorder
}

func New(filePath string) Store {
	return &fileStore{path: filePath}
}
