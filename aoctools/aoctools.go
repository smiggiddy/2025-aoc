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
