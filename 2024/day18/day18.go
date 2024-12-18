package day18

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

var transformations = map[Direction]image.Point{
	NORTH: image.Point{X: 0, Y: -1}, EAST: image.Point{X: 1, Y: 0},
	SOUTH: image.Point{X: 0, Y: 1}, WEST: image.Point{X: -1, Y: 0},
}

type PriorityQueue struct {
	Items      []image.Point
	Scores     map[image.Point]int
	Heuristics map[image.Point]int
}

func (pq PriorityQueue) Len() int { return len(pq.Items) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.Heuristics[pq.Items[i]] < pq.Heuristics[pq.Items[j]]
}
func (pq PriorityQueue) Swap(i, j int)       { pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i] }
func (pq *PriorityQueue) Push(x interface{}) { pq.Items = append(pq.Items, x.(image.Point)) }
func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.Items)
	item := pq.Items[n-1]
	pq.Items = pq.Items[:n-1]
	return item
}

func (pq *PriorityQueue) Add(point image.Point) { heap.Push(pq, point) }
func (pq *PriorityQueue) Remove() image.Point   { return heap.Pop(pq).(image.Point) }

func manhattanDistance(a, b image.Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func aStar(grid map[image.Point]bool, initial, end image.Point, rows, cols int) map[image.Point]int {
	bounds := image.Rect(0, 0, cols+1, rows+1)
	pq := &PriorityQueue{[]image.Point{}, map[image.Point]int{}, map[image.Point]int{}}
	heap.Init(pq)

	visited := make(map[image.Point]bool)

	pq.Scores[initial] = 0
	pq.Heuristics[initial] = 0 + manhattanDistance(initial, end)
	pq.Push(initial)

	for pq.Len() > 0 {
		current := pq.Remove()

		visited[current] = true

		if current == end {
			return pq.Scores
		}

		for _, direction := range []Direction{NORTH, EAST, SOUTH, WEST} {
			next := current.Add(transformations[direction])
			if !next.In(bounds) || grid[next] {
				continue
			}

			newScore := pq.Scores[current] + 1
			if existingScore, exists := pq.Scores[next]; !exists || newScore < existingScore {
				pq.Scores[next] = newScore
				pq.Heuristics[next] = newScore + manhattanDistance(next, end)
				heap.Push(pq, next)
			}
		}
	}

	return pq.Scores
}

func Part1(bytes []image.Point, rows, cols, steps int) int {
	initial := image.Point{X: 0, Y: 0}
	target := image.Point{X: cols, Y: rows}

	grid := make(map[image.Point]bool)
	for _, point := range bytes[:steps] {
		grid[point] = true
	}

	results := aStar(grid, initial, target, rows, cols)

	return results[target]
}

func Part2(bytes []image.Point, rows, cols int) string {
	left, right := 0, len(bytes)-1
	for left < right {
		middle := (left + right + 1) / 2
		if Part1(bytes, rows, cols, middle) != 0 {
			left = middle
		} else {
			right = middle - 1
		}
	}

	return fmt.Sprintf("%d,%d", bytes[left].X, bytes[left].Y)
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var bytes []image.Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			break
		}

		parts := strings.Split(line, ",")

		xStr, yStr := parts[0], parts[1]
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)

		bytes = append(bytes, image.Point{X: x, Y: y})
	}

	var params []int
	line := scanner.Text()
	for _, partStr := range strings.Split(line, ",") {
		part, _ := strconv.Atoi(partStr)
		params = append(params, part)
	}

	rows, cols, steps := params[0], params[1], params[2]

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(bytes, rows, cols, steps)
	part2 := Part2(bytes, rows, cols)

	return part1, part2
}
