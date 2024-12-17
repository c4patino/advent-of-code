package day17

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Program struct {
	Registers    map[string]int
	Instructions []int
	Pointer      int
	Output       []int
}

func (p *Program) Interpret() {
	opcode := p.Instructions[p.Pointer]
	operand := p.Instructions[p.Pointer+1]

	value := float64(p.Registers["A"]) / math.Pow(2, float64(p.readComboOperand()))
	switch opcode {
	case 0: // adv
		p.Registers["A"] = int(value)
	case 1: // bxl
		p.Registers["B"] ^= operand
	case 2: // bst
		p.Registers["B"] = p.readComboOperand() % 8
	case 3: // jnz
		if p.Registers["A"] != 0 {
			p.Pointer = operand
			return
		}
	case 4: // bxc
		p.Registers["B"] ^= p.Registers["C"]
	case 5: // out
		p.Output = append(p.Output, p.readComboOperand()%8)
	case 6: // bdv
		p.Registers["B"] = int(value)
	case 7: // cdv
		p.Registers["C"] = int(value)
	}

	p.Pointer += 2
}

func (p *Program) readComboOperand() int {
	switch p.Instructions[p.Pointer+1] {
	case 0, 1, 2, 3:
		return p.Instructions[p.Pointer+1]
	case 4:
		return p.Registers["A"]
	case 5:
		return p.Registers["B"]
	case 6:
		return p.Registers["C"]
	default:
		panic("invalid operand")
	}
}

func (p *Program) Run() {
	for !p.Terminated() {
		p.Interpret()
	}
}

func (p *Program) Terminated() bool {
	return p.Pointer >= len(p.Instructions)
}

func (p *Program) ReadOutput() string {
	strNums := make([]string, len(p.Output))
	for i, num := range p.Output {
		strNums[i] = strconv.Itoa(num)
	}

	return strings.Join(strNums, ",")
}

func Part1(registers map[string]int, instructions []int) string {
	p := Program{
		Registers:    registers,
		Instructions: instructions,
		Pointer:      0,
		Output:       []int{},
	}

	p.Run()

	return p.ReadOutput()
}

func Part2(instructions []int) int {
	place := len(instructions) - 1
	a := int(math.Pow(8, float64(place)))

	var new Program
	for true {
		new = Program{
			Registers:    map[string]int{"A": a, "B": 0, "C": 0},
			Instructions: instructions,
			Pointer:      0,
			Output:       []int{},
		}

		new.Run()

		if slices.Equal(new.Output, new.Instructions) {
			break
		}

		for i := len(new.Output) - 1; i >= 0; i-- {
			if new.Output[i] != new.Instructions[i] {
				place = i
				break
			}
		}

		a += int(math.Pow(8, float64(place)))
	}

	return a
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`Register ([A-Z]): (\d+)`)

	registers := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			break
		}

		matches := re.FindStringSubmatch(line)
		value, _ := strconv.Atoi(matches[2])
		registers[matches[1]] = value
	}

	instructions := []int{}
	for _, num := range strings.Split(strings.Split(scanner.Text(), " ")[1], ",") {
		n, _ := strconv.Atoi(num)
		instructions = append(instructions, n)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(registers, instructions)
	part2 := Part2(instructions)

	return part1, part2
}
