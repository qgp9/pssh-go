package parser

// for blesed parser* by parserHelper[T entryable]
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
type entryable interface {
	GetSshConfig(c *PConfig) string
	SetValue(string)
}
