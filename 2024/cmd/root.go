package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"cpatino.com/advent-of-code/2024/cmd/generate"
	"github.com/spf13/cobra"

	"cpatino.com/advent-of-code/2024/day01"
	"cpatino.com/advent-of-code/2024/day02"
	"cpatino.com/advent-of-code/2024/day03"
	"cpatino.com/advent-of-code/2024/day04"
	"cpatino.com/advent-of-code/2024/day05"
	"cpatino.com/advent-of-code/2024/day06"
	"cpatino.com/advent-of-code/2024/day07"
	"cpatino.com/advent-of-code/2024/day08"
	"cpatino.com/advent-of-code/2024/day09"
	"cpatino.com/advent-of-code/2024/day10"
	"cpatino.com/advent-of-code/2024/day11"
	"cpatino.com/advent-of-code/2024/day12"
	"cpatino.com/advent-of-code/2024/day13"
	"cpatino.com/advent-of-code/2024/day14"
	"cpatino.com/advent-of-code/2024/day15"
	"cpatino.com/advent-of-code/2024/day16"
	"cpatino.com/advent-of-code/2024/day17"
	"cpatino.com/advent-of-code/2024/day18"
	"cpatino.com/advent-of-code/2024/day19"
	"cpatino.com/advent-of-code/2024/day20"
	"cpatino.com/advent-of-code/2024/day21"
)

var days = map[int]func(string) (interface{}, interface{}){
	1:  day01.Run,
	2:  day02.Run,
	3:  day03.Run,
	4:  day04.Run,
	5:  day05.Run,
	6:  day06.Run,
	7:  day07.Run,
	8:  day08.Run,
	9:  day09.Run,
	10: day10.Run,
	11: day11.Run,
	12: day12.Run,
	13: day13.Run,
	14: day14.Run,
	15: day15.Run,
	16: day16.Run,
	17: day17.Run,
	18: day18.Run,
	19: day19.Run,
	20: day20.Run,
	21: day21.Run,
}

var rootCmd = &cobra.Command{
	Use:   "advent-of-code",
	Short: "Advent of Code 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		file, _ := cmd.Flags().GetString("file")
		if file == "" {
			file = fmt.Sprintf("./day%02d/input.txt", day)
		}

		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Fatalf("file does not exist: %s", file)
		}

		log.Printf("Day: %d, Input File: %s\n", day, file)

		part1, part2 := days[day](file)

		log.Printf("Part 1: %v\n", part1)
		log.Printf("Part 2: %v\n", part2)
	},
}

func init() {
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.Flags().String("file", "", "Path to a file")
}

func Execute() error {
	if len(os.Args) == 1 {
		return rootCmd.Execute()
	}

	subCmd := os.Args[1]
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == subCmd {
			return cmd.Execute()
		}
	}

	return rootCmd.Execute()
}
