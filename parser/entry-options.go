package parser

import (
	"regexp"
	"strings"

	"github.com/qgp9/pssh-go/utils"
)

type parserOption struct {
	parserHelper[*entryOption]
}

var reOptionPrefix = regexp.MustCompile(`^([\s\|]*)-\s*(.*)`)

func (e *parserOption) Selector(line string) bool {
	return reOptionPrefix.MatchString(line)
}

func (e *parserOption) Parse(c *PConfig) int {
	// TODO: comment
	var entry = entryOption{}
	value := reOptionPrefix.ReplaceAllString(c.currentLine(), "$2")
	var option, comment string
	utils.SplitAssign(value, "#", &option, &comment)
	strings.TrimSpace(option)
	strings.TrimRight(option, "|")
	strings.TrimSpace(option)
	option = c.applyVariable(option)
	value = "# " + comment + "\n" + option
	entry.value = value
	c.addEntry(&entry)
	return 1
}

func NewParserOption() *parserOption {
	return NewParser[*parserOption]()
}

type entryOption struct {
	entryHelper
}
