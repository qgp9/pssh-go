package parser

type parserIgnore struct {
	parserHelper[*entryIgnore]
}

type entryIgnore struct {
	entryHelper
}

func (e *entryIgnore) GetSshConfig(c *PConfig) string {
	return ""
}

func NewParserIgnore() *parserIgnore {
	return NewParserWithSelector[*parserIgnore](`^\s*\+`)
}
