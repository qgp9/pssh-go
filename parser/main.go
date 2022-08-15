package parser

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
)

type parseable interface {
	Selector(string, string) bool
	Parse(c *PConfig) int
}

type entryer interface {
	GetSshConfig(c *PConfig) string
}

func parseEntry(e parseable, c *PConfig) int {
	ii := e.Parse(c)
	if ii < 0 {
		log.Printf("%T: %s", e, c.lines[c.i])
		return c.i
	}
	return ii
}

func ParsePConfigFromFile(path string) *PConfig {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return ParsePConfigString(string(content))
}

func ParsePConfigString(str string) *PConfig {
	scanner := bufio.NewScanner(strings.NewReader(str))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return ParsePConfigStringSlice(lines)

}

func ParsePConfigStringSlice(lines []string) *PConfig {
	p := NewPConfig()
	p.lines = lines
	parsers := []parseable{
		&parserEmpty{},    // ''
		&parserOption{},   // begins with -
		&parserNode{},     // begins with |
		&parserComment{},  // begins with #
		&parserVariable{}, // begins with $
		&parserIgnore{},   // begins with +
		&parserGroup{},    // begins with @@
		&parserUnknown{},  // all
	}
	for p.i = 0; p.i < len(lines); p.i++ {
		// p.i may be going to modified by each parser.
		line := lines[p.i]
		l := strings.TrimSpace(line)
		for _, parser := range parsers {
			if parser.Selector(line, l) {
				parser.Parse(p)
				break
			}
		}
	}
	return p
}
