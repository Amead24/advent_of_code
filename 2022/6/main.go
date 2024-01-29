package main

import (
	"fmt"

	"github.com/amead24/advent_of_code/aoc"
)

func findStartOfMessageMarker(str string, length int) (int, error) {
	// I remember a leetcode problem similiar and you can use a third pointer
	// to move checker up a little bit smarter - tbd
	for i := 0; i < len(str); i++ {
		isStart := true
		seen := make(map[rune]bool)
		for _, r := range str[i : i+length] {
			if _, ok := seen[r]; ok {
				isStart = false
				break
			} else {
				seen[r] = true
			}
		}

		if isStart {
			return i + length, nil
		}
	}
	return 0, fmt.Errorf("didn't find it ¯\\_(ツ)_/¯")
}

func p1ProcessLines(lines []string) (int, error) {
	return findStartOfMessageMarker(lines[0], 4)
}

func p2ProcessLines(lines []string) (int, error) {
	return findStartOfMessageMarker(lines[0], 14)
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
