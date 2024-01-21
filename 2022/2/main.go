package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		opponent, me := splitLine[0], splitLine[1]
		switch opponent {
		case "A":
			switch me {
			case "X":
				sum += 1 + 3
			case "Y":
				sum += 2 + 6
			case "Z":
				sum += 3 + 0
			}
		case "B":
			switch me {
			case "X":
				sum += 1 + 0
			case "Y":
				sum += 2 + 3
			case "Z":
				sum += 3 + 6
			}
		case "C":
			switch me {
			case "X":
				sum += 1 + 6
			case "Y":
				sum += 2 + 0
			case "Z":
				sum += 3 + 3
			}
		}
	}
	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	// score := map[int]string{1: "X"}
	sum := 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		opponent, me := splitLine[0], splitLine[1]
		switch opponent {
		case "A":
			switch me {
			case "X":
				sum += 3 + 0
			case "Y":
				sum += 1 + 3
			case "Z":
				sum += 2 + 6
			}
		case "B":
			switch me {
			case "X":
				sum += 1 + 0
			case "Y":
				sum += 2 + 3
			case "Z":
				sum += 3 + 6
			}
		case "C":
			switch me {
			case "X":
				sum += 2 + 0
			case "Y":
				sum += 3 + 3
			case "Z":
				sum += 1 + 6
			}
		}
	}
	return sum, nil
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
