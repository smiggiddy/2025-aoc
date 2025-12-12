package main

import (
	"fmt"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func main() {
	data := aoctools.LoadFile("./day05/sample.txt")
	seq := strings.Split(data, "\n\n")
	var groups, numRange []string

	// var found bool
	for _, c := range seq {
		fmt.Println(c)
		// if i == "" {
		// 	found = true
		// 	continue
		// }
		// if !found {
		// 	groups = append(groups, i)
		// } else {
		// 	numRange = append(numRange, i)
		// }
	}

	fmt.Println(groups)
	fmt.Println(numRange)
}
