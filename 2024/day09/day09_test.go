package day09

import (
	"os"
	"testing"
)

func TestDay9(t *testing.T) {
	filename := "test.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := 1928
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%d\ngot:\t%d", expectedPart1, part1)
	}

	expectedPart2 := 2858
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%d\ngot:\t%d", expectedPart2, part2)
	}
}

func TestDay9Solution(t *testing.T) {
	filename := "input.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := 6435922584968
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%d\ngot:\t%d", expectedPart1, part1)
	}

	expectedPart2 := 6469636832766
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%d\ngot:\t%d", expectedPart2, part2)
	}
}
