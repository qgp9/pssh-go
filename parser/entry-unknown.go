package parser

type parserUnknown parserBase

type entryUnknown struct {
	entryBase
}

func (e *parserUnknown) Selector(line string, trimed string) bool {
	return true
}

func (e *parserUnknown) Parse(c *PConfig) int {
	entry := entryUnknown{}
	entry.value = c.currentLine()
	c.addEntry(&entry)
	return 1
}

func (e *entryUnknown) GetSshConfig(c *PConfig) string {
	return e.value
}
