package historystore

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/0x6B30/pls/pkg/domains/history"
)

type fileStore struct {
	path string
}

func (s *fileStore) Entries() ([]history.Entry, error) {
	file, err := os.Open(s.path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return decode(file)
}

func (s *fileStore) Record(path string, at time.Time) error {
	if !isRecordable(path) {
		return ErrInvalidPath
	}

	entries, err := s.Entries()
	if err != nil {
		return err
	}

	entries = history.Visit(entries, filepath.Clean(path), at)
	entries = history.Age(entries)
	return writeAtomic(s.path, encode(entries))
}

func isRecordable(path string) bool {
	return filepath.IsAbs(path) && !strings.ContainsAny(path, "\t\n\r")
}
