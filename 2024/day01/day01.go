package day01

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1(firstCol []int, secondCol []int) int {
	slices.Sort(firstCol)
	slices.Sort(secondCol)

	sum := 0.0
	for i := 0; i < len(firstCol); i++ {
		sum += math.Abs(float64(firstCol[i] - secondCol[i]))
	}

	return int(sum)
}

func Part2(firstCol []int, secondCol []int) int {
	counts := make(map[int]int)
	for _, num := range secondCol {
		counts[num]++
	}

	sum := 0
	for _, num := range firstCol {
		sum += num * counts[num]
	}

	return sum
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var firstCol []int
	var secondCol []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		var input []int
		for _, word := range words {
			word, err := strconv.Atoi(word)
			if err != nil {
				panic(err)
			}

			input = append(input, word)
		}

		firstCol = append(firstCol, input[0])
		secondCol = append(secondCol, input[1])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(firstCol, secondCol)
	part2 := Part2(firstCol, secondCol)

	return part1, part2
}
