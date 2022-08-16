package parser

import (
	"log"
	"reflect"
	"regexp"
)

// Helper to implement interface parceable
type parserHelper[T entryer] struct {
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
	entry := reflect.New(reflect.TypeOf(new(T)).Elem().Elem()).Interface().(T)
	entry.SetValue(c.currentLine())
	c.addEntry(entry)
	return 1
}

type entryBase struct {
	line  int
	value string
}

func (e entryBase) GetSshConfig(c *PConfig) string {
	return e.value
}

func (e *entryBase) SetValue(str string) {
	e.value = str
}

func NewParserWithSelector[P parseableHelper](selector string) P {
	p := reflect.New(reflect.TypeOf(new(P)).Elem().Elem()).Interface().(P)
	p.setRegexp(selector)
	return p
}

func NewParserWithSelector2[P parseableHelper](p P, selector string) P {
	p.setRegexp(selector)
	return p
}

func NewParser[P parseableHelper]() P {
	p := reflect.New(reflect.TypeOf(new(P)).Elem().Elem()).Interface().(P)
	s := p.getSelectorString()
	if s != "" {
		p.setRegexp(s)
	}
	return p
}
