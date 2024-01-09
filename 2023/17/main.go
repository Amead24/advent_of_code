package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type Graph struct {
	g [][]int
}

func (g Graph) inBounds(n Node) bool {
	if (n.v < 0) || (n.v >= len(g.g)) || (n.h < 0) || (n.h >= len(g.g[0])) {
		return false
	}

	return true
}

func (g Graph) neighors(n Node) []Node {
	var neighbors []Node

	if g.inBounds(Node{n.v - 1, n.h}) {
		neighbors = append(neighbors, Node{n.v - 1, n.h})
	}

	if g.inBounds(Node{n.v + 1, n.h}) {
		neighbors = append(neighbors, Node{n.v + 1, n.h})
	}

	if g.inBounds(Node{n.v, n.h - 1}) {
		neighbors = append(neighbors, Node{n.v, n.h - 1})
	}

	if g.inBounds(Node{n.v, n.h + 1}) {
		neighbors = append(neighbors, Node{n.v, n.h + 1})
	}

	return neighbors
}

type Node struct {
	v int
	h int
}

func (l Node) Equal(r Node) bool {
	if l.v == r.v && l.h == r.h {
		return true
	}
	return false
}

func p1ProcessLines(lines []string) (int, error) {
	// it's been awhile
	// onwards to graphs: https://www.redblobgames.com/pathfinding/a-star/introduction.html

	var graph Graph
	for _, line := range lines {
		ints := aoc.MapInt(strings.SplitAfter(line, ""))
		graph.g = append(graph.g, ints)
	}

	// fmt.Printf("Graph: %v\n", graph)

	start := Node{v: 0, h: 0}
	// fmt.Printf("Neighbors: %v\n", graph.neighors(start))

	frontier := make([]Node, 0)
	frontier = append(frontier, start)

	seen := map[Node]bool{}
	seen[start] = true

	cameFrom := map[Node]Node{}
	cameFrom[start] = Node{}

	for len(frontier) >= 1 {
		current := frontier[0]  // pop off first element
		frontier = frontier[1:] // remove first element
		for _, next := range graph.neighors(current) {
			// if _, alreadySeen := seen[next]; !alreadySeen {
			// 	frontier = append(frontier, next)
			// 	seen[next] = true
			// }
			if _, alreadySeen := cameFrom[next]; !alreadySeen {
				frontier = append(frontier, next)
				cameFrom[next] = current
			}
		}
	}

	// fmt.Printf("CameFrom: %v\n", cameFrom)

	// now do the path finding
	current := Node{12, 12} // goal node
	path := make([]Node, 0)
	for !current.Equal(start) {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, start)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	fmt.Printf("Path: %v\n", path)

	return 0, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
}

func main() {
	lines := aoc.ReadLines("")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 1 Error: %e\n", err)
	}

	fmt.Printf("Part 1 Sum = %d\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 2 Error: %e\n", err)
	}

	fmt.Printf("Part 2 Sum = %d\n", sum)
}
