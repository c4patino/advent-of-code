package day16

import (
	"bufio"
	"container/heap"
	"image"
	"math"
	"os"
)

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type State struct {
	position  image.Point
	direction Direction
}

var transformations = map[Direction]image.Point{
	NORTH: image.Point{X: 0, Y: -1}, EAST: image.Point{X: 1, Y: 0},
	SOUTH: image.Point{X: 0, Y: 1}, WEST: image.Point{X: -1, Y: 0},
}

type PriorityQueue struct {
	items  []State
	scores map[State]int
}

func (pq PriorityQueue) Len() int            { return len(pq.items) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq.scores[pq.items[i]] < pq.scores[pq.items[j]] }
func (pq PriorityQueue) Swap(i, j int)       { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }
func (pq *PriorityQueue) Push(x interface{}) { pq.items = append(pq.items, x.(State)) }
func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.items)
	item := pq.items[n-1]
	pq.items = pq.items[:n-1]
	return item
}

func (pq *PriorityQueue) Add(state State) { heap.Push(pq, state) }
func (pq *PriorityQueue) Remove() State   { return heap.Pop(pq).(State) }

func djikstras(grid [][]rune, initial, end image.Point) map[State]int {
	pq := &PriorityQueue{[]State{}, map[State]int{}}
	heap.Init(pq)

	visited := make(map[State]bool)

	initialState := State{position: initial, direction: EAST}
	pq.scores[initialState] = 0
	pq.Push(initialState)

	for pq.Len() > 0 {
		current := pq.Remove()
		position := current.position
		direction := current.direction

		visited[current] = true

		if position == end {
			return pq.scores
		}

		next := State{position.Add(transformations[direction]), direction}
		if grid[next.position.Y][next.position.X] != '#' {
			newScore := pq.scores[current] + 1
			if existingScore, exists := pq.scores[next]; !exists || newScore < existingScore {
				pq.scores[next] = newScore
				heap.Push(pq, next)
			}
		}

		right := State{position, (direction + 3) % 4}
		newScore := pq.scores[current] + 1000
		if existingScore, exists := pq.scores[right]; !exists || newScore < existingScore {
			pq.scores[right] = newScore
			heap.Push(pq, right)
		}

		left := State{position, (direction + 1) % 4}
		newScore = pq.scores[current] + 1000
		if existingScore, exists := pq.scores[left]; !exists || newScore < existingScore {
			pq.scores[left] = newScore
			heap.Push(pq, left)
		}
	}

	return pq.scores
}

func Part1(grid [][]rune, initial, end image.Point) int {
	scores := djikstras(grid, initial, end)

	globalScores = scores

	best := math.MaxInt32
	for direction := 0; direction < 4; direction++ {
		if score, exists := scores[State{end, Direction(direction)}]; exists && score < best {
			best = score
		}
	}

	return best
}

func Part2(grid [][]rune, initial, end image.Point) int {
	bests := make(map[image.Point]bool)

	stack := []State{}
	best := math.MaxInt32
	for direction := 0; direction < 4; direction++ {
		state := State{position: end, direction: Direction(direction)}
		if score, exists := globalScores[state]; exists && score < best {
			best = score
			stack = []State{state}
		} else if exists && score == best {
			stack = append(stack, state)
		}
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		bests[current.position] = true

		currentScore := globalScores[current]
		position := current.position
		direction := current.direction

		prevState := State{position: position.Sub(transformations[direction]), direction: direction}
		if prevScore, exists := globalScores[prevState]; exists && prevScore+1 == currentScore {
			stack = append(stack, prevState)
		}

		for _, d := range []Direction{(direction + 1) % 4, (direction + 3) % 4} {
			if prevScore, exists := globalScores[State{position, d}]; exists && prevScore+1000 == currentScore {
				stack = append(stack, State{position, d})
			}
		}
	}

	return len(bests)
}

var globalScores map[State]int

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := [][]rune{}
	var initial, end image.Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune{}
		for _, r := range line {
			row = append(row, r)
			switch r {
			case 'S':
				initial = image.Point{X: len(row) - 1, Y: len(grid)}
			case 'E':
				end = image.Point{X: len(row) - 1, Y: len(grid)}
			}
		}

		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(grid, initial, end)
	part2 := Part2(grid, initial, end)

	return part1, part2
}
