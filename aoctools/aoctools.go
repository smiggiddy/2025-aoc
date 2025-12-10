package aoctools

import (
	"os"
	"strings"
)

func LoadFile(file string) string {
	data, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	str := strings.Trim(string(data), "")

	return str
}

func AllSame(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, ch := range s {
		if ch != rune(s[0]) {
			return false
		}
	}

	return true
}
