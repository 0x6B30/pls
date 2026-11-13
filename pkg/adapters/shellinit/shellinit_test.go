package shellinit

import (
	"errors"
	"strings"
	"testing"
)

func TestEverySupportedShellDefinesCommands(t *testing.T) {
	for shell := range scriptsByShell {
		script, err := New().Script(shell)
		if err != nil {
			t.Fatalf("%s: %v", shell, err)
		}
		for _, fn := range []string{"pls()", "plsj()"} {
			if !strings.Contains(script, fn) {
				t.Errorf("%s script missing %s", shell, fn)
			}
		}
	}
}

func TestUnsupportedShell(t *testing.T) {
	if _, err := New().Script("tcsh"); !errors.Is(err, ErrUnsupportedShell) {
		t.Fatalf("got %v, want ErrUnsupportedShell", err)
	}
}
