package parser

type parserEmpty struct {
	parserHelper[*entryEmpty]
}

type entryEmpty struct {
	entryHelper
}

func NewParserEmpty() *parserEmpty {
	return NewParserWithSelector[*parserEmpty](`^\s*$`)
}
