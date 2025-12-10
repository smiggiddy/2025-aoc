package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func main() {
	pathPtr := flag.String("day", "", "Creates a dir and starter files at pwd")

	flag.Parse()

	if len(*pathPtr) == 0 {
		fmt.Println("Please include a path")
		os.Exit(1)

	}

	aoctools.SetupDay(*pathPtr)

}
