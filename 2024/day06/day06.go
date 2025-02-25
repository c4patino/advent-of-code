package day06

import (
	"bufio"
	"os"
	"sync"
)

var transformations = map[Direction]func(loc Location) Location{
	NORTH: func(loc Location) Location { return Location{X: loc.X, Y: loc.Y - 1} },
	SOUTH: func(loc Location) Location { return Location{X: loc.X, Y: loc.Y + 1} },
	EAST:  func(loc Location) Location { return Location{X: loc.X + 1, Y: loc.Y} },
	WEST:  func(loc Location) Location { return Location{X: loc.X - 1, Y: loc.Y} },
}

func applyTransform(loc Location, direction Direction, grid [][]string) (Location, bool) {
	newLoc := transformations[direction](loc)
	if newLoc.X < 0 || newLoc.X >= len(grid[0]) || newLoc.Y < 0 || newLoc.Y >= len(grid) {
		return loc, false
	}

	return newLoc, true
}

type Direction string
type Location struct{ Y, X int }

const (
	NORTH Direction = "N"
	SOUTH Direction = "S"
	EAST  Direction = "E"
	WEST  Direction = "W"
)

func turnRight(direction Direction) Direction {
	switch direction {
	case NORTH:
		return EAST
	case EAST:
		return SOUTH
	case SOUTH:
		return WEST
	case WEST:
		return NORTH
	}
	return ""
}

func checkLoop(grid [][]string, initial, obstruction Location) bool {
	visited := make([][][]Direction, len(grid))
	for i := range visited {
		visited[i] = make([][]Direction, len(grid[0]))
	}

	direction := NORTH
	loc := initial
	for true {
		for _, dir := range visited[loc.Y][loc.X] {
			if dir == direction {
				return true
			}
		}

		visited[loc.Y][loc.X] = append(visited[loc.Y][loc.X], direction)

		newLoc, status := applyTransform(loc, direction, grid)
		if !status {
			break
		}

		if grid[newLoc.Y][newLoc.X] == "#" || newLoc == obstruction {
			newLoc = loc
			direction = turnRight(direction)
		}

		loc = newLoc
	}

	return false
}

func Part1(grid [][]string, initial Location) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	direction := NORTH
	loc := initial
	for true {
		visited[loc.Y][loc.X] = true

		newLoc, status := applyTransform(loc, direction, grid)
		if !status {
			break
		}

		if grid[newLoc.Y][newLoc.X] == "#" {
			newLoc = loc
			direction = turnRight(direction)
		}

		loc = newLoc
	}

	visitedCount := 0
	for row := range visited {
		for col := range visited[row] {
			if visited[row][col] {
				visitedCount++
			}
		}
	}

	return visitedCount
}

func Part2(grid [][]string, initial Location) int {
	var wg sync.WaitGroup
	res := make(chan bool)

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	direction := NORTH
	loc := initial
	for true {
		visited[loc.Y][loc.X] = true

		newLoc, status := applyTransform(loc, direction, grid)
		if !status {
			break
		}

		if grid[newLoc.Y][newLoc.X] == "#" {
			newLoc = loc
			direction = turnRight(direction)
		}

		loc = newLoc
	}

	for row := range grid {
		for col := range grid[row] {
			if !visited[row][col] {
				continue
			}

			wg.Add(1)

			go func(row, col int) {
				defer wg.Done()
				res <- checkLoop(grid, initial, Location{X: col, Y: row})
			}(row, col)
		}
	}

	go func() {
		defer close(res)
		wg.Wait()
	}()

	total := 0
	for result := range res {
		if result {
			total += 1
		}
	}

	return total
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := [][]string{}
	initialLocation := Location{X: -1, Y: -1}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []string{}
		for _, char := range line {
			row = append(row, string(char))

			if char == '^' {
				initialLocation.Y = len(grid)
				initialLocation.X = len(row) - 1
			}
		}

		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(grid, initialLocation)
	part2 := Part2(grid, initialLocation)

	return part1, part2
}
