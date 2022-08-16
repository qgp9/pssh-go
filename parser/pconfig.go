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
	EntryList []entryable
	variables map[string]*entryVariable
	nodes     map[string]*entryNode
	groups    map[string]*entryGroup
}

func (p *PConfig) WriteSshConfig() {
	var sshConfigList []string
	for _, v := range p.EntryList {
		sshConfigList = append(sshConfigList, v.GetSshConfig(p))
	}
	err := utils.WriteStringToFile("output.txt", strings.Join(sshConfigList, ""))
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

func (p *PConfig) addEntry(e entryable) {
	p.EntryList = append(p.EntryList, e)
}

func (p *PConfig) addVariable(e *entryVariable) {
	_, ok := p.variables[e.name]
	if ok == true {
		log.Fatal("Variable " + e.name + " is duplicated!") // Better error handling
	} else {
		p.variables[e.name] = e
	}
}

func (p *PConfig) addNode(e *entryNode) {
	_, ok := p.variables[e.Host]
	if ok == true {
		log.Fatal("Node " + e.Host + " is duplicated!") // Better error handling
	} else {
		p.nodes[e.Host] = e
	}
}

func (p *PConfig) addGroup(e *entryGroup) {
	utils.AddToMap(p.groups, e.Host, e, func() {
		log.Panic("Group " + e.Host + " is duplicated!") // Better error handling
	})
}

func (p *PConfig) currentLine() string {
	return p.lines[p.i]
}

func NewPConfig() *PConfig {
	p := new(PConfig)
	p.variables = make(map[string]*entryVariable)
	p.nodes = make(map[string]*entryNode)
	p.groups = make(map[string]*entryGroup)

	return p
}
