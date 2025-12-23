package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/smiggiddy/2025-aoc/aoctools"
)

func main() {
	setupCmd := flag.NewFlagSet("setup", flag.ExitOnError)
	setupDay := setupCmd.String("day", "", "Creates a dir and starter files at pwd")

	getInputCmd := flag.NewFlagSet("getinput", flag.ExitOnError)
	getInputDay := getInputCmd.Int("day", 0, "AOC Day")
	getInputYear := getInputCmd.Int("year", 0, "AOC Year")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'setup' or 'getinput' subcommands.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "setup":
		setupCmd.Parse(os.Args[2:])
		if len(*setupDay) == 0 {
			fmt.Println("Please specify which day to create directory and files.")
			os.Exit(1)

		}

		if *setupDay != "" {
			aoctools.SetupDay(*setupDay)
		}

	case "getinput":
		{
			getInputCmd.Parse(os.Args[2:])
			aoctools.GetInput(*getInputDay, *getInputYear)

		}
	default:
		fmt.Println("Expected 'setup' or 'input' subcommands")
		os.Exit(1)

	}

}
