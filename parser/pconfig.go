package parser

import (
	"bufio"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/qgp9/pssh-go/utils"
)

type PConfig struct {
	lines     []string
	i         int
	EntryList []entryer
	variables map[string]*entryVariable
	nodes     map[string]*entryNode
}

func (p *PConfig) LoadConfigFromFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	p.LoadConfigFromString(string(content))

}

func (p *PConfig) LoadConfigFromString(str string) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	p.Parse(lines)
}

func (p *PConfig) Parse(lines []string) {
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
}

func (p *PConfig) WriteSshConfig() {
	var sshConfigList []string
	for _, v := range p.EntryList {
		sshConfigList = append(sshConfigList, v.GetSshConfig(p))
	}
	err := utils.WriteStringToFile("output.txt", strings.Join(sshConfigList, "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

var reVar = regexp.MustCompile(`\$\w+`)

func (p *PConfig) applyVariable(str string) string {
	ret := strings.Clone(str)
	for reVar.MatchString(ret) {
		ret = reVar.ReplaceAllStringFunc(str, func(r string) string {
			rn := strings.Trim(r, "$")
			variable, ok := p.variables[rn]
			if ok != true {
				log.Fatal("Variable " + r + " doesn't exists!")
			}
			return variable.value
		})
	}
	return ret
}

func (p *PConfig) addEntry(e entryer) {
	p.EntryList = append(p.EntryList, e)
}

func (p *PConfig) currentLine() string {
	return p.lines[p.i]
}

func NewPConfig() *PConfig {
	p := new(PConfig)
	p.variables = make(map[string]*entryVariable)
	return p
}
