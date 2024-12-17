package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Peek() T {
	return q.items[0]
}

func Part1(text string) int {
	mulRe := regexp.MustCompile(`(?i)mul\((\d+),(\d+)\)`)
	mulReMatches := mulRe.FindAllStringIndex(text, -1)

	result := 0
	for _, match := range mulReMatches {
		lineMatch := text[match[0]:match[1]]
		captures := mulRe.FindStringSubmatch(lineMatch)

		num1, err1 := strconv.Atoi(captures[1])
		num2, err2 := strconv.Atoi(captures[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers:", err1, err2)
			continue
		}

		result += num1 * num2
	}

	return result
}

func Part2(text string) int {
	mulRe := regexp.MustCompile(`(?i)mul\((\d+),(\d+)\)`)
	mulReMatches := mulRe.FindAllStringIndex(text, -1)

	doRe := regexp.MustCompile("(?i)do\\(\\)")
	doReMatches := doRe.FindAllStringIndex(text, -1)

	dontRe := regexp.MustCompile("(?i)don't\\(\\)")
	dontReMatches := dontRe.FindAllStringIndex(text, -1)

	lastDoIndex := 0
	lastDontIndex := 0
	result := 0
	for _, match := range mulReMatches {
		lastDo := -1
		for i := lastDoIndex; i < len(doReMatches); i++ {
			m := doReMatches[i]
			if match[0] < m[0] {
				break
			}
			lastDo = m[0]
			lastDoIndex = i
		}

		lastDont := -1
		for i := lastDontIndex; i < len(dontReMatches); i++ {
			m := dontReMatches[i]
			if match[0] < m[0] {
				break
			}
			lastDont = m[0]
			lastDontIndex = i
		}

		if lastDont > lastDo {
			continue
		}

		lineMatch := text[match[0]:match[1]]
		captures := mulRe.FindStringSubmatch(lineMatch)

		num1, err1 := strconv.Atoi(captures[1])
		num2, err2 := strconv.Atoi(captures[2])
		if err1 != nil || err2 != nil {
			continue
		}

		result += num1 * num2
	}

	return result
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(text)

	part2 := Part2(text)

	return part1, part2
}
