package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func main() {
	// data := aoctools.LoadFile("./day03/sample.txt")
	data := aoctools.LoadFile("./day03/input.txt")
	str := strings.SplitSeq(data, "\n")
	sum := 0
	var num int

	for s := range str {
		var largest int = 0

		for i := 0; i <= len(s)-2; i++ {
			for j := 1; j <= len(s)-1; j++ {
				if j > i {
					tempStr := string(s[i]) + string(s[j])
					num, _ = strconv.Atoi(tempStr)

					if num > largest {
						largest = num
					}

				}

			}

		}
		sum += largest

	}

	fmt.Println(sum)

}
