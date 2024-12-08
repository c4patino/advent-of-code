package day2

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSafe(row []int) bool {
	increasing := true
	decreasing := true
	valid := true

	for i := 1; i < len(row); i++ {
		diff := math.Abs(float64(row[i]) - float64(row[i-1]))

		if diff < 1 || diff > 3 {
			valid = false
			break
		}

		if row[i] < row[i-1] {
			increasing = false
		}

		if row[i] > row[i-1] {
			decreasing = false
		}
	}

	if valid && (increasing || decreasing) {
		return true
	}

	return false
}

func Part1(rows [][]int) int {
	safe := 0
	for _, row := range rows {
		if isSafe(row) {
			safe += 1
		}
	}

	return safe
}

func Part2(rows [][]int) int {
	safe := 0
	for _, row := range rows {
		if isSafe(row) {
			safe += 1
			continue
		}

		for i := 0; i < len(row); i++ {
			modified := make([]int, len(row)-1)
			copy(modified[:i], row[:i])
			copy(modified[i:], row[i+1:])

			if isSafe(modified) {
				safe += 1
				break
			}
		}
	}

	return safe
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rows [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var row []int

		for _, num := range strings.Fields(line) {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			row = append(row, n)
		}

		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(rows)
	part2 := Part2(rows)

	return part1, part2
}
