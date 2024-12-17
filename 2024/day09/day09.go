package day09

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Block struct {
	id   interface{}
	size int
}

func calculateChecksum(blocks []Block) int {
	checksum := 0
	currentIndex := 0
	for _, block := range blocks {
		if block.id == nil {
			currentIndex += block.size
			continue
		}

		blockId := block.id.(int)
		for i := 0; i < block.size; i++ {
			checksum += currentIndex * blockId
			currentIndex += 1
		}
	}

	return checksum
}

func mergeBlocks(blocks []Block) []Block {
	if len(blocks) == 0 {
		return blocks
	}

	merged := blocks[:0]
	for _, block := range blocks {
		if block.id != nil {
			merged = append(merged, block)
			continue
		}

		if len(merged) > 0 && merged[len(merged)-1].id == nil {
			merged[len(merged)-1].size += block.size
		} else {
			merged = append(merged, block)
		}
	}

	return merged
}

func Part1(blocks []Block) int {
	reordered := []Block{}

	for len(blocks) > 0 {
		currentBlock := blocks[0]
		if currentBlock.id != nil {
			reordered = append(reordered, currentBlock)
			blocks = blocks[1:]
			continue
		}

		lastBlock := blocks[len(blocks)-1]
		for lastBlock.id == nil && len(blocks) > 1 {
			blocks = blocks[:len(blocks)-1]
			lastBlock = blocks[len(blocks)-1]
		}

		newBlock := Block{id: lastBlock.id, size: -1}
		newBlock.size = int(math.Min(float64(lastBlock.size), float64(currentBlock.size)))

		reordered = append(reordered, newBlock)

		lastBlock.size -= newBlock.size
		if lastBlock.size == 0 {
			blocks = blocks[:len(blocks)-1]
		} else {
			blocks[len(blocks)-1] = lastBlock
		}

		currentBlock.size -= newBlock.size
		if currentBlock.size == 0 {
			blocks = blocks[1:]
		} else {
			blocks[0] = currentBlock
		}
	}

	return calculateChecksum(reordered)
}

func Part2(blocks []Block) int {
	maxId := 0
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].id != nil {
			maxId = blocks[i].id.(int)
			break
		}
	}

	for i := maxId; i >= 0; i-- {
		blockIndex := -1
		for index, block := range blocks {
			if block.id == i {
				blockIndex = index
				break
			}
		}
		if blockIndex == -1 {
			continue
		}

		emptyIndex := -1
		for index, block := range blocks[:blockIndex] {
			if block.id == nil && block.size >= blocks[blockIndex].size {
				emptyIndex = index
				break
			}
		}
		if emptyIndex == -1 {
			continue
		}

		newBlock := Block{id: i, size: blocks[blockIndex].size}

		blocks[blockIndex].id = nil

		blocks[emptyIndex].size -= blocks[blockIndex].size
		if blocks[emptyIndex].size > 0 {
			blocks = append(blocks[:emptyIndex+1], blocks[emptyIndex+1:]...)
		} else {
			blocks = append(blocks[:emptyIndex], blocks[emptyIndex+1:]...)
			blockIndex -= 1
		}

		blocks = append(blocks[:emptyIndex], append([]Block{newBlock}, blocks[emptyIndex:]...)...)
		blocks = mergeBlocks(blocks)
	}

	return calculateChecksum(blocks)
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	blocks := []Block{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		currentIndex := 0
		for i, num := range strings.Split(line, "") {
			num, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			if num == 0 {
				continue
			}

			if i%2 == 0 {
				blocks = append(blocks, Block{size: num, id: currentIndex})
				currentIndex += 1
			} else {
				blocks = append(blocks, Block{size: num, id: nil})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1Blocks := append([]Block(nil), blocks...)
	part2Blocks := append([]Block(nil), blocks...)

	part1 := Part1(part1Blocks)
	part2 := Part2(part2Blocks)

	return part1, part2
}
