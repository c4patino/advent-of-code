package day21

import (
	"bufio"
	"image"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Keypad interface {
	GetButtonPosition(button rune) image.Point
}

type NumericKeypad struct {
	ButtonPositions map[rune]image.Point
}

func (k *NumericKeypad) GetButtonPosition(button rune) image.Point {
	return k.ButtonPositions[button]
}

type DirectionalKeypad struct {
	ButtonPositions map[rune]image.Point
}

func (k *DirectionalKeypad) GetButtonPosition(button rune) image.Point {
	return k.ButtonPositions[button]
}

var numeric = NumericKeypad{
	ButtonPositions: map[rune]image.Point{
		'A': image.Point{2, 0},
		'0': image.Point{1, 0},
		'1': image.Point{0, 1},
		'2': image.Point{1, 1},
		'3': image.Point{2, 1},
		'4': image.Point{0, 2},
		'5': image.Point{1, 2},
		'6': image.Point{2, 2},
		'7': image.Point{0, 3},
		'8': image.Point{1, 3},
		'9': image.Point{2, 3},
	},
}

var directional = DirectionalKeypad{
	ButtonPositions: map[rune]image.Point{
		'A': image.Point{2, 1},
		'^': image.Point{1, 1},
		'<': image.Point{0, 0},
		'v': image.Point{1, 0},
		'>': image.Point{2, 0},
	},
}

func getDirectional(input string, start rune) string {
	curr := directional.GetButtonPosition(start)
	seq := []rune{}

	for _, char := range input {
		dest := directional.GetButtonPosition(char)
		dx, dy := dest.X-curr.X, dest.Y-curr.Y

		horiz, vert := []rune{}, []rune{}

		for i := 0; i < int(math.Abs(float64(dx))); i++ {
			if dx >= 0 {
				horiz = append(horiz, '>')
			} else {
				horiz = append(horiz, '<')
			}
		}

		for i := 0; i < int(math.Abs(float64(dy))); i++ {
			if dy >= 0 {
				vert = append(vert, '^')
			} else {
				vert = append(vert, 'v')
			}
		}

		if curr.X == 0 && dest.Y == 1 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if curr.Y == 1 && dest.X == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, 'A')
	}

	return string(seq)
}

func getNumeric(input string, start rune) string {
	curr := numeric.GetButtonPosition(start)
	seq := []rune{}

	for _, char := range input {
		dest := numeric.GetButtonPosition(char)
		dx, dy := dest.X-curr.X, dest.Y-curr.Y

		horiz, vert := []rune{}, []rune{}
		for i := 0; i < int(math.Abs(float64(dx))); i++ {
			if dx >= 0 {
				horiz = append(horiz, '>')
			} else {
				horiz = append(horiz, '<')
			}
		}
		for i := 0; i < int(math.Abs(float64(dy))); i++ {
			if dy >= 0 {
				vert = append(vert, '^')
			} else {
				vert = append(vert, 'v')
			}
		}

		if curr.Y == 0 && dest.X == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if curr.X == 0 && dest.Y == 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, 'A')
	}

	return string(seq)
}

func split(input string) []string {
	var result []string
	var current string

	for _, char := range input {
		current += string(char)
		if char == 'A' {
			result = append(result, current)
			current = ""
		}
	}

	return result
}

func count(input string, maxRobots, robot int) int {
	if val, ok := cache[input]; ok && robot <= len(val) && val[robot-1] != 0 {
		return val[robot-1]
	}

	seq := getDirectional(input, 'A')
	if robot == maxRobots {
		return len(seq)
	}

	if _, ok := cache[input]; !ok {
		cache[input] = make([]int, maxRobots)
	}

	steps := split(seq)
	total := 0
	for _, step := range steps {
		c := count(step, maxRobots, robot+1)
		total += c
	}

	cache[input][robot-1] = total
	return total
}

var cache map[string][]int

func Part1(sequences []string) int {
	cache = make(map[string][]int)
	re := regexp.MustCompile(`\d+`)

	complexity := 0
	for _, sequence := range sequences {
		numericCode := re.FindString(sequence)
		value, err := strconv.Atoi(numericCode)
		if err != nil {
			panic(err)
		}

		moves := getNumeric(sequence, 'A')
		length := count(moves, 2, 1)

		complexity += length * value
	}

	return complexity
}

func Part2(sequences []string) int {
	cache = make(map[string][]int)
	re := regexp.MustCompile(`\d+`)

	complexity := 0
	for _, sequence := range sequences {
		numericCode := re.FindString(sequence)
		value, err := strconv.Atoi(numericCode)
		if err != nil {
			panic(err)
		}

		moves := getNumeric(sequence, 'A')
		length := count(moves, 25, 1)

		complexity += length * value
	}

	return complexity
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sequences []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		sequences = append(sequences, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(sequences)
	part2 := Part2(sequences)

	return part1, part2
}
