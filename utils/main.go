package utils

import (
	"log"
	"os"
	"strings"
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
