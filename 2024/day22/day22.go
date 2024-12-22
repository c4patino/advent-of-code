package day22

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

type SharedStore struct {
	Store map[string]int
	Lock  sync.Mutex
}

func secretIterator(initial int, limit int) func() (int, bool) {
	current := initial
	count := 0
	return func() (int, bool) {
		if count >= limit {
			return -1, false
		}

		current = (current ^ (current * 64)) % 16777216
		current = (current ^ (current / 32)) % 16777216
		current = (current ^ (current * 2048)) % 16777216

		count += 1
		return current, true
	}
}

func Part1(secrets []int) int {
	sum := 0
	for _, secret := range secrets {
		iter := secretIterator(secret, 2000)

		last := -1
		for secret, ok := iter(); ok; secret, ok = iter() {
			last = secret
		}

		sum += last
	}

	return sum
}

func Part2(secrets []int) int {
	var wg sync.WaitGroup

	rewards := SharedStore{Store: make(map[string]int)}
	for _, secret := range secrets {
		wg.Add(1)
		go func(secret int) {
			defer wg.Done()

			iter := secretIterator(secret, 2000)
			previousPrice := secret % 10

			buyer := make(map[string]int)
			var changes []int
			for secret, ok := iter(); ok; secret, ok = iter() {
				currentPrice := secret % 10
				change := currentPrice - previousPrice
				previousPrice = currentPrice

				changes = append(changes, change)

				if len(changes) > 4 {
					changes = changes[1:]
				}

				if len(changes) == 4 {
					key := ""
					for _, change := range changes {
						key += strconv.Itoa(change) + ","
					}

					if _, ok := buyer[key]; !ok {
						buyer[key] = currentPrice
					}
				}
			}

			rewards.Lock.Lock()
			defer rewards.Lock.Unlock()
			for key, reward := range buyer {
				rewards.Store[key] += reward
			}
		}(secret)
	}

	wg.Wait()

	max := 0
	for _, reward := range rewards.Store {
		if reward > max {
			max = reward
		}
	}

	return max
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var secrets []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		secrets = append(secrets, number)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(secrets)
	part2 := Part2(secrets)

	return part1, part2
}
