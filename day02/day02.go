package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func solveOne(str iter.Seq[string]) int {
	// math log 10 get the digits

	total := 0
	for s := range str {
		if len(s) == 0 {
			continue
		}
		nums := strings.Split(s, "-")
		start, _ := strconv.Atoi(strings.Trim(nums[0], "\n"))
		end, _ := strconv.Atoi(strings.Trim(nums[1], "\n"))

		for i := start; i < end+1; i++ {
			if len(strconv.Itoa(i))%2 == 0 {

				sn := strconv.Itoa(i)
				l := len(sn) / 2

				if strings.HasPrefix(sn, "0") {
					continue
				}
				// check if all numbers are the same
				if aoctools.AllSame(sn) {
					total += i
					continue
				}

				// check if number repeats
				if l > 1 {
					if strings.Compare(sn[:l], sn[l:]) == 0 {
						total += i
					}
				}
			}
		}
	}
	return total
}

func solveTwo(str iter.Seq[string]) int {
	total := 0
	for s := range str {
		if len(s) == 0 {
			continue
		}
		nums := strings.Split(s, "-")
		start, _ := strconv.Atoi(strings.Trim(nums[0], "\n"))
		end, _ := strconv.Atoi(strings.Trim(nums[1], "\n"))

		for i := start; i < end+1; i++ {
			sn := strconv.Itoa(i)

			// fail quick
			if strings.HasPrefix(sn, "0") || len(sn) < 2 {
				continue
			}

			// check if all numbers are the same
			if aoctools.AllSame(sn) {
				total += i
				continue
			}

			if checkSequences(sn) {
				total += i
			}

		}
	}
	return total
}

func checkSequences(s string) bool {

	count := 2

	for {

		if count > len(s)/2 {
			return false
		}
		var found bool = true
		for i := count; i < len(s); i += count {

			if i+count > len(s) {
				found = false
				break
			}
			if s[:count] == s[i:i+count] {
				continue
			} else if len(s)%count != 0 {
				found = false
				break
			} else {
				found = false
				break
			}
		}

		if found == true {
			return true
		}
		count += 1
	}

}

func main() {
	// data := aoctools.LoadFile("day02/sample_data.txt")
	data := aoctools.LoadFile("day02/input.txt")
	s := strings.SplitSeq(data, ",")
	partOne := solveOne(s)
	sTwo := strings.SplitSeq(data, ",")
	partTwo := solveTwo(sTwo)

	fmt.Printf("Part One=%v\n", partOne)
	fmt.Printf("Part Two=%v\n", partTwo)
}
