package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"cpatino.com/advent-of-code/2024/cmd/generate"
	"cpatino.com/advent-of-code/2024/cmd/submit"
	"github.com/spf13/cobra"

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
	10: day10.Run,
	11: day11.Run,
	2:  day2.Run,
	3:  day3.Run,
	4:  day4.Run,
	5:  day5.Run,
	6:  day6.Run,
	7:  day7.Run,
	8:  day8.Run,
	9:  day9.Run,
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
			file = fmt.Sprintf("./day%d/input.txt", day)
		}

		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Fatalf("file does not exist: %s", file)
		}

		log.Printf("Day: %d, Input File: %s\n", day, file)

		part1, part2 := days[day](file)

		log.Printf("Part 1: %d\n", part1)
		log.Printf("Part 2: %d\n", part2)
	},
}

func init() {
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(submit.SubmitCmd)
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
