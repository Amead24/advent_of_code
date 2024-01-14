package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type Node struct {
	L string
	R string
}

func p1ProcessLines(lines []string) (int, error) {
	nodes := make(map[string]Node)
	for _, line := range lines[2:] {
		reNode := regexp.MustCompile(`^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$`)
		matches := reNode.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic("mismatch")
		}

		nodes[matches[1]] = Node{matches[2], matches[3]}
	}

	// fmt.Printf("Nodes: %v\n", nodes)

	counter := 0
	position := "AAA"
	steps := strings.SplitAfter(lines[0], "")
	for position != "ZZZ" {
		// fmt.Printf("POS: %v, Counter: %d\n", position, counter)
		switch steps[counter%len(steps)] {
		case "L":
			position = nodes[position].L
		case "R":
			position = nodes[position].R
		}

		counter++
	}

	return counter, nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func walk(nodes map[string]Node, steps []string, start string) int {
	counter := 0
	for !strings.HasSuffix(start, "Z") {
		switch steps[counter%len(steps)] {
		case "L":
			start = nodes[start].L
		case "R":
			start = nodes[start].R
		}
		counter++
	}
	return counter
}

func p2ProcessLines(lines []string) (int, error) {
	nodes := make(map[string]Node)
	starting_nodes := []string{}
	for _, line := range lines[2:] {
		reNode := regexp.MustCompile(`^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$`)
		matches := reNode.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic(fmt.Sprintf("mismatch on %s w/ %v\n", line, matches))
		}

		new_node := Node{matches[2], matches[3]}
		if strings.HasSuffix(matches[1], "A") {
			// fmt.Printf("match: %v\n", matches[1])
			starting_nodes = append(starting_nodes, matches[1])
		}

		nodes[matches[1]] = new_node
	}

	results := []int{}
	steps := strings.SplitAfter(lines[0], "")
	for _, start := range starting_nodes {
		results = append(results, walk(nodes, steps, start))
	}

	fmt.Printf("sn: %v\n", starting_nodes)
	fmt.Printf("results: %v\n", results)

	// sn: [BGA SLA PTA AAA XJA JNA]
	// results: [18961 12169 17263 13301 14999 16697]
	val := results[0]
	for i := 1; i < len(results); i++ {
		val = lcm(val, results[i])
	}

	return val, nil
}

func main() {
	lines := aoc.ReadLines("./input.txt")

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
