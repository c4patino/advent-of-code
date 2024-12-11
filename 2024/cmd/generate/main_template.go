package generate

import (
	"fmt"
	"os"
	"os/exec"
)

var mainCodeTemplate = `package day%02d

import (
	"bufio"
	"fmt"
	"os"
)


func Part1() int {
	return -1
}

func Part2() int {
	return -1
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1()
	part2 := Part2()

	return part1, part2
}
`

func generateMainCode(day int) error {
	testCode := fmt.Sprintf(mainCodeTemplate, day)

	filename := fmt.Sprintf("./day%02d/day%02d.go", day, day)

	if err := os.WriteFile(filename, []byte(testCode), 0755); err != nil {
		return err
	}

	cmd := exec.Command("go", "fmt", filename)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
