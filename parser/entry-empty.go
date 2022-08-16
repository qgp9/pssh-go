package parser

type parserEmpty struct {
	parserHelper[*entryEmpty]
}

func (p parserEmpty) getSelectorString() string {
	return `^\s*$`
}

type entryEmpty struct {
	entryHelper
}
