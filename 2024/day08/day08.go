package day08

import (
	"bufio"
	"os"
)

type Point struct {
	x, y int
}

func isValid(p Point, rows, cols int) bool {
	return p.x >= 0 && p.x < rows && p.y >= 0 && p.y < cols
}

func Part1(grid map[rune][]Point, rows, cols int) int {
	antinodes := make(map[Point]bool)

	for _, antennas := range grid {
		if len(antennas) < 2 {
			continue
		}

		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1, a2 := antennas[i], antennas[j]
				dx, dy := a2.x-a1.x, a2.y-a1.y

				forward := Point{a1.x - dx, a1.y - dy}
				if isValid(forward, rows, cols) {
					antinodes[forward] = true
				}

				backward := Point{a2.x + dx, a2.y + dy}
				if isValid(backward, rows, cols) {
					antinodes[backward] = true
				}
			}
		}
	}

	return len(antinodes)
}

func Part2(grid map[rune][]Point, rows, cols int) int {
	antinodes := make(map[Point]bool)
	for _, antennas := range grid {
		if len(antennas) < 2 {
			continue
		}

		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1, a2 := antennas[i], antennas[j]
				dx, dy := a2.x-a1.x, a2.y-a1.y

				antinodes[a1] = true
				antinodes[a2] = true

				forward := Point{a1.x - dx, a1.y - dy}
				for isValid(forward, rows, cols) {
					antinodes[forward] = true
					forward = Point{forward.x - dx, forward.y - dy}
				}

				backward := Point{a2.x + dx, a2.y + dy}
				for isValid(backward, rows, cols) {
					antinodes[backward] = true
					backward = Point{backward.x + dx, backward.y + dy}
				}
			}
		}
	}

	return len(antinodes)
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make(map[rune][]Point)
	scanner := bufio.NewScanner(file)
	rows, cols := 0, 0
	for scanner.Scan() {
		line := scanner.Text()

		for c, ch := range line {
			if ch != '.' {
				grid[ch] = append(grid[ch], Point{rows, c})
			}
		}

		cols = len(line)
		rows++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(grid, rows, cols)
	part2 := Part2(grid, rows, cols)

	return part1, part2
}
