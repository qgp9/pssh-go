package parser

import (
	"log"
	"regexp"

	"github.com/qgp9/pssh-go/utils"
)

// Helper to implement interface parceable
type parserHelper[T entryable] struct {
	regstr *regexp.Regexp
}

func (e *parserHelper[T]) getSelectorString() string {
	return ""
}

func (e *parserHelper[T]) setRegexp(regStr string) (err error) {
	e.regstr, err = regexp.Compile(regStr)
	if err != nil {
		log.Panic(err)
	}
	return err
}

func (e *parserHelper[T]) Selector(line string) bool {
	if e.regstr != nil {
		if true == e.regstr.MatchString(line) {
			//log.Printf("Selector %T: %s", new(T), line)
			return true
		}
	}
	return false
}

func (e *parserHelper[T]) Parse(c *PConfig) int {
	entry := utils.NewElem[T]()
	entry.SetValue(c.currentLine())
	c.addEntry(entry)
	return 1
}

type entryHelper struct {
	pos   [2]int
	value string
}

func (e entryHelper) GetSshConfig(c *PConfig) string {
	return e.value + "\n"
}

func (e *entryHelper) SetValue(str string) {
	e.value = str
}

func NewParserWithSelector[P parseableHelper](selector string) P {
	p := utils.NewElem[P]()
	p.setRegexp(selector)
	return p
}

func NewParserWithSelector2[P parseableHelper](p P, selector string) P {
	p.setRegexp(selector)
	return p
}

func NewParser[P parseableHelper]() P {
	p := utils.NewElem[P]()
	s := p.getSelectorString()
	if s != "" {
		p.setRegexp(s)
	}
	return p
}
