package parser

type parserIgnore struct {
	parser[*entryIgnore]
}

type entryIgnore struct {
	entryBase
}

func (e *entryIgnore) GetSshConfig(c *PConfig) string {
	return ""
}

func NewParserIgnore() *parserIgnore {
	return NewParserWithSelector[*parserIgnore](`^\s*\+`)
}
