package historystore

import (
	"bufio"
	"bytes"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/0x6B30/pls/pkg/domains/history"
)

const fieldSeparator = "\t"

func encode(entries []history.Entry) []byte {
	var buf bytes.Buffer
	for _, entry := range entries {
		buf.WriteString(entry.Path)
		buf.WriteString(fieldSeparator)
		buf.WriteString(strconv.FormatFloat(entry.Visits, 'f', -1, 64))
		buf.WriteString(fieldSeparator)
		buf.WriteString(strconv.FormatInt(entry.LastVisit.Unix(), 10))
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func decode(r io.Reader) ([]history.Entry, error) {
	var entries []history.Entry
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		entry, ok := decodeLine(scanner.Text())
		if ok {
			entries = append(entries, entry)
		}
	}
	return entries, scanner.Err()
}

// decodeLine skips anything malformed so a damaged file never blocks a cd.
func decodeLine(line string) (history.Entry, bool) {
	fields := strings.Split(line, fieldSeparator)
	if len(fields) != 3 || fields[0] == "" {
		return history.Entry{}, false
	}

	visits, err := strconv.ParseFloat(fields[1], 64)
	if err != nil || !isValidVisitCount(visits) {
		return history.Entry{}, false
	}

	unixSeconds, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return history.Entry{}, false
	}

	return history.Entry{Path: fields[0], Visits: visits, LastVisit: time.Unix(unixSeconds, 0)}, true
}

func isValidVisitCount(visits float64) bool {
	return visits > 0 && !math.IsInf(visits, 0)
}
