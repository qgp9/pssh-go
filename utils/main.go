package utils

import (
	"log"
	"os"
	"reflect"
	"strings"
	"unicode"
)

func SplitAndTrim(s string, sep string) []string {
	return TrimSpaceSlice(strings.Split(s, sep))
}

func TrimSpaceSlice(s []string) []string {
	trimed := make([]string, 0, len(s))
	for _, v := range s {
		trimed = append(trimed, strings.TrimSpace(v))
	}
	return trimed
}

func SplitAssign(str string, sep string, args ...*string) {
	fields := strings.Split(str+strings.Repeat(sep, len(args)), sep)
	for i, v := range args {
		*v = strings.TrimSpace(fields[i])
	}
}

func WriteStringToFile(path string, s string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	_, err = f.WriteString(s)
	if err != nil {
		log.Fatal(err)
		return err
	}
	f.Sync()

	return nil
}

func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) == false {
			return false
		}
	}
	return true
}

func AddToMap[V any](m map[string]V, key string, value V, log func()) bool {
	_, ok := m[key]
	if ok == true {
		log()
		return false
	}
	m[key] = value
	return true
}

func NewElem[T any]() T {
	return reflect.New(reflect.TypeOf(new(T)).Elem().Elem()).Interface().(T)
}
