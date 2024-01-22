package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	overlap := 0
	for _, line := range lines {
		reSections := regexp.MustCompile("([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)")
		sections := reSections.FindStringSubmatch(line)
		if len(sections) != 5 {
			log.Fatalf("Expected 5, got %d, line: %v\n", len(sections), line)
		}

		numbers := aoc.MapInt(sections[1:])

		if (numbers[0] <= numbers[2]) && (numbers[1] >= numbers[3]) ||
			(numbers[2] <= numbers[0]) && (numbers[3] >= numbers[1]) {
			// slog.Warn(fmt.Sprintf("Overlap: %v\n", line))
			overlap++
		}
	}
	return overlap, nil
}

func p2ProcessLines(lines []string) (int, error) {
	overlap := 0
	for _, line := range lines {
		reSections := regexp.MustCompile("([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)")
		sections := reSections.FindStringSubmatch(line)
		if len(sections) != 5 {
			log.Fatalf("Expected 5, got %d, line: %v\n", len(sections), line)
		}

		numbers := aoc.MapInt(sections[1:])
		if (numbers[1] >= numbers[2]) && (numbers[3] >= numbers[0]) {
			// slog.Warn(fmt.Sprintf("Overlap: %v\n", line))
			overlap++
		}
	}
	return overlap, nil
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
