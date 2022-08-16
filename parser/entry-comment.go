package parser

// interface: parseable
type parserComment struct {
	parserHelper[*entryComment]
}

// interface: entryable
type entryComment struct {
	entryHelper
}

func NewParserComment() *parserComment {
	p := new(parserComment)
	p.setRegexp(`^\s*#`)
	return p
}
