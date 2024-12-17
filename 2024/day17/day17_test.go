package day17

import (
	"os"
	"testing"
)

func TestDay17(t *testing.T) {
	filename := "test.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := "5,7,3,0"
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%s\ngot:\t%s", expectedPart1, part1)
	}

	expectedPart2 := 117440
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%d\ngot:\t%d", expectedPart2, part2)
	}
}

func TestDay17Solutions(t *testing.T) {
	filename := "input.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := "2,3,6,2,1,6,1,2,1"
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%s\ngot:\t%s", expectedPart1, part1)
	}

	expectedPart2 := 90938893795561
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%d\ngot:\t%d", expectedPart2, part2)
	}
}
