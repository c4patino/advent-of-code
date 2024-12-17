package generate

import (
	"fmt"
	"os"
	"os/exec"
)

const testCodeTemplate = `package day%02d

import (
	"os"
	"testing"
)

type TestCase struct {
	FileName string
	Part1    interface{}
	Part2    interface{}
}

func TestDay%02d(t *testing.T) {
	tests := []TestCase{
		{"test.txt", -1, -1},
		// {"input.txt", -1, -1},
	}

	for _, test := range tests {
		if _, err := os.Stat(test.FileName); os.IsNotExist(err) {
			t.Fatalf("test file does not exist: %%v", test.FileName)
		}

		part1, part2 := Run(test.FileName)
		if test.Part1 != nil && part1 != test.Part1 {
			t.Fatalf("%%s: unexpected part1:\nwant:\t%%v\ngot:\t%%v", test.FileName, test.Part1, part1)
		}

		if test.Part2 != nil && part2 != test.Part2 {
			t.Fatalf("%%s: unexpected part2:\nwant:\t%%v\ngot:\t%%v", test.FileName, test.Part2, part2)
		}
	}
}
`

func generateTestCode(day int) error {
	testCode := fmt.Sprintf(testCodeTemplate, day, day)

	filename := fmt.Sprintf("./day%02d/day%02d_test.go", day, day)

	if err := os.WriteFile(filename, []byte(testCode), 0755); err != nil {
		return err
	}

	cmd := exec.Command("go", "fmt", filename)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
