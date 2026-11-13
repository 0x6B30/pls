package shellinit

import "errors"

var ErrUnsupportedShell = errors.New("unsupported shell (supported: zsh, bash)")

type Generator interface {
	Script(shell string) (string, error)
}

func New() Generator {
	return embeddedScripts{}
}
