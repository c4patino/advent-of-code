package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Button struct {
	X, Y int
	cost int
}

type Machine struct {
	A, B Button
	X, Y int
}

func solveMachine(machine Machine) int {
	/*
		a * aX + b * bX = pX
		a * aY + b * bY = pY
	*/
	tX := machine.X
	tY := machine.Y
	aX := machine.A.X
	aY := machine.A.Y
	bX := machine.B.X
	bY := machine.B.Y

	a := float64(tX*bY-tY*bX) / float64(aX*bY-aY*bX)
	b := float64(tY*aX-tX*aY) / float64(aX*bY-aY*bX)

	if a == math.Trunc(a) && b == math.Trunc(b) {
		return int(a*3 + b)
	}

	return 0
}

func Part1(machines []Machine) int {
	totalCost := 0
	for _, machine := range machines {
		totalCost += solveMachine(machine)
	}

	return totalCost
}

func Part2(machines []Machine) int {
	for i := range machines {
		machines[i].X += 10000000000000
		machines[i].Y += 10000000000000
	}

	return Part1(machines)
}

func Run(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		lines += fmt.Sprintf("%s\n", line)
	}

	machines := []Machine{}

	regex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)
	matches := regex.FindAllStringSubmatch(lines, -1)
	for _, machine := range matches {
		fields := make([]int, 6)
		for i, match := range machine[1:] {
			conv, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}

			fields[i] = conv
		}

		buttonA := Button{X: fields[0], Y: fields[1], cost: 3}
		buttonB := Button{X: fields[2], Y: fields[3], cost: 1}

		machines = append(machines, Machine{A: buttonA, B: buttonB, X: fields[4], Y: fields[5]})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(machines)
	part2 := Part2(machines)

	return part1, part2
}
