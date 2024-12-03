package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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
	doReMatches := doRe.FindAllStringIndex(text, -1)

	// Regex to match do quoted strings
	dontRe := regexp.MustCompile("(?i)don't\\(\\)")
	dontReMatches := dontRe.FindAllStringIndex(text, -1)

	lastDoIndex := 0
	lastDontIndex := 0
	result := 0
	for _, match := range mulReMatches {
		// Find the last instance of do before the current match
		lastDo := -1
		for i := lastDoIndex; i < len(doReMatches); i++ {
			m := doReMatches[i]
			if match[0] < m[0] {
				break
			}

			lastDo = m[0]
			lastDoIndex = i
		}

		// Find the last instance of don't before the current match
		lastDont := -1
		for i := lastDontIndex; i < len(dontReMatches); i++ {
			m := dontReMatches[i]
			if match[0] < m[0] {
				break
			}

			lastDont = m[0]
			lastDontIndex = i
		}

		// If the previous don't is after the previous do, skip this match
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
