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

	var outputFile, printFile bool
	getInputCmd := flag.NewFlagSet("getinput", flag.ExitOnError)
	getInputDay := getInputCmd.Int("day", 0, "AOC Day")
	getInputYear := getInputCmd.Int("year", 0, "AOC Year")
	getInputCmd.BoolVar(&printFile, "print", false, "Print the input to stdout")
	getInputCmd.BoolVar(&printFile, "p", false, "Print the input to stdout")
	getInputCmd.BoolVar(&outputFile, "output", false, "Outputs the input to a dir defined by day")
	getInputCmd.BoolVar(&outputFile, "o", false, "Outputs the input to a dir defined by day")

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
			input, err := aoctools.GetInput(*getInputDay, *getInputYear)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v", err)
				os.Exit(1)
			}

			if printFile {
				fmt.Println(string(input))
			}

			if outputFile {
				dayString := aoctools.FormatDayString(*getInputDay)
				filePath := fmt.Sprintf("day%s/input.txt", dayString)

				err = os.WriteFile(filePath, input, 0644)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}

			}

		}
	default:
		fmt.Fprintf(os.Stderr, "Expected 'setup' or 'input' subcommands")
		os.Exit(1)

	}

}
