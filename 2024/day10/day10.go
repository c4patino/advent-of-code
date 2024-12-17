package day10

import (
	"bufio"
	"os"
)

type Direction string
type Location struct{ Y, X int }

const (
	NORTH Direction = "N"
	SOUTH Direction = "S"
	EAST  Direction = "E"
	WEST  Direction = "W"
)

var transformations = map[Direction]func(loc Location) Location{
	NORTH: func(loc Location) Location { return Location{X: loc.X, Y: loc.Y - 1} },
	SOUTH: func(loc Location) Location { return Location{X: loc.X, Y: loc.Y + 1} },
	EAST:  func(loc Location) Location { return Location{X: loc.X + 1, Y: loc.Y} },
	WEST:  func(loc Location) Location { return Location{X: loc.X - 1, Y: loc.Y} },
}

func applyTransform(loc Location, direction Direction, grid [][]int) (Location, bool) {
	newLoc := transformations[direction](loc)
	if newLoc.X < 0 || newLoc.X >= len(grid[0]) || newLoc.Y < 0 || newLoc.Y >= len(grid) {
		return loc, false
	}

	return newLoc, true
}

func findTrailEndings(loc Location, elevation int, grid [][]int, found map[Location]bool) int {
	if _, ok := found[loc]; !ok && grid[loc.Y][loc.X] == 9 {
		found[loc] = true
		return 1
	}

	sum := 0
	for direction, _ := range transformations {
		newLoc, ok := applyTransform(loc, direction, grid)
		if !ok {
			continue
		}

		if grid[newLoc.Y][newLoc.X] == elevation+1 {
			sum += findTrailEndings(newLoc, elevation+1, grid, found)
		}
	}

	return sum
}

func findUniqueTrails(loc Location, elevation int, grid [][]int) int {
	if grid[loc.Y][loc.X] == 9 {
		return 1
	}

	sum := 0
	for direction, _ := range transformations {
		newLoc, ok := applyTransform(loc, direction, grid)
		if !ok {
			continue
		}

		if grid[newLoc.Y][newLoc.X] == elevation+1 {
			sum += findUniqueTrails(newLoc, elevation+1, grid)
		}
	}

	return sum
}

func Part1(grid [][]int, rows, cols int) int {
	sum := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				var visited = make(map[Location]bool)
				sum += findTrailEndings(Location{Y: row, X: col}, 0, grid, visited)
			}
		}
	}

	return sum
}

func Part2(grid [][]int, rows, cols int) int {
	sum := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				sum += findUniqueTrails(Location{Y: row, X: col}, 0, grid)
			}
		}
	}

	return sum

}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, char := range line {
			row = append(row, int(char-'0'))
		}

		grid = append(grid, row)
	}

	rows, cols := len(grid), len(grid[0])

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(grid, rows, cols)
	part2 := Part2(grid, rows, cols)

	return part1, part2
}
