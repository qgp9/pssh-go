package parser

import (
	"log"
	"strings"
)

type parserIgnore parserBase

func (e *parserIgnore) Selector(line string, trimed string) bool {
	return strings.HasPrefix(trimed, "+")
}

func (e *parserIgnore) Parse(c *PConfig) int {
	log.Println("entryIgnore: ", c.currentLine())
	var entry = entryIgnore{}
	entry.value = c.lines[c.i] //FIXME: remove @@
	c.addEntry(&entry)
	return 1
}

type entryIgnore struct {
	entryBase
}

func (e *entryIgnore) GetSshConfig(c *PConfig) string {
	return ""
}
