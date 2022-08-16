package parser

import (
	"log"
	"reflect"
	"regexp"
)

// Helper to implement interface parceable
type parser[T entryer] struct {
	regstr *regexp.Regexp
}

func (e *parser[T]) setRegexp(regStr string) (err error) {
	e.regstr, err = regexp.Compile(regStr)
	if err != nil {
		log.Panic(err)
	}
	return err
}

func (e *parser[T]) Selector(line string) bool {
	if e.regstr != nil {
		if true == e.regstr.MatchString(line) {
			//log.Printf("Selector %T: %s", new(T), line)
			return true
		}
	}
	return false
}

func (e *parser[T]) Parse(c *PConfig) int {
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

func NewParserWithSelector[P parseable](selector string) P {
	p := reflect.New(reflect.TypeOf(new(P)).Elem().Elem()).Interface().(P)
	p.setRegexp(selector)
	return p
}

func NewParser[P parseable]() P {
	p := reflect.New(reflect.TypeOf(new(P)).Elem().Elem()).Interface().(P)
	return p
}
