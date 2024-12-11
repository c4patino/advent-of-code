package day11

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func update(stones map[int]int) map[int]int {
	processed := make(map[int]int)
	for stone, count := range stones {
		if stone == 0 {
			processed[1] += count
			continue
		}

		places := int(math.Log10(float64(stone))) + 1
		if places%2 == 0 {
			half := places / 2
			divisor := int(math.Pow10(half))

			left := stone / divisor
			right := stone % divisor

			processed[left] += count
			processed[right] += count
			continue
		}

		processed[stone*2024] += count
	}

	return processed
}

func Part1(stones map[int]int, steps int) int {
	for i := 0; i < steps; i++ {
		stones = update(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func Part2(stones map[int]int, steps int) int {
	return Part1(stones, steps)
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stones := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, c := range strings.Split(line, " ") {
			id, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}

			stones[id] += 1
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(stones, 25)
	part2 := Part2(stones, 75)

	return part1, part2
}
