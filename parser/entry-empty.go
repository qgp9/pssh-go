package parser

type parserEmpty struct {
	parser[*entryEmpty]
}

type entryEmpty struct {
	entryBase
}

func (e *entryEmpty) GetSshConfig(c *PConfig) string {
	return ""
}

func NewParserEmpty() *parserEmpty {
	return NewParserWithSelector[*parserEmpty](`^\s*$`)
}
