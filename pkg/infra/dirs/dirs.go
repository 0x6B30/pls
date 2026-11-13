package dirs

type Checker interface {
	Exists(path string) bool
}

type Creator interface {
	Ensure(path string) error
}

type Dirs interface {
	Checker
	Creator
}

func New() Dirs {
	return osDirs{}
}
