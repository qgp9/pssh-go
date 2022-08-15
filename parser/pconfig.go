package parser

import (
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
