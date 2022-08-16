package parser

// for blesed parser* by parserHelper[T entryer]
type parseableHelper interface {
	parseable
	setRegexp(string) error
	getSelectorString() string
}

// for parser* struct
type parseable interface {
	Selector(string) bool
	Parse(c *PConfig) int
}

// for entry* struct
type entryer interface {
	GetSshConfig(c *PConfig) string
	SetValue(string)
}
