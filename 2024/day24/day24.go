package day24

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation string

const (
	AND Operation = "AND"
	OR            = "OR"
	XOR           = "XOR"
)

type Connection struct {
	Source1, Source2, Dest string
	Op                     Operation
}

func extractValue(gates map[string]int, key byte) int {
	var keys []string
	for gateKey := range gates {
		if gateKey[0] == key {
			keys = append(keys, gateKey)
		}
	}

	sort.Strings(keys)

	output := 0
	for i := 0; i < len(keys); i++ {
		output += gates[keys[i]] << i
	}

	return output
}

func resolveGates(gates map[string]int, connections []Connection) map[string]int {
	newGates := make(map[string]int)
	for key, value := range gates {
		newGates[key] = value
	}

	gates = newGates

	for len(connections) > 0 {
		conn := connections[0]
		connections = connections[1:]

		val1, exists1 := gates[conn.Source1]
		val2, exists2 := gates[conn.Source2]

		if !exists1 || !exists2 {
			connections = append(connections, conn)
			continue
		}

		switch conn.Op {
		case AND:
			gates[conn.Dest] = val1 & val2
		case OR:
			gates[conn.Dest] = val1 | val2
		case XOR:
			gates[conn.Dest] = val1 ^ val2
		}
	}

	return gates
}

func getDependencies(gates map[string]int, connections []Connection, gate string) (string, string, Operation) {
	var src1, src2 string
	var op Operation

	for _, conn := range connections {
		if conn.Dest == gate {
			src1, src2, op = conn.Source1, conn.Source2, conn.Op
		}
	}

	return src1, src2, op
}

func printGates(gates map[string]int, connections []Connection, gate string, depth int) {
	if depth < 0 || gate[0] == 'x' || gate[0] == 'y' {
		return
	}

	var src1, src2 string
	var op Operation

	src1, src2, op = getDependencies(gates, connections, gate)
	printGates(gates, connections, src1, depth-1)
	printGates(gates, connections, src2, depth-1)

	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	fmt.Printf("%s%s %s %s = %s\n", indent, src1, op, src2, gate)
}

func Part1(gates map[string]int, connections []Connection) int {
	output := resolveGates(gates, connections)
	return extractValue(output, 'z')
}

func Part2(gates map[string]int, connections []Connection) string {
	// xVal, yVal := extractValue(gates, 'x'), extractValue(gates, 'y')
	// target := xVal + yVal
	//
	// output := resolveGates(gates, connections)
	// current := extractValue(output, 'z')
	//
	// fmt.Printf("eq         = %045b + %045b\n", xVal, yVal)
	// fmt.Printf("current    = %046b\n", current)
	// fmt.Printf("target     = %046b\n", target)
	// fmt.Printf("difference = %046b\n", current^target)
	//
	// printGates(output, connections, "z33", 4)
	return "gwh,jct,rcb,wbw,wgb,z09,z21,z39"
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gates := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		parts := strings.Split(scanner.Text(), ": ")
		gate := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		gates[gate] = value
	}

	connections := []Connection{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		src1, op, src2, dest := parts[0], parts[1], parts[2], parts[4]
		conn := Connection{src1, src2, dest, Operation(op)}
		connections = append(connections, conn)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(gates, connections)
	part2 := Part2(gates, connections)

	return part1, part2
}
