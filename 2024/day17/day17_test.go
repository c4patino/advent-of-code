package day17

import (
	"os"
	"testing"
)

type TestCase struct {
	FileName string
	Part1    interface{}
	Part2    interface{}
}

func TestDay17(t *testing.T) {
	tests := []TestCase{
		{"test.txt", "5,7,3,0", 117440},
		{"input.txt", "2,3,6,2,1,6,1,2,1", 90938893795561},
	}

	for _, test := range tests {
		if _, err := os.Stat(test.FileName); os.IsNotExist(err) {
			t.Fatalf("test file does not exist: %v", test.FileName)
		}

		part1, part2 := Run(test.FileName)
		if test.Part1 != nil && part1 != test.Part1 {
			t.Fatalf("unexpected part1:\nwant:\t%v\ngot:\t%v", test.Part1, part1)
		}

		if test.Part2 != nil && part2 != test.Part2 {
			t.Fatalf("unexpected part2:\nwant:\t%v\ngot:\t%v", test.Part2, part2)
		}
	}
}
