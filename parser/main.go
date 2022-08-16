package parser

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
)

type parseable interface {
	Selector(string) bool
	Parse(c *PConfig) int
	setRegexp(string) error
}

type entryer interface {
	GetSshConfig(c *PConfig) string
	SetValue(string)
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
		NewParserEmpty(),
		NewParserEmpty(),    // ''
		NewParserOption(),   // begins with -
		NewParserNode(),     // begins with |
		NewParserComment(),  // begins with #
		NewParserVariable(), // begins with $
		NewParserIgnore(),   // begins with +
		NewParserGroup(),    // begins with @@
		NewParserUnknown(),  // all
	}
	for p.i = 0; p.i < len(lines); p.i++ {
		// p.i may be going to modified by each parser.
		line := lines[p.i]
		for _, parser := range parsers {
			if parser.Selector(line) == true {
				parser.Parse(p)
				break
			}
		}
	}
	return p
}
