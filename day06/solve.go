package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func partOne(data string) {

	data = strings.TrimSpace(data)
	seq := strings.SplitSeq(data, "\n")

	var group [][]string

	for l := range seq {
		group = append(group, strings.Fields(l))
	}

	operatorRowLen := len(group)
	colLen := len(group[operatorRowLen-1])
	sums := make([]int, colLen)

	for r := range operatorRowLen - 1 {
		for c := range colLen {
			num, _ := strconv.Atoi(group[r][c])
			if r == 0 {
				sums[c] += num
			} else {
				switch group[operatorRowLen-1][c] {
				case "*":
					sums[c] *= num
				case "+":
					sums[c] += num
				}
			}
		}
	}

	sum := 0
	for _, v := range sums {
		sum += v
	}

	// fmt.Println(group)
	fmt.Println(sum)
}

func getNumber(input []string, char, rowLen int) string {
	item := ""

	for i := range rowLen {
		item += string(input[i][char])
	}
	item = strings.TrimSpace(item)

	return item

}

func partTwo(data string) {
	seq := strings.Split(data, "\n")
	str := seq[len(seq)-1]

	// this is to get the width of each col as whitespace varies
	re := regexp.MustCompile(`([*+]\s*)`)
	lens := re.FindAllString(str, -1)
	colLen := len(lens)

	group := make([][]string, len(seq))
	sums := make([]int, colLen)

	for i := range len(seq) {
		group[i] = make([]string, colLen)
	}

	operators := strings.Fields(seq[len(seq)-1])

	char := len(seq[0]) - 1
	for col := range colLen {
		for row := range len(seq) {
			var word string
			var newCol bool
			for {
				if char < 0 {
					break
				}
				word = getNumber(seq, char, len(seq)-1)
				char -= 1
				if word != "" {
					break
				} else {
					newCol = true
					break
				}

			}

			if newCol == true {
				break
			}
			group[row][col] += word
		}
	}

	// solve
	slices.Reverse(operators)

	for c := range len(lens) {
		for r := range len(group) {
			num, _ := strconv.Atoi(group[r][c])
			if num == 0 {
				continue
			}
			if r == 0 {
				sums[c] += num
			} else {
				switch operators[c] {
				case "*":

					sums[c] *= num
				case "+":
					sums[c] += num
				}

			}
		}
	}
	sum := 0
	for _, v := range sums {
		sum += v
	}
	fmt.Println(sum)

}

func main() {
	// data := aoctools.LoadFile("./day06/sample.txt")
	data := aoctools.LoadFile("./day06/input.txt")
	partOne(data)
	partTwo(data)

}
