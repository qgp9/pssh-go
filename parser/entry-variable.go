package parser

import (
	"log"
	"regexp"
	"strings"
)

type parserVariable parserBase

type entryVariable struct {
	entryBase
	name    string
	value   string
	comment string
}

func (e *parserVariable) Selector(line string, trimed string) bool {
	return strings.HasPrefix(trimed, "$")
}

func (e *parserVariable) Parse(c *PConfig) int {
	log.Println("entryVariable: ", c.lines[c.i])
	var entry = entryVariable{}
	rSingle, _ := regexp.Compile(`^\$([\w]+)\s*=\s*([^#]+)(.*)$`)
	rBlock, _ := regexp.Compile(`^\$([\w]+)\s*=\s*\{\s*(#.*)?$`)
	rBlockEnd, _ := regexp.Compile(`^\s*\}\s*(#.*)?$`)
	line := c.lines[c.i]
	//rBlock.MatchString(line)
	res := rBlock.FindStringSubmatch(line)
	if len(res) > 0 {
		log.Println("DEBUG===: ", line)
		var values []string
		entry.name = res[1]
		entry.comment = res[2]
		for c.i < len(c.lines) {
			c.i++
			line = c.lines[c.i]
			res = rBlockEnd.FindStringSubmatch(line)
			if len(res) > 0 {
				entry.value = strings.Join(values, "\n")
				log.Println(e)
				break // FIXME return number of processed lines
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
