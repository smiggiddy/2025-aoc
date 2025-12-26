package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

type battery struct {
	x, y, z int
}

func main() {
	data := aoctools.LoadFile("./day08/sample.txt")
	lines := strings.SplitSeq(data, "\n")

	var batteries []battery

	for l := range lines {
		str := strings.TrimSpace(l)
		s := strings.Split(str, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		z, _ := strconv.Atoi(s[2])
		batt := battery{
			x: x,
			y: y,
			z: z,
		}
		batteries = append(batteries, batt)
	}

	fmt.Println(batteries)

}
