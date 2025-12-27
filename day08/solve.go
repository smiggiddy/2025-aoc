package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

type battery struct {
	x, y, z int
}

func getDistance(b1, b2 battery) int {
	x := float64((b1.x - b2.x) * (b1.x - b2.x))
	y := float64((b1.y - b2.y) * (b1.y - b2.y))
	z := float64((b1.z - b2.z) * (b1.z - b2.z))

	distance := math.Sqrt(x + y + z)

	return int(distance)

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

	for r := range len(batteries) {
		distance := 0
		for i := range batteries {
			if i == r {
				continue
			}

			d := getDistance(batteries[r], batteries[i])
			if distance == 0 {
				distance = d
			}
			fmt.Printf("Distance=%v, batt1=%v, batt2=%v\n", d, batteries[r], batteries[i])

			distance = min(d, distance)
		}
		fmt.Printf("Smallest Distance=%v\n\n", distance)

	}

}
