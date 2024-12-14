package day14

import (
	"os"
	"testing"
)

func TestDay14(t *testing.T) {
	filename := "test.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, _ := Run(filename)

	expectedPart1 := 21
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%d\ngot:\t%d", expectedPart1, part1)
	}
}

func TestDay14Solutions(t *testing.T) {
	filename := "input.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := 211692000
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%d\ngot:\t%d", expectedPart1, part1)
	}

	expectedPart2 := 6587
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%d\ngot:\t%d", expectedPart2, part2)
	}
}
