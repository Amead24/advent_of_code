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

type P2Node struct {
	V string
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

func p2ProcessLines(lines []string) (int, error) {
	nodes := make(map[string]Node)
	starting_nodes := []P2Node{}
	for _, line := range lines[2:] {
		reNode := regexp.MustCompile(`^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$`)
		matches := reNode.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic(fmt.Sprintf("mismatch on %s w/ %v\n", line, matches))
		}

		if matches[1][2] == 'A' {
			// fmt.Printf("match: %v\n", matches[1])
			starting_nodes = append(starting_nodes, P2Node{matches[1], matches[2], matches[3]})
		}
		nodes[matches[1]] = Node{matches[2], matches[3]}
	}

	counter := 0
	steps := strings.SplitAfter(lines[0], "")
outer:
	for {
		arrived := true
		for i := 0; i < len(starting_nodes); i++ {
			if starting_nodes[i].V[2] != 'Z' {
				arrived = false
			}

			// else we need to move each node forward
			switch steps[counter%len(steps)] {
			case "L":
				// fmt.Printf("Step: L")
				starting_nodes[i] = P2Node{starting_nodes[i].L, nodes[starting_nodes[i].L].L, nodes[starting_nodes[i].L].R}
			case "R":
				// fmt.Printf("Step: R")
				starting_nodes[i] = P2Node{starting_nodes[i].R, nodes[starting_nodes[i].R].L, nodes[starting_nodes[i].R].R}
			}

			// fmt.Printf("Moving SN: %v\n", starting_nodes)
		}

		if arrived {
			break outer
		}

		counter++
	}

	return counter, nil
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
