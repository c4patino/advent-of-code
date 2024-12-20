package day20

import (
	"bufio"
	"container/heap"
	"image"
	"math"
	"os"
	"strconv"
)

type Direction int
type Cheat struct {
	start image.Point
	end   image.Point
}

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

func cheat(results map[image.Point]int, maxDistance int) map[Cheat]int {
	skips := map[Cheat]int{}

	for start, startScore := range results {
		for end, endScore := range results {
			distance := manhattanDistance(start, end)
			if distance > maxDistance {
				continue
			}

			improvement := endScore - startScore - distance
			if improvement <= 0 {
				continue
			}

			if score, exists := skips[Cheat{start, end}]; !exists || improvement < score {
				skips[Cheat{start, end}] = improvement
			}
		}
	}

	return skips
}

func Part1(grid map[image.Point]bool, initial, target image.Point, rows, cols, threshold int) int {
	results := aStar(grid, initial, target, rows, cols)

	skips := cheat(results, 2)

	count := 0
	for _, improvement := range skips {
		if improvement >= threshold {
			count += 1
		}
	}

	return count
}

func Part2(grid map[image.Point]bool, initial, target image.Point, rows, cols, threshold int) int {
	results := aStar(grid, initial, target, rows, cols)

	skips := cheat(results, 20)

	count := 0
	for _, improvement := range skips {
		if improvement >= threshold {
			count += 1
		}
	}

	return count
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := make(map[image.Point]bool)
	var start, end image.Point

	scanner := bufio.NewScanner(file)
	rows, cols := 0, 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			scanner.Scan()
			break
		}

		cols = len(scanner.Text())
		for x, char := range scanner.Text() {
			switch char {
			case 'S':
				start = image.Point{X: x, Y: rows}
			case 'E':
				end = image.Point{X: x, Y: rows}
			case '#':
				grid[image.Point{X: x, Y: rows}] = true
			}
		}

		rows += 1
	}

	threshold, _ := strconv.Atoi(scanner.Text())

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(grid, start, end, rows, cols, threshold)
	part2 := Part2(grid, start, end, rows, cols, threshold)

	return part1, part2
}
