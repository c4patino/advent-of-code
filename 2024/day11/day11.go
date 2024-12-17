package day11

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Key struct {
	times, stone int
}

var splitCache = make(map[int][]int)
var splitCacheHits, splitCacheMisses int

var cache = make(map[Key]int)
var cacheHits, cacheMisses int

func split(stone int) (int, int) {
	if val, ok := splitCache[stone]; ok {
		splitCacheHits++
		return val[0], val[1]
	}

	splitCacheMisses++

	half := (int(math.Log10(float64(stone))) + 1) / 2
	divisor := int(math.Pow10(half))

	left := stone / divisor
	right := stone % divisor

	splitCache[stone] = []int{left, right}

	return left, right
}

func cacheKey(key Key, value int) int {
	cache[key] = value
	return value
}

func update(stone int, depth int) int {
	key := Key{depth, stone}

	if count, exists := cache[key]; exists {
		cacheMisses++
		return count
	}

	cacheHits++

	if depth == 0 {
		return cacheKey(key, 1)
	}

	if stone == 0 {
		return cacheKey(key, update(1, depth-1))
	}

	places := int(math.Log10(float64(stone))) + 1
	if places%2 == 0 {
		split1, split2 := split(stone)
		result := update(split1, depth-1) + update(split2, depth-1)
		return cacheKey(key, result)
	}

	return cacheKey(key, update(stone*2024, depth-1))
}

func Part1(stones []int, steps int) int {
	result := 0
	for _, stone := range stones {
		result += update(stone, steps)
	}

	return result
}

func Part2(stones []int, steps int) int {
	return Part1(stones, steps)
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stones := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, c := range strings.Split(line, " ") {
			id, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}

			stones = append(stones, id)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(stones, 25)
	part2 := Part2(stones, 75)

	return part1, part2
}
