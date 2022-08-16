package parser

type parserGroup struct {
	parserHelper[*entryGroup]
}

type entryGroup struct {
	entryHelper
}

func (e *entryGroup) GetSshConfig(c *PConfig) string {
	return "# " + e.value
}

func NewParserGroup() *parserGroup {
	return NewParserWithSelector[*parserGroup](`^\s*@@`)
}
