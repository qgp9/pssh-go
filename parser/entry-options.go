package parser

import "regexp"

type parserOption struct {
	parser[*entryOption]
}

var reOptionPrefix = regexp.MustCompile(`^(\s*\|\s*)*-\s*`)

func (e *parserOption) Selector(line string) bool {
	return reOptionPrefix.MatchString(line)
}

func (e *parserOption) Parse(c *PConfig) int {
	var entry = entryOption{}
	entry.value = reOptionPrefix.ReplaceAllString(c.currentLine(), "")
	entry.value = c.applyVariable(entry.value)
	c.addEntry(&entry)
	return 1
}

func NewParserOption() *parserOption {
	return NewParser[*parserOption]()
}

type entryOption struct {
	entryBase
}
