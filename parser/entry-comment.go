package parser

// interface: parseable
type parserComment struct {
	parserHelper[*entryComment]
}

// interface: entryable
type entryComment struct {
	entryBase
}

func NewParserComment() *parserComment {
	p := new(parserComment)
	p.setRegexp(`^\s*#`)
	return p
}
