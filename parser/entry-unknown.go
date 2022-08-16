package parser

type parserUnknown struct {
	parser[*entryUnknown]
}

func NewParserUnknown() *parserUnknown {
	return NewParserWithSelector[*parserUnknown](`.`)
}

type entryUnknown struct {
	entryBase
}
