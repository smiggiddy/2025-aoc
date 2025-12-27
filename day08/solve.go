package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

type battery struct {
	x, y, z   int
	inCircuit bool
}

type circuit struct {
	batteries []battery
}

type comparisons struct {
	b1, b2   battery
	distance int
}

func getDistance(b1, b2 battery) int {
	x := float64((b1.x - b2.x) * (b1.x - b2.x))
	y := float64((b1.y - b2.y) * (b1.y - b2.y))
	z := float64((b1.z - b2.z) * (b1.z - b2.z))

	distance := math.Sqrt(x + y + z)

	return int(distance)

}

func getBatts(d string) (batts []battery) {
	lines := strings.SplitSeq(d, "\n")

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
		batts = append(batts, batt)
	}
	return

}

func main() {
	data := aoctools.LoadFile("./day08/sample.txt")
	batts := getBatts(data)

	var distances []comparisons
	visited := make(map[int]bool)

	for r := range len(batts) {
		dist := 0
		for i := range batts {
			if i == r {
				continue
			}

			d := getDistance(batts[r], batts[i])
			c := comparisons{
				b1:       batts[r],
				b2:       batts[i],
				distance: d,
			}
			if _, ok := visited[d]; !ok {
				distances = append(distances, c)
				visited[d] = true
			}
			if dist == 0 {
				dist = d
			}
			dist = min(d, dist)
		}

	}

	slices.SortFunc(distances, func(a, b comparisons) int {
		return cmp.Compare(a.distance, b.distance)
	})
	fmt.Println(distances)

}
