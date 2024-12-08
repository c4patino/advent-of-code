package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

type sharedResource struct {
	resource int
	mutex    sync.Mutex
}

func reorderPages(rules [][]int, update []int) []int {
	dependents := make(map[int][]int)
	dependencies := make(map[int]int)

	inUpdate := make(map[int]bool)
	for _, num := range update {
		inUpdate[num] = true
	}

	for _, rule := range rules {
		pageA, pageB := rule[0], rule[1]
		if !inUpdate[pageA] || !inUpdate[pageB] {
			continue
		}

		dependents[pageA] = append(dependents[pageA], pageB)
		dependencies[pageB]++
	}

	ordered := []int{}
	queue := []int{}

	for _, num := range update {
		if dependencies[num] == 0 {
			queue = append(queue, num)
		}
	}

	for len(queue) > 0 {
		page := queue[0]
		queue = queue[1:]
		ordered = append(ordered, page)

		for _, dependent := range dependents[page] {
			dependencies[dependent]--
			if dependencies[dependent] == 0 {
				queue = append(queue, dependent)
			}
		}
	}

	pageSeen := make(map[int]bool)
	for _, page := range ordered {
		pageSeen[page] = true
	}
	for _, page := range update {
		if !pageSeen[page] {
			ordered = append(ordered, page)
		}
	}

	return ordered
}

func determineFailure(rules [][]int, update []int) bool {
	pageIndex := make(map[int]int)
	for i, num := range update {
		pageIndex[num] = i
	}

	failed := false
	for _, rule := range rules {
		pageA, pageB := rule[0], rule[1]

		indexA, existsA := pageIndex[pageA]
		indexB, existsB := pageIndex[pageB]

		if existsA && existsB && indexA > indexB {
			failed = true
			break
		}
	}

	return failed
}

func Part1(rules, updates [][]int) int {
	var wg sync.WaitGroup
	middleSum := sharedResource{resource: 0, mutex: sync.Mutex{}}

	for _, update := range updates {
		wg.Add(1)

		go func(update []int) {
			defer wg.Done()

			failed := determineFailure(rules, update)
			if failed {
				return
			}

			middleSum.mutex.Lock()
			middleSum.resource += update[len(update)/2]
			middleSum.mutex.Unlock()
		}(update)
	}

	wg.Wait()

	return middleSum.resource
}

func Part2(rules, updates [][]int) int {
	var wg sync.WaitGroup
	middleSum := sharedResource{resource: 0, mutex: sync.Mutex{}}

	for _, update := range updates {
		wg.Add(1)

		go func(update []int) {
			defer wg.Done()

			failed := determineFailure(rules, update)
			if !failed {
				return
			}

			fixedUpdate := reorderPages(rules, update)

			middleSum.mutex.Lock()
			middleSum.resource += fixedUpdate[len(fixedUpdate)/2]
			middleSum.mutex.Unlock()
		}(update)
	}

	wg.Wait()

	return middleSum.resource
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rules [][]int
	var updates [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		nums := strings.Split(line, "|")
		if len(nums) == 2 {
			rule := make([]int, 2)
			for i, num := range nums {
				rule[i], err = strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
			}

			rules = append(rules, rule)
		} else {
			nums := strings.Split(line, ",")
			update := []int{}
			for _, num := range nums {
				parsedNum, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}

				update = append(update, parsedNum)
			}

			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(rules, updates)
	part2 := Part2(rules, updates)

	return part1, part2
}
