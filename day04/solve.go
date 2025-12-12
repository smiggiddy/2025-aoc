package main

import (
	"fmt"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func checkAdjacentPositions(x int, y int, l *[]string) bool {
	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {0, -1}, {1, 0}, {1, 1}, {1, -1}}
	adjPositions := 0
	length := len(*l)
	arr := *l

	for _, d := range directions {
		newX := d[0] + x
		newY := d[1] + y

		if newX > -1 && newY > -1 && newX < length && newY < length {
			if arr[newX][newY] == '@' {
				adjPositions += 1
			}
		}
	}

	return adjPositions < 4

}
func checkAdjacentPos(x int, y int, l *[][]byte) bool {
	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {0, -1}, {1, 0}, {1, 1}, {1, -1}}
	adjPositions := 0
	length := len(*l)
	arr := *l

	for _, d := range directions {
		newX := d[0] + x
		newY := d[1] + y

		if newX > -1 && newY > -1 && newX < length && newY < length {
			if arr[newX][newY] == '@' {
				adjPositions += 1
			}
		}
	}

	return adjPositions < 4

}

func partOne(input string) {
	lines := strings.Split(input, "\n")

	rowLen := len(lines[0])
	numOfRoles := 0

	for i := range len(lines) {
		for j := range rowLen {

			if lines[i][j] == '@' {
				if checkAdjacentPositions(i, j, &lines) {
					numOfRoles += 1
				}

			}

		}
	}
	fmt.Println(numOfRoles)
}

func partTwo(input string) {
	lines := strings.Split(input, "\n")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	rowLen := len(grid[0])
	numOfRoles := 0

	var madeChange bool = false
	for {
		madeChange = false
		for i := range len(grid) {
			for j := range rowLen {
				if grid[i][j] == '@' {
					if checkAdjacentPos(i, j, &grid) {
						numOfRoles += 1
						grid[i][j] = 'x'
						madeChange = true
					}
				}
			}
		}
		// separator := []byte("\n")
		// fmt.Println(string(bytes.Join(grid, separator)))
		if madeChange == false {
			break
		}

	}

	fmt.Println(numOfRoles)
}

func main() {
	// data := aoctools.LoadFile("./day04/sample.txt")
	data := aoctools.LoadFile("./day04/input.txt")
	partOne(data)
	partTwo(data)

}
