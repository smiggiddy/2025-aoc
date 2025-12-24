package main

import (
	"fmt"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func checkPresence(i map[int]bool, key int) bool {
	_, check := i[key]

	return check == true
}
func checkPresenceInt(i map[int]int, key int) bool {
	_, check := i[key]

	return check == true
}

func partOne(data string) {
	text := strings.SplitSeq(data, "\n")

	positions := make(map[int]bool)

	splitCount := 0
	for line := range text {
		for i, c := range line {
			if c == 'S' {
				positions[i] = true
				continue
			}

			if c == '^' && positions[i] == true {
				delete(positions, i)

				newPosMin := max(0, i-1)
				newPostMax := min(len(line)-1, i+1)

				if !checkPresence(positions, newPosMin) {
					positions[newPosMin] = true
				}

				if !checkPresence(positions, newPostMax) {
					positions[newPostMax] = true
				}
				splitCount += 1
			}

		}
	}
	fmt.Println("Arrow Count:", splitCount)

}
func partTwo(data string) {
	text := strings.SplitSeq(data, "\n")

	positions := make(map[int]int)

	for line := range text {
		for i, c := range line {
			if c == 'S' {
				positions[i] = 1
				continue
			}

			if c == '^' && positions[i] == 1 {
				delete(positions, i)

				newPosMin := max(0, i-1)
				newPostMax := min(len(line)-1, i+1)

				if !checkPresenceInt(positions, newPosMin) {
					positions[newPosMin] += 1
				}

				if !checkPresenceInt(positions, newPostMax) {
					positions[newPostMax] += 1
				}
			}

		}
	}
	fmt.Println("Arrow Count:", positions)

}

func main() {
	// data := aoctools.LoadFile("./day07/sample.txt")
	data := aoctools.LoadFile("./day07/input.txt")
	partOne(data)
	partTwo(data)
}
