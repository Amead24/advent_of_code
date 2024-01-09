package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	var newLines [][]string
	for i := 0; i < len(lines); i++ {
		newLines = append(newLines, strings.SplitAfter(lines[i], ""))
	}

	for i := 0; i < len(newLines); i++ {
		for ii := 0; ii < len(newLines[i]); ii++ {
			for j := i; j > 0; j-- {
				if newLines[j][ii] == "O" && newLines[j-1][ii] == "." {
					newLines[j-1][ii], newLines[j][ii] = newLines[j][ii], newLines[j-1][ii]
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(newLines); i++ {
		fmt.Printf("%v\n", newLines[i])
		for ii := 0; ii < len(newLines[0]); ii++ {
			if newLines[i][ii] == "O" {
				sum += len(newLines) - i
			}
		}
	}

	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Error: %e\n", err)
	}

	fmt.Printf("Sum = %d\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Error: %e\n", err)
	}

	fmt.Printf("Sum = %d\n", sum)
}
