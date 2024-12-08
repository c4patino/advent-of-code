package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	if len(flag.Args()) != 2 {
		panic("Usage: go run . <day> <input file>")
	}

	day, inputFile := flag.Args()[0], flag.Args()[1]

	if _, err := strconv.Atoi(day); err != nil {
		panic("Please provide a valid numeric day")
	}

	fmt.Printf("Day: %s, Input File: %s\n", day, inputFile)

	var part1, part2 int
	switch day {
	case "1":
		part1, part2 = day1.Run(inputFile)
	case "2":
		part1, part2 = day2.Run(inputFile)
	case "3":
		part1, part2 = day3.Run(inputFile)
	case "4":
		part1, part2 = day4.Run(inputFile)
	case "5":
		part1, part2 = day5.Run(inputFile)
	case "6":
		part1, part2 = day6.Run(inputFile)
	case "7":
		part1, part2 = day7.Run(inputFile)
	case "8":
		part1, part2 = day8.Run(inputFile)
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
