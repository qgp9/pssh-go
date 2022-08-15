package parser

type parserBase struct {
}

func (e *parserBase) Parse(c *PConfig) int {
	return -1
}

type entryBase struct {
	line  int
	value string
}

func (e entryBase) GetSshConfig(c *PConfig) string {
	return ""
}
