package parser

import (
	"regexp"
	"strings"
)

type parserVariable struct {
	parserHelper[*entryVariable]
}

func NewParserVariable() *parserVariable {
	return NewParserWithSelector[*parserVariable](`^\s*\$`)
}

type entryVariable struct {
	entryHelper
	name    string
	comment string
}

func (e *parserVariable) Parse(c *PConfig) int {
	//log.Println("entryVariable: ", c.lines[c.i])
	var entry = entryVariable{}
	rSingle, _ := regexp.Compile(`^\$([\w]+)\s*=\s*([^#]+)(.*)$`)
	rBlock, _ := regexp.Compile(`^\$([\w]+)\s*=\s*\{\s*(#.*)?$`)
	rBlockEnd, _ := regexp.Compile(`^\s*\}\s*(#.*)?$`)
	line := c.lines[c.i]
	lineStart := c.i
	res := rBlock.FindStringSubmatch(line)
	if len(res) > 0 {
		var values []string
		entry.name = res[1]
		entry.comment = res[2]
		for c.i < len(c.lines) {
			c.i++
			line = c.lines[c.i]
			res = rBlockEnd.FindStringSubmatch(line)
			if len(res) > 0 {
				entry.value = strings.Join(values, "\n")
				entry.pos = [2]int{lineStart, c.i}
				break
			} else {
				values = append(values, line) // FIXME Trim? maybe no
			}
		}
		if c.i >= len(c.lines) {
			return 0 // FIXME throw error.
		}
	} else {
		res = rSingle.FindStringSubmatch(line)

		if len(res) > 0 {
			entry.pos = [2]int{lineStart, lineStart}
			entry.name = res[1]
			entry.value = res[2]
			entry.comment = res[3]
		}
	}

	if entry.name != "" {
		c.variables[entry.name] = &entry //FIXME: Sure?
		c.addEntry(&entry)
		return 1 //FIXME: return number of processed lines
	}
	return 0 // FIXME throw error
}

func (e entryVariable) GetSshConfig(c *PConfig) string {
	return "# " + strings.Join(c.lines[e.pos[0]:e.pos[1]+1], "\n# ") + "\n"
}
