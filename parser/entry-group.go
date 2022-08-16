package parser

type parserGroup struct {
	parser[*entryGroup]
}

type entryGroup struct {
	entryBase
}

func (e *entryGroup) GetSshConfig(c *PConfig) string {
	return "# " + e.value
}

func NewParserGroup() *parserGroup {
	return NewParserWithSelector[*parserGroup](`^\s*@@`)
}
