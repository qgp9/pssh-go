package parser

type parserUnknown struct {
	parserHelper[*entryUnknown]
}

func NewParserUnknown() *parserUnknown {
	return NewParserWithSelector2(new(parserUnknown), `.`)
}

type entryUnknown struct {
	entryHelper
}
