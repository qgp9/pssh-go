package parser

type parserGroup struct {
	parserHelper[*entryGroup]
}

type entryGroup struct {
	entryNode
	entryHelper
}

func (e *entryGroup) GetSshConfig(c *PConfig) string {
	return "# GROUP: " + e.value + "\n"
}

func NewParserGroup() *parserGroup {
	return NewParserWithSelector[*parserGroup](`^\s*@@`)
}
