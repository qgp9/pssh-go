package parser

import "regexp"

type parserOption parserBase

var reOption = regexp.MustCompile(`^(\|\s*)*-\s*`)

func (e *parserOption) Selector(line string, trimed string) bool {
	return reOption.MatchString(trimed)
	//return strings.HasPrefix(trimed, "-")
}

var reOptionPrefix = reOption

func (e *parserOption) Parse(c *PConfig) int {
	var entry = entryOption{}
	entry.value = reOptionPrefix.ReplaceAllString(c.currentLine(), "")
	entry.value = c.applyVariable(entry.value)
	//entry.value = strings.Trim(strings.TrimSpace(c.lines[c.i]), "-")
	c.addEntry(&entry)
	return 1
}

type entryOption struct {
	entryBase
}

func (e *entryOption) GetSshConfig(c *PConfig) string {
	return e.value
}
