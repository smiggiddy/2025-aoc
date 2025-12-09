package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func partOne(input string) int {
	lastPosition := 50
	var sum int

	for line := range strings.SplitSeq(input, "\n") {

		if len(line) == 0 {
			continue
		}

		movement, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			lastPosition -= movement

		} else {
			lastPosition += movement
		}

		if lastPosition%100 == 0 {
			sum += 1
		}

	}
	return sum
}

func partTwo(input string) int {
	dial := 50
	ret := 0

	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			continue
		}
		num, _ := strconv.Atoi(line[1:])

		if line[0] == 'L' {
			num = -num
			if dial+num <= 0 {
				start := 0
				if dial != 0 {
					start = 1
				}
				sweeps := start + abs((dial+num)/100)
				ret += sweeps
			}
		} else {
			ret += (dial + num) / 100
		}

		dial = ((dial+num)%100 + 100) % 100
	}

	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func partTwo(input string) int {
// 	lastPosition := 50
// 	var sum int = 0
//
// 	for line := range strings.SplitSeq(input, "\n") {
// 		if len(line) == 0 {
// 			continue
// 		}
//
// 		movement, _ := strconv.Atoi(line[1:])
// 		startPos := lastPosition
// 		if line[0] == 'L' {
// 			if movement > startPos {
// 				crossings := math.Floor(float64((movement - startPos + 99) / 100))
// 				sum += int(crossings)
// 				lastPosition = ((movement - startPos) % 100)
// 			} else {
// 				lastPosition -= movement
// 				lastPosition = lastPosition % 100
//
// 			}
//
// 		} else {
// 			if movement > startPos {
// 				crossings := math.Floor(float64((startPos + movement) / 100))
// 				sum += int(crossings)
// 				lastPosition = ((startPos + movement) % 100)
// 			} else {
// 				lastPosition += movement
//
// 			}
// 		}
// 		lastPosition = lastPosition % 100
// 		if lastPosition == 0 {
// 			sum += 1
// 		}
//
// 	}
// 	return sum
//
// }

func main() {
	input := aoctools.LoadFile("./day01/input.txt")
	sum := partOne(input)
	fmt.Println(sum)
	fmt.Println(partTwo(input))
}
