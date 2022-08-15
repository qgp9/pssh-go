package parser

import "strings"

type parserGroup parserBase

func (e *parserGroup) Selector(line string, trimed string) bool {
	return strings.HasPrefix(trimed, "@@")
}

func (e *parserGroup) Parse(c *PConfig) int {
	var entry = entryGroup{}
	entry.value = c.lines[c.i] //FIXME: remove @@
	c.addEntry(&entry)
	return 1
}

type entryGroup struct {
	entryBase
}

func (e *entryGroup) GetSshConfig(c *PConfig) string {
	return "# " + e.value
}
