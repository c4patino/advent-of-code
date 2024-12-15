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

func TestDay%02d(t *testing.T) {
	filename := "test.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("test file does not exist: %%v", filename)
	}

	part1, part2 := Run(filename)

	expectedPart1 := -1
	if part1 != expectedPart1 {
		t.Fatalf("unexpected part1:\nwant:\t%%d\ngot:\t%%d", expectedPart1, part1)
	}

	expectedPart2 := -1
	if part2 != expectedPart2 {
		t.Fatalf("unexpected part2:\nwant:\t%%d\ngot:\t%%d", expectedPart2, part2)
	}
}

// func TestDay%02dSolutions(t *testing.T) {
// 	filename := "input.txt"
// 	if _, err := os.Stat(filename); os.IsNotExist(err) {
// 		t.Fatalf("test file does not exist: %%v", filename)
// 	}
//
// 	part1, part2 := Run(filename)
//
// 	expectedPart1 := -1
// 	if part1 != expectedPart1 {
// 		t.Fatalf("unexpected part1:\nwant:\t%%d\ngot:\t%%d", expectedPart1, part1)
// 	}
//
// 	expectedPart2 := -1
// 	if part2 != expectedPart2 {
// 		t.Fatalf("unexpected part2:\nwant:\t%%d\ngot:\t%%d", expectedPart2, part2)
// 	}
// }
`

func generateTestCode(day int) error {
	testCode := fmt.Sprintf(testCodeTemplate, day, day, day)

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
