package day23

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Graph map[string]map[string]bool

func (g *Graph) getNodes() []string {
	nodes := []string{}
	for node := range *g {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g *Graph) getNeighbors(node string) []string {
	neighbors := []string{}
	for neighbor := range (*g)[node] {
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func (g *Graph) getIntersection(node string, b []string) []string {
	set := make(map[string]bool)
	for _, v := range g.getNeighbors(node) {
		set[v] = true
	}

	var intersect []string
	for _, v := range b {
		if set[v] {
			intersect = append(intersect, v)
		}
	}

	return intersect
}

func (g Graph) findMaxClique() []string {
	var maxClique []string
	var bronKerbosch func(r, p, x []string)

	bronKerbosch = func(r, p, x []string) {
		if len(p) == 0 && len(x) == 0 {
			if len(r) > len(maxClique) {
				maxClique = append([]string{}, r...)
			}
			return
		}

		for i := 0; i < len(p); i++ {
			node := p[i]
			newR := append(r, node)
			newP := g.getIntersection(node, p)
			newX := g.getIntersection(node, x)
			bronKerbosch(newR, newP, newX)

			p = append(p[:i], p[i+1:]...)
			x = append(x, node)
		}
	}

	bronKerbosch([]string{}, g.getNodes(), []string{})
	return maxClique
}

func Part1(graph Graph) int {
	cliques := make(map[string]bool)
	count := 0

	for node, neighbors := range graph {
		neighborList := graph.getNeighbors(node)

		for i := 0; i < len(neighbors); i++ {
			for j := i + 1; j < len(neighbors); j++ {
				if !graph[neighborList[i]][neighborList[j]] {
					continue
				}

				triangle := []string{node, neighborList[i], neighborList[j]}
				sort.Strings(triangle)

				key := strings.Join(triangle, ",")
				if cliques[key] {
					continue
				}

				cliques[key] = true
				for _, n := range triangle {
					if strings.HasPrefix(n, "t") {
						count++
						break
					}
				}
			}
		}
	}

	return count
}

func Part2(graph Graph) string {
	maxClique := graph.findMaxClique()
	sort.Strings(maxClique)

	return strings.Join(maxClique, ",")
}

func Run(filename string) (interface{}, interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	graph := make(map[string]map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		a, b := nodes[0], nodes[1]

		if _, exists := graph[a]; !exists {
			graph[a] = make(map[string]bool)
		}
		graph[a][b] = true

		if _, exists := graph[b]; !exists {
			graph[b] = make(map[string]bool)
		}
		graph[b][a] = true
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1 := Part1(graph)
	part2 := Part2(graph)

	return part1, part2
}
