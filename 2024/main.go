package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"cpatino.com/advent-of-code/2024/day1"
	"cpatino.com/advent-of-code/2024/day10"
	"cpatino.com/advent-of-code/2024/day11"
	"cpatino.com/advent-of-code/2024/day2"
	"cpatino.com/advent-of-code/2024/day3"
	"cpatino.com/advent-of-code/2024/day4"
	"cpatino.com/advent-of-code/2024/day5"
	"cpatino.com/advent-of-code/2024/day6"
	"cpatino.com/advent-of-code/2024/day7"
	"cpatino.com/advent-of-code/2024/day8"
	"cpatino.com/advent-of-code/2024/day9"
)

var days = map[int]func(string) (int, int){
	1:  day1.Run,
	2:  day2.Run,
	3:  day3.Run,
	4:  day4.Run,
	5:  day5.Run,
	6:  day6.Run,
	7:  day7.Run,
	8:  day8.Run,
	9:  day9.Run,
	10: day10.Run,
	11: day11.Run,
}

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
		fmt.Printf("test file does not exist: %v\n", inputFile)
		os.Exit(1)
	}

	fmt.Printf("Day: %d, Input File: %s\n", day, inputFile)

	part1, part2 := days[day](inputFile)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
