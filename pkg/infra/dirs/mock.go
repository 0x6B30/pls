package dirs

type Mock struct {
	ExistsFunc func(path string) bool
	EnsureFunc func(path string) error
}

func (m *Mock) Exists(path string) bool {
	return m.ExistsFunc(path)
}

func (m *Mock) Ensure(path string) error {
	return m.EnsureFunc(path)
}
