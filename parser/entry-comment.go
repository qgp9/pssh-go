package parser

import (
	"strings"
)

type parserComment parserBase

func (e *parserComment) Selector(line string, trimed string) bool {
	return strings.HasPrefix(trimed, "#")
}

func (e *parserComment) Parse(c *PConfig) int {
	var entry = entryComment{}
	entry.value = c.lines[c.i]
	c.addEntry(&entry)
	return 1
}

type entryComment struct {
	entryBase
}

func (e *entryComment) GetSshConfig(c *PConfig) string {
	return e.value
}
