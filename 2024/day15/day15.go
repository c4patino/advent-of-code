package day15

import (
	"bufio"
	"image"
	"os"
)

type Direction string
type Object string

const (
	NORTH Direction = "^"
	SOUTH Direction = "v"
	EAST  Direction = ">"
	WEST  Direction = "<"
)

const (
	Player   Object = "@"
	Box      Object = "O"
	Wall     Object = "#"
	Empty    Object = "."
	BoxLeft  Object = "["
	BoxRight Object = "]"
)

var transformations = map[Direction]image.Point{
	NORTH: image.Point{X: 0, Y: -1},
	SOUTH: image.Point{X: 0, Y: 1},
	EAST:  image.Point{X: 1, Y: 0},
	WEST:  image.Point{X: -1, Y: 0},
}

func checkCanMove(grid [][]Object, position image.Point, direction Direction) bool {
	newPosition := position.Add(transformations[direction])
	object := grid[newPosition.Y][newPosition.X]

	var leftPos, rightPos image.Point
	switch object {
	case BoxLeft:
		leftPos, rightPos = newPosition, newPosition.Add(transformations[EAST])
	case BoxRight:
		leftPos, rightPos = newPosition.Add(transformations[WEST]), newPosition
	}

	switch object {
	case Wall:
		return false
	case Box:
		return checkCanMove(grid, newPosition, direction)
	case BoxLeft, BoxRight:
		switch direction {
		case NORTH, SOUTH:
			return checkCanMove(grid, leftPos, direction) && checkCanMove(grid, rightPos, direction)
		case EAST:
			return checkCanMove(grid, rightPos, direction)
		case WEST:
			return checkCanMove(grid, newPosition, direction)
		}
	}

	return true
}

func move(grid [][]Object, position image.Point, direction Direction) image.Point {
	newPosition := position.Add(transformations[direction])
	if !checkCanMove(grid, position, direction) {
		return position
	}

	object := grid[position.Y][position.X]
	switch grid[newPosition.Y][newPosition.X] {
	case Box:
		move(grid, newPosition, direction)
	case BoxLeft:
		switch direction {
		case NORTH, SOUTH:
			move(grid, newPosition, direction)
			move(grid, newPosition.Add(transformations[EAST]), direction)
		case EAST, WEST:
			move(grid, newPosition, direction)
		}
	case BoxRight:
		switch direction {
		case NORTH, SOUTH:
			move(grid, newPosition, direction)
			move(grid, newPosition.Add(transformations[WEST]), direction)
		case EAST, WEST:
			move(grid, newPosition, direction)
		}
	}

	grid[position.Y][position.X] = Empty
	grid[newPosition.Y][newPosition.X] = object

	return newPosition
}

func takeScore(grid [][]Object) int {
	score := 0

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == Box || grid[y][x] == BoxLeft {
				score += 100*y + x
			}
		}
	}

	return score
}

func Part1(grid [][]Object, position image.Point, movements []Direction) int {
	for _, direction := range movements {
		position = move(grid, position, direction)
	}

	return takeScore(grid)
}

func Part2(grid [][]Object, position image.Point, movements []Direction) int {
	for _, direction := range movements {
		position = move(grid, position, direction)
	}

	return takeScore(grid)
}

func parse(file *os.File) ([][]Object, image.Point, []Direction) {
	var grid [][]Object
	var position image.Point

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		row := make([]Object, len(line))
		for x, char := range line {
			if Object(char) == Player {
				position = image.Point{X: x, Y: y}
			}

			row[x] = Object(char)
		}

		grid = append(grid, row)
		y++
	}

	var movements []Direction
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			direction := Direction(char)
			movements = append(movements, direction)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid, position, movements
}

func parseLarge(file *os.File) ([][]Object, image.Point, []Direction) {
	var grid [][]Object
	var position image.Point

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var row []Object
		for x, char := range line {
			switch Object(char) {
			case Box:
				row = append(append(row, BoxLeft), BoxRight)
			case Wall:
				row = append(append(row, Wall), Wall)
			case Empty:
				row = append(append(row, Empty), Empty)
			case Player:
				row = append(append(row, Player), Empty)
				position = image.Point{X: x * 2, Y: y}
			}
		}

		grid = append(grid, row)
		y++
	}

	var movements []Direction
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			direction := Direction(char)
			movements = append(movements, direction)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid, position, movements
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid, position, movements := parse(file)
	part1 := Part1(grid, position, movements)

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	grid, position, movements = parseLarge(file)
	part2 := Part2(grid, position, movements)

	return part1, part2
}
