package historystore

import (
	"time"

	"github.com/0x6B30/pls/pkg/domains/history"
)

type Mock struct {
	EntriesFunc func() ([]history.Entry, error)
	RecordFunc  func(path string, at time.Time) error
}

func (m *Mock) Entries() ([]history.Entry, error) {
	return m.EntriesFunc()
}

func (m *Mock) Record(path string, at time.Time) error {
	return m.RecordFunc(path, at)
}
