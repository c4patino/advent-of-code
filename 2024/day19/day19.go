package day19

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

type Cache struct {
	cache map[string]int
	lock  sync.Mutex
}

func (c *Cache) Get(key string) (int, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	val, ok := c.cache[key]
	return val, ok
}

func (c *Cache) Set(key string, value int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache[key] = value
}

var cache Cache

func match(towels []string, pattern string) int {
	if pattern == "" {
		return 1
	}

	if result, ok := cache.Get(pattern); ok {
		return result
	}

	total := 0
	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) {
			total += match(towels, strings.TrimPrefix(pattern, towel))
		}
	}

	cache.Set(pattern, total)

	return total
}

func Part1(towels []string, patterns []string) int {
	var wg sync.WaitGroup
	ch := make(chan int)

	for _, pattern := range patterns {
		wg.Add(1)
		go func(pattern string) {
			defer wg.Done()
			ch <- match(towels, pattern)
		}(pattern)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	count := 0
	for result := range ch {
		if result > 0 {
			count++
		}
	}

	return count
}

func Part2(patterns []string) int {
	count := 0
	for _, pattern := range patterns {
		count += cache.cache[pattern]
	}

	return count
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var towels []string
	for _, towel := range strings.Split(scanner.Text(), ", ") {
		towels = append(towels, towel)
	}

	var patterns []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		patterns = append(patterns, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	cache = Cache{cache: make(map[string]int)}

	part1 := Part1(towels, patterns)
	part2 := Part2(patterns)

	return part1, part2
}
