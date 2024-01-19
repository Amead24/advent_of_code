package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/amead24/advent_of_code/aoc"
)

func countElves(lines []string) ([]int, error) {
	var sum int
	elves := []int{}
	lines = append(lines, "") // too lazy to solve the last row w/ iteration
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			index := sort.SearchInts(elves, sum)
			elves = append(elves, 0)
			copy(elves[index+1:], elves[index:len(elves)-1])
			elves[index] = sum
			sum = 0
		} else {
			calories, err := strconv.Atoi(lines[i])
			if err != nil {
				return nil, fmt.Errorf("Error converting %s\n", lines[i])
			}

			sum += calories
		}
	}
	return elves, nil
}

func p1ProcessLines(lines []string) (int, error) {
	elves, err := countElves(lines)
	if err != nil {
		log.Panicf("error: %v\n", err)
	}

	return elves[len(elves)-1], nil
}

func p2ProcessLines(lines []string) (int, error) {
	elves, err := countElves(lines)
	if err != nil {
		log.Panicf("error: %v\n", err)
	}
	len := len(elves)
	return elves[len-3] + elves[len-2] + elves[len-1], nil
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
