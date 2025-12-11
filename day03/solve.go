package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func getLargetNumber(r []rune) int {
	maxFirstDigit := 0
	bestCombo := 0

	for i := range r {
		combo := maxFirstDigit*10 + int(r[i]) - 48
		bestCombo = max(bestCombo, combo)
		maxFirstDigit = max(maxFirstDigit, int(r[i])-48)
	}
	return bestCombo
}

func getLargetNumberGreedy(r []rune, start int, end int) (int, int) {
	var largest int = 0
	var nextIndex int = 0

	for i := start; i < len(r)-end; i++ {
		if ch := int(r[i]) - 48; largest < ch {
			largest = ch
			nextIndex = i + 1
		}

	}
	return largest, nextIndex
}

func partOne(data string) {
	str := strings.SplitSeq(data, "\n")
	sum := 0

	for s := range str {
		if len(s) == 0 {
			continue
		}

		runes := []rune(s)
		largest := getLargetNumber(runes)

		sum += largest

	}

	fmt.Println("part one", sum)

}

func partTwo(data string) {
	str := strings.SplitSeq(data, "\n")
	sum := 0

	for s := range str {
		if len(s) == 0 {
			continue
		}

		largest := ""
		nextI := 0
		num := 0
		runes := []rune(s)

		for i := 11; i >= 0; i-- {
			num, nextI = getLargetNumberGreedy(runes, nextI, i)
			largest += strconv.Itoa(num)

		}
		num, _ = strconv.Atoi(largest)
		sum += num

	}

	fmt.Println("part two", sum)

}

func main() {
	// data := aoctools.LoadFile("./day03/sample.txt")
	data := aoctools.LoadFile("./day03/input.txt")
	partOne(data)
	partTwo(data)
}
