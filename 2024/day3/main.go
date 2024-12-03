package main

import (
	"bufio"
	"flag"
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
	// Regex to match mul(x, y)
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
	// Regex to match mul(x, y)
	mulRe := regexp.MustCompile(`(?i)mul\((\d+),(\d+)\)`)
	mulReMatches := mulRe.FindAllStringIndex(text, -1)

	// Regex to match do quoted strings
	doRe := regexp.MustCompile("(?i)do\\(\\)")
	doQueue := Queue[int]{}
	for _, match := range doRe.FindAllStringIndex(text, -1) {
		doQueue.Enqueue(match[0])
	}

	// Regex to match do quoted strings
	dontRe := regexp.MustCompile("(?i)don't\\(\\)")
	dontQueue := Queue[int]{}
	for _, match := range dontRe.FindAllStringIndex(text, -1) {
		dontQueue.Enqueue(match[0])
	}

	result := 0
	for _, match := range mulReMatches {
		// Find the last instance of do before the current match
		for !doQueue.IsEmpty() && doQueue.Peek() < match[0] {
			_ = doQueue.Dequeue()
		}

		// Find the last instance of don't before the current match
		for !dontQueue.IsEmpty() && dontQueue.Peek() < match[0] {
			_ = dontQueue.Dequeue()
		}

		// If the previous don't is after the previous do, skip this match
		if !dontQueue.IsEmpty() && !doQueue.IsEmpty() && dontQueue.Peek() < doQueue.Peek() {
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

	scanner := bufio.NewScanner(file)
	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	answer := Part1(text)
	fmt.Println(answer)

	answer = Part2(text)
	fmt.Println(answer)
}
