package parser

import (
	"regexp"
)

type parserOption struct {
	parserHelper[*entryOption]
}

var reOptionPrefix = regexp.MustCompile(`^([\s\|]*)-\s*(.*?)([\s\|]*)$`)

func (e *parserOption) Selector(line string) bool {
	return reOptionPrefix.MatchString(line)
}

func (e *parserOption) Parse(c *PConfig) int {
	// TODO: comment
	var entry = entryOption{}
	entry.value = reOptionPrefix.ReplaceAllString(c.currentLine(), "$2")
	entry.value = c.applyVariable(entry.value)
	c.addEntry(&entry)
	return 1
}

func NewParserOption() *parserOption {
	return NewParser[*parserOption]()
}

type entryOption struct {
	entryHelper
}
