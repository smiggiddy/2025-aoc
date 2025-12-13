package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

type ranges struct {
	start int
	end   int
}

func sortGroup(s string) []ranges {
	allRanges := []ranges{}
	groups := strings.SplitSeq(s, "\n")

	for group := range groups {
		g := strings.Split(group, "-")
		start, _ := strconv.Atoi(g[0])
		end, _ := strconv.Atoi(g[1])
		r := ranges{
			start: start,
			end:   end,
		}
		allRanges = append(allRanges, r)
	}

	return allRanges

}

func partOne(data string) {
	data = strings.TrimSpace(data)
	seq := strings.Split(data, "\n\n")
	g := seq[0]
	n := seq[1]
	groups := sortGroup(g)
	numbers := strings.Split(n, "\n")
	slices.Sort(numbers)
	var sum int = 0

	for _, s := range numbers {
		target, _ := strconv.Atoi(s)
		for _, group := range groups {
			if group.start <= target && target <= group.end {
				sum += 1
				break
			}
		}
	}

	fmt.Println(sum)

}

func main() {
	data := aoctools.LoadFile("./day05/input.txt")
	// data := aoctools.LoadFile("./day05/sample.txt")
	partOne(data)

}
