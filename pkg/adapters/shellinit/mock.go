package shellinit

type Mock struct {
	ScriptFunc func(shell string) (string, error)
}

func (m *Mock) Script(shell string) (string, error) {
	return m.ScriptFunc(shell)
}
