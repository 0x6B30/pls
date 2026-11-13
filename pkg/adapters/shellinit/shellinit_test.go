package shellinit

import (
	"errors"
	"strings"
	"testing"
)

func TestEverySupportedShellDefinesPls(t *testing.T) {
	for shell := range scriptsByShell {
		script, err := New().Script(shell)
		if err != nil || !strings.Contains(script, "pls()") {
			t.Errorf("%s script missing pls function (err %v)", shell, err)
		}
	}
}

func TestUnsupportedShell(t *testing.T) {
	if _, err := New().Script("tcsh"); !errors.Is(err, ErrUnsupportedShell) {
		t.Fatalf("got %v, want ErrUnsupportedShell", err)
	}
}
