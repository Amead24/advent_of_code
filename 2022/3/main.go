package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	sum := 0
	var matches []string
	for _, line := range lines {
		sliceLine := strings.SplitAfter(line, "")
	loop:
		for i := 0; i < len(sliceLine)/2; i++ {
			for j := len(sliceLine) / 2; j < len(sliceLine); j++ {
				if sliceLine[i] == sliceLine[j] {
					matches = append(matches, sliceLine[i])
					break loop
				}
			}
		}
	}

	for k := 0; k < len(matches); k++ {
		// A == 65, a = 97
		value := int(rune(matches[k][0]))
		if value >= 97 {
			sum += (value - 96)
		} else {
			sum += (value - 64 + 26)
		}
	}
	return sum, nil
}

func set(slice string) map[rune]bool {
	newSet := map[rune]bool{}
	for _, r := range slice {
		newSet[r] = true
	}
	return newSet
}

func intersect(s1, s2 string) string {
	intersection := map[rune]bool{}

	s2Set := set(s2)
	for k, _ := range set(s1) {
		if s2Set[k] {
			intersection[k] = true
		}
	}

	values := []rune{}
	for k, _ := range intersection {
		values = append(values, k)
	}
	return string(values)
}

func p2ProcessLines(lines []string) (int, error) {
	sum := 0
	var matches []string
	for i := 0; i < len(lines); i += 3 {
		set := intersect(intersect(lines[i+0], lines[i+1]), lines[i+2])
		matches = append(matches, set)
	}

	for k := 0; k < len(matches); k++ {
		// A == 65, a = 97
		var priority int
		value := int(rune(matches[k][0]))
		if value >= 97 {
			priority = (value - 96)
		} else {
			priority = (value - 64 + 26)
		}
		// fmt.Printf("rune: %s (%v)\n", string(matches[k]), priority)
		sum += priority
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
