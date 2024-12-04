package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var transformations = map[string]func(x, y int) (int, int){
	"N":  func(x, y int) (int, int) { return x, y - 1 },
	"S":  func(x, y int) (int, int) { return x, y + 1 },
	"E":  func(x, y int) (int, int) { return x + 1, y },
	"W":  func(x, y int) (int, int) { return x - 1, y },
	"NW": func(x, y int) (int, int) { return x - 1, y - 1 },
	"NE": func(x, y int) (int, int) { return x + 1, y - 1 },
	"SW": func(x, y int) (int, int) { return x - 1, y + 1 },
	"SE": func(x, y int) (int, int) { return x + 1, y + 1 },
}

var opposites = map[string]string{
	"N":  "S",
	"S":  "N",
	"E":  "W",
	"W":  "E",
	"NW": "SE",
	"NE": "SW",
	"SW": "NE",
	"SE": "NW",
}
var perpendiculars = map[string][]string{
	"N":  {"E", "W"},
	"S":  {"E", "W"},
	"E":  {"N", "S"},
	"W":  {"N", "S"},
	"NW": {"NE", "SW"},
	"SE": {"NE", "SW"},
	"SW": {"NW", "SE"},
	"NE": {"NW", "SE"},
}

func searchForChar(text [][]string, char string, direction string, x, y int) (int, int, bool) {
	xNew, yNew := transformations[direction](x, y)
	if xNew < 0 || xNew >= len(text[0]) || yNew < 0 || yNew >= len(text) {
		return -1, -1, false
	}

	if text[yNew][xNew] != char {
		return -1, -1, false
	}

	return xNew, yNew, true
}

func Part1(text [][]string) int {
	count := 0

	word := []string{"X", "M", "A", "S"}

	for y, line := range text {
		for x, char := range line {
			xNew, yNew := 0, 0
			success := false

			if char == word[0] {
				for key, _ := range transformations {
					xNew, yNew, success = searchForChar(text, "M", key, x, y)
					if !success {
						continue
					}

					xNew, yNew, success = searchForChar(text, "A", key, xNew, yNew)
					if !success {
						continue
					}

					xNew, yNew, success = searchForChar(text, "S", key, xNew, yNew)
					if !success {
						continue
					}

					count += 1
				}
			}
		}
	}

	return count
}

func Part2(text [][]string) int {
	count := 0

	for y, line := range text {
		for x, char := range line {

			if char == "A" {
				for key, _ := range transformations {
					if key == "N" || key == "S" || key == "E" || key == "W" {
						continue
					}

					_, _, success := searchForChar(text, "M", key, x, y)
					if !success {
						continue
					}

					_, _, success = searchForChar(text, "S", opposites[key], x, y)
					if !success {
						continue
					}

					found := false
					perpendicular := perpendiculars[key]
					for _, direction := range perpendicular {
						_, _, foundM := searchForChar(text, "M", direction, x, y)
						_, _, foundS := searchForChar(text, "S", opposites[direction], x, y)

						found = found || (foundM && foundS)
					}

					if found {
						count += 1
						break
					}
				}
			}
		}
	}

	return count
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		panic("Please provide a filename")
	}

	filename := flag.Args()[0]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var text [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		text = append(text, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	answer := Part1(text)
	fmt.Println(answer)

	answer = Part2(text)
	fmt.Println(answer)
}
