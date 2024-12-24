package day25

import (
	"bufio"
	"fmt"
	"os"
)

type Key []int
type Lock []int

func Part1(keys []Key, locks []Lock) int {
	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			check := true
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					check = false
					break
				}
			}

			if check {
				count += 1
			}
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

	keys := make([]Key, 0)
	locks := make([]Lock, 0)

	current := []int{-1, -1, -1, -1, -1}
	status := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if status == "" {
			switch line {
			case "#####":
				status = "key"
			case ".....":
				status = "lock"
			}
		}

		if line == "" {
			switch status {
			case "key":
				keys = append(keys, current)
			case "lock":
				locks = append(locks, current)
			}
			status = ""
			current = []int{-1, -1, -1, -1, -1}
		}

		for i, c := range line {
			if c == '#' {
				current[i] += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(keys, locks)

	return part1, nil
}
