package day07

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Equation struct {
	target  int
	factors []int
}

type Operation string

const (
	ADD Operation = "+"
	MUL Operation = "*"
	CAT Operation = "||"
)

var transformations = map[Operation]func(x, y int) int{
	ADD: func(x, y int) int { return x + y },
	MUL: func(x, y int) int { return x * y },
	CAT: func(x, y int) int {
		cat := strconv.Itoa(x) + strconv.Itoa(y)
		num, err := strconv.Atoi(cat)
		if err != nil {
			panic(err)
		}

		return num
	},
}

func evaluateAllPossibleExpressions(equation Equation, operations []Operation, agg chan<- int) {
	target := equation.target
	factors := equation.factors

	possible := []int{factors[0]}
	for _, factor := range factors[1:] {
		newPossible := make([]int, 0)

		for _, num := range possible {
			for _, operation := range operations {
				newPossible = append(newPossible, transformations[operation](num, factor))
			}
		}

		possible = newPossible
	}

	for _, num := range possible {
		if num == target {
			agg <- target
			break
		}
	}
}

func Part1(equations []Equation) int {
	var wg sync.WaitGroup
	res := make(chan int)

	legalOperations := []Operation{ADD, MUL}
	for _, equation := range equations {
		wg.Add(1)

		go func() {
			defer wg.Done()
			evaluateAllPossibleExpressions(equation, legalOperations, res)
		}()
	}

	go func() {
		wg.Wait()
		defer close(res)
	}()

	sum := 0
	for num := range res {
		sum += num
	}

	return sum
}

func Part2(equations []Equation) int {
	var wg sync.WaitGroup
	res := make(chan int)

	legalOperations := []Operation{ADD, MUL, CAT}
	for _, equation := range equations {
		wg.Add(1)

		go func() {
			defer wg.Done()
			evaluateAllPossibleExpressions(equation, legalOperations, res)
		}()
	}

	go func() {
		wg.Wait()
		defer close(res)
	}()

	sum := 0
	for num := range res {
		sum += num
	}

	return sum
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	equations := make([]Equation, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		factorsStr := strings.TrimSpace(parts[1])
		factors := make([]int, 0)
		for _, numStr := range strings.Split(factorsStr, " ") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			factors = append(factors, num)
		}

		equations = append(equations, Equation{target, factors})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(equations)
	part2 := Part2(equations)

	return part1, part2
}
