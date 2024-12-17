package day14

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	p image.Point
	v image.Point
}

func deepcopyRobots(robots []Robot) []Robot {
	newRobots := make([]Robot, len(robots))
	for step, robot := range robots {
		newRobots[step] = Robot{p: robot.p, v: robot.v}
	}

	return newRobots
}

func printField(robots []Robot, xMax, yMax int) {
	counts := make(map[image.Point]int)
	for _, robot := range robots {
		counts[robot.p]++
	}

	for j := 0; j < yMax; j++ {
		for i := 0; i < xMax; i++ {
			if count, ok := counts[image.Pt(i, j)]; !ok {
				fmt.Print(" ")
			} else {
				fmt.Print(count)
			}
		}
		fmt.Println()
	}
}

func moveRobots(robots []Robot, xMax, yMax int) {
	rect := image.Rect(0, 0, xMax, yMax)

	for i := range robots {
		robot := &robots[i]

		robot.p = robot.p.Add(robot.v)
		robot.p = robot.p.Mod(rect)
	}
}

func calculateSafety(robots []Robot, xMax, yMax int) int {
	quadrants := []int{0, 0, 0, 0}

	middleX := xMax / 2
	middleY := yMax / 2

	bounds := []image.Rectangle{
		image.Rect(0, 0, middleX, middleY),
		image.Rect(middleX, 0, xMax, middleY),
		image.Rect(0, middleY, middleX, yMax),
		image.Rect(middleX, middleY, xMax, yMax),
	}

	for _, robot := range robots {
		if robot.p.X == middleX || robot.p.Y == middleY {
			continue
		}

		for i, bound := range bounds {
			if robot.p.In(bound) {
				quadrants[i]++
				break
			}
		}
	}

	safety := 1
	for _, quad := range quadrants {
		safety *= quad
	}

	return safety
}

func calculateSTD(robots []Robot) float64 {
	numRobots := len(robots)

	var sum image.Point

	for _, robot := range robots {
		sum.X += robot.p.X
		sum.Y += robot.p.Y
	}

	xMean := float64(sum.X) / float64(numRobots)
	yMean := float64(sum.Y) / float64(numRobots)

	var variance float64
	for _, robot := range robots {
		dx := float64(robot.p.X) - xMean
		dy := float64(robot.p.Y) - yMean
		variance += dx*dx + dy*dy
	}

	return math.Sqrt(variance / float64(numRobots))
}

func Part1(robots []Robot, xMax, yMax int) int {
	for step := 0; step < 100; step++ {
		moveRobots(robots, xMax, yMax)
	}

	return calculateSafety(robots, xMax, yMax)
}

func Part2(robots []Robot, xMax, yMax int) int {
	for step := 0; true; step++ {
		moveRobots(robots, xMax, yMax)

		std := calculateSTD(robots)
		if std < 30 {
			return step + 1
		}
	}

	return -1
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var robots []Robot
	regex := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := make([]int, 4)
		for i, match := range regex.FindStringSubmatch(scanner.Text())[1:] {
			fields[i], _ = strconv.Atoi(match)
		}

		robots = append(robots, Robot{p: image.Pt(fields[0], fields[1]), v: image.Pt(fields[2], fields[3])})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(deepcopyRobots(robots), 101, 103)
	part2 := Part2(deepcopyRobots(robots), 101, 103)

	return part1, part2
}
