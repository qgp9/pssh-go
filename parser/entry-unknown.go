package parser

type parserUnknown struct {
	parserHelper[*entryUnknown]
}

func NewParserUnknown() *parserUnknown {
	return NewParserWithSelector[*parserUnknown](`.`)
}

type entryUnknown struct {
	entryHelper
}
