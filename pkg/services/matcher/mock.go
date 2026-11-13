package matcher

type Mock struct {
	BestFunc func(query, exclude string) (string, error)
}

func (m *Mock) Best(query, exclude string) (string, error) {
	return m.BestFunc(query, exclude)
}
