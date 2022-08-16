package parser

// interface: parseable
type parserComment struct {
	parserHelper[*entryComment]
}

// interface: entryer
type entryComment struct {
	entryBase
}

func NewParserComment() *parserComment {
	p := new(parserComment)
	p.setRegexp(`^\s*#`)
	return p
}
