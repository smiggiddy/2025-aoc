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

	str := strings.Trim(string(data), "\n")

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

func SetupDay(path string) {

	err := os.MkdirAll(path, 0750)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path+"/input.txt", []byte(""), 0640)
	err = os.WriteFile(path+"/sample.txt", []byte(""), 0640)
	err = os.WriteFile(path+"/solve.go", []byte(""), 0640)

}
