package parser

type parserEmpty parserBase

func (e *parserEmpty) Selector(line string, trimed string) bool {
	return trimed == ""
}

func (e *parserEmpty) Parse(c *PConfig) int {
	c.addEntry(&entryEmpty{})
	return 1
}

type entryEmpty struct {
	entryBase
}

func (e *entryEmpty) GetSshConfig(c *PConfig) string {
	return ""
}
