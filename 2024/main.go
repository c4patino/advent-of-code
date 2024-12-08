package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"cpatino.com/advent-of-code/2024/day1"
	"cpatino.com/advent-of-code/2024/day2"
	"cpatino.com/advent-of-code/2024/day3"
	"cpatino.com/advent-of-code/2024/day4"
	"cpatino.com/advent-of-code/2024/day5"
	"cpatino.com/advent-of-code/2024/day6"
	"cpatino.com/advent-of-code/2024/day7"
	"cpatino.com/advent-of-code/2024/day8"
)

func main() {
	var fileArg = flag.String("file", "", "input filename")
	flag.Parse()

	args := flag.Args()
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: go run . [OPTIONS] <day>")
		os.Exit(1)
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	var inputFile = fmt.Sprintf("day%d/input.txt", day)
	if *fileArg != "" {
		inputFile = *fileArg
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Println("test file does not exist: %v", inputFile)
		os.Exit(1)
	}

	fmt.Printf("Day: %d, Input File: %s\n", day, inputFile)

	var part1, part2 int
	switch day {
	case 1:
		part1, part2 = day1.Run(inputFile)
	case 2:
		part1, part2 = day2.Run(inputFile)
	case 3:
		part1, part2 = day3.Run(inputFile)
	case 4:
		part1, part2 = day4.Run(inputFile)
	case 5:
		part1, part2 = day5.Run(inputFile)
	case 6:
		part1, part2 = day6.Run(inputFile)
	case 7:
		part1, part2 = day7.Run(inputFile)
	case 8:
		part1, part2 = day8.Run(inputFile)
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
