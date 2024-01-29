package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	sizes := map[string]int{}
	stack := aoc.Stack[string]{}
	stack.Push("/")

	neighbors := map[string][]string{}

	for _, line := range lines {
		if line == "$ cd .." {
			stack.Pop()
		} else if strings.HasPrefix(line, "$ cd") {
			reChangeDirectory := regexp.MustCompile(`\$ cd (.*)$`)
			directory := reChangeDirectory.FindStringSubmatch(line)
			if len(directory) != 2 {
				log.Fatalf("length mismsatch, %s => %v\n", line, directory)
			}

			stack.Push(directory[1])
		} else { // the output of an ls command
			cwd := stack.Peek()
			splitLine := strings.Split(line, " ")
			if !strings.HasPrefix(line, "dir") {
				size, _ := strconv.Atoi(splitLine[0])
				sizes[cwd] += size
			} else {
				neighbors[cwd] = append(neighbors[cwd], splitLine[1])
			}
		}
	}

	fmt.Printf("Sizes: %v\n", sizes)
	fmt.Printf("Neighbors: %v\n", neighbors)
	fmt.Printf("Stack: %v\n", stack)

	seen := map[string]bool{}
	stack = aoc.Stack[string]{}
	stack.Push("/")

	sum := 0
	// for i := 0; i < 3; i++ {
	for !stack.Empty() {
		cwd := stack.Peek()
		fmt.Printf("cwd: %s and stack: %v\n", cwd, stack)
		if dirs, hasNeighbors := neighbors[cwd]; !hasNeighbors { // is an edge node
			seen[cwd] = true
			if sizes[cwd] <= 100_000 {
				sum += sizes[cwd]
			}
			stack.Pop()
		} else {
			cwdTotal := sizes[cwd] // this is the total of any non-dirs found already
			seenAllNeighbors := true

			// see if all the neighbors have been visited
			for _, dir := range dirs {
				if !seen[dir] {

					// TODO - Implement `Stack.Contains(dir)`
					// to check for the possibility of infinite loops
					seenAllNeighbors = false
					stack.Push(dir)
					break
				} else {
					cwdTotal += sizes[dir]
				}
			}

			if seenAllNeighbors {
				// then it's okay to add this to the cwd size
				seen[cwd] = true

				sizes[cwd] = cwdTotal
				if sizes[cwd] <= 100_000 {
					sum += sizes[cwd]
				}

				stack.Pop()
			}
		}
	}

	// 1340491 - to low
	// 1280098
	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
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
