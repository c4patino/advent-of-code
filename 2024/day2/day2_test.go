package day2

import (
	"os"
	"testing"
)

func TestDay2(t *testing.T) {
	filename := "test.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := 2
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%d\ngot:\t%d", expectedPart1, part1)
	}

	expectedPart2 := 4
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%d\ngot:\t%d", expectedPart2, part2)
	}
}

func TestDay2Solution(t *testing.T) {
	filename := "input.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := 663
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%d\ngot:\t%d", expectedPart1, part1)
	}

	expectedPart2 := 692
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%d\ngot:\t%d", expectedPart2, part2)
	}
}
