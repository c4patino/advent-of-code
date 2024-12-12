package day12

import (
	"bufio"
	"maps"
	"os"
	"slices"
	"sort"
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

func applyTransform(loc Location, direction Direction, grid [][]rune) (Location, bool) {
	newLoc := transformations[direction](loc)
	if newLoc.X < 0 || newLoc.X >= len(grid[0]) || newLoc.Y < 0 || newLoc.Y >= len(grid) {
		return newLoc, false
	}

	return newLoc, true
}

func countSides(region map[Direction]map[Location]bool) int {
	sides := 0
	for dir, locations := range region {
		locs := slices.Collect(maps.Keys(locations))

		dirSides := len(locs)
		switch dir {
		case NORTH, SOUTH:
			sort.Slice(locs, func(i, j int) bool {
				if locs[i].Y != locs[j].Y {
					return locs[i].Y < locs[j].Y
				}
				return locs[i].X < locs[j].X
			})

			for i := 1; i < len(locs); i++ {
				difX := locs[i].X - locs[i-1].X
				difY := locs[i].Y - locs[i-1].Y

				if difY == 0 && difX == 1 {
					dirSides -= 1
				}
			}
		case EAST, WEST:
			sort.Slice(locs, func(i, j int) bool {
				if locs[i].X != locs[j].X {
					return locs[i].X < locs[j].X
				}
				return locs[i].Y < locs[j].Y
			})

			for i := 1; i < len(locs); i++ {
				difX := locs[i].X - locs[i-1].X
				difY := locs[i].Y - locs[i-1].Y

				if difX == 0 && difY == 1 {
					dirSides -= 1
				}
			}
		}

		sides += dirSides
	}

	return sides
}

func Part1(grid [][]rune) int {
	var visited = make(map[Location]bool)

	var dfs func(loc Location, char rune) (int, int)
	dfs = func(loc Location, char rune) (int, int) {
		visited[loc] = true

		area := 1
		perimeter := 0
		for dir, _ := range transformations {
			newLoc, ok := applyTransform(loc, dir, grid)
			if !ok || grid[newLoc.Y][newLoc.X] != char {
				perimeter++
				continue
			}

			if !visited[newLoc] {
				a, p := dfs(newLoc, char)
				area += a
				perimeter += p
			}
		}

		return area, perimeter
	}

	totalPrice := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if !visited[Location{y, x}] {
				area, perimeter := dfs(Location{y, x}, grid[y][x])
				totalPrice += area * perimeter
			}
		}
	}

	return totalPrice
}

func Part2(grid [][]rune) int {
	var visited = make(map[Location]bool)

	var dfs func(region map[Direction]map[Location]bool, loc Location, char rune) int
	dfs = func(region map[Direction]map[Location]bool, loc Location, char rune) int {
		visited[loc] = true

		area := 1

		for direction, _ := range transformations {
			newLoc, ok := applyTransform(loc, direction, grid)
			if !ok || grid[newLoc.Y][newLoc.X] != char {
				region[direction][newLoc] = true
				continue
			}

			if !visited[newLoc] {
				area += dfs(region, newLoc, char)
			}
		}

		return area
	}

	totalPrice := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if !visited[Location{y, x}] {
				region := make(map[Direction]map[Location]bool)
				for dir, _ := range transformations {
					region[dir] = make(map[Location]bool)
				}

				area := dfs(region, Location{y, x}, grid[y][x])
				sides := countSides(region)

				totalPrice += area * sides
			}
		}
	}

	return totalPrice
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune{}
		for _, r := range line {
			row = append(row, r)
		}

		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(grid)
	part2 := Part2(grid)

	return part1, part2
}
