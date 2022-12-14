package parser

import (
	"regexp"
	"strings"

	"github.com/qgp9/pssh-go/utils"
)

type parserNode struct {
	parserHelper[*entryNode]
}

func NewParserNode() *parserNode {
	return NewParserWithSelector[*parserNode](`^\s*\|`)
}

var reNode = regexp.MustCompile(`^|\s*\s*|\s*|\s+`)

func (e *parserNode) Parse(c *PConfig) int {
	line := c.lines[c.i]
	var dummy, address, option string
	var host, key, user, hostname, port, comment string
	utils.SplitAssign(line, "|", &dummy, &host, &address, &key, &option)
	utils.SplitAssign(address, "@", &user, &hostname)
	if hostname == "" {
		hostname, user = user, ""
	}
	utils.SplitAssign(hostname, ":", &hostname, &port)
	utils.SplitAssign(option, "#", &option, &comment)
	if host == "" {
		host = hostname
		hostname = ""
	}
	var entry = entryNode{
		Host:     host,
		User:     user,
		Hostname: hostname,
		Port:     port,
		Key:      key,
		Comment:  comment,
		Option:   option,
	}
	entry.pos[0] = c.i
	if utils.IsUpper(host) == false {
		c.addEntry(&entry)
		c.addNode(&entry)
	} else {
		entryG := entryGroup{}
		entryG.entryNode = entry
		entryG.SetValue(line)
		c.addEntry(&entryG)
		c.addGroup(&entryG)
	}

	return 1
}

type entryNode struct {
	Name     string
	Host     string
	User     string
	Hostname string
	Port     string
	Key      string
	Option   string
	Comment  string
	entryHelper
}

func (e *entryNode) GetSshConfig(c *PConfig) string {
	var strs = []string{""}
	addEntry := func(value, prefix string) {
		if value != "" {
			value = c.applyVariable(value)
			strs = append(strs, prefix+value)
		}
	}
	addEntry(e.Comment, "# ")
	addEntry(e.Host, "Host ")
	addEntry(e.Hostname, "HostName ")
	addEntry(e.User, "User ")
	addEntry(e.Port, "Port ")
	addEntry(e.Key, "IdentityFile ")
	addEntry(e.Option, "")

	return strings.Join(strs, "\n") + "\n"
}
