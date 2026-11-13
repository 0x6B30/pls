package shellinit

import _ "embed"

//go:embed scripts/zsh.sh
var zshScript string

//go:embed scripts/bash.sh
var bashScript string

var scriptsByShell = map[string]string{
	"zsh":  zshScript,
	"bash": bashScript,
}

type embeddedScripts struct{}

func (embeddedScripts) Script(shell string) (string, error) {
	script, ok := scriptsByShell[shell]
	if !ok {
		return "", ErrUnsupportedShell
	}
	return script, nil
}
