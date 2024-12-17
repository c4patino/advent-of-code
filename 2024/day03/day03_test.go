package day03

import (
	"os"
	"testing"
)

type TestCase struct {
	FileName string
	Part1    interface{}
	Part2    interface{}
}

func TestDay3(t *testing.T) {
	tests := []TestCase{
		{"test.txt", 161, 48},
		{"input.txt", 187833789, 94455185},
	}

	for _, test := range tests {
		if _, err := os.Stat(test.FileName); os.IsNotExist(err) {
			t.Fatalf("test file does not exist: %v", test.FileName)
		}

		part1, part2 := Run(test.FileName)
		if test.Part1 != nil && part1 != test.Part1 {
			t.Fatalf("%s: unexpected part1:\nwant:\t%v\ngot:\t%v", test.FileName, test.Part1, part1)
		}

		if test.Part2 != nil && part2 != test.Part2 {
			t.Fatalf("%s: unexpected part2:\nwant:\t%v\ngot:\t%v", test.FileName, test.Part2, part2)
		}
	}
}
