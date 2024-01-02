package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type Digger struct {
	vertical   int
	horizontal int
}

func (digger *Digger) Move(direction string, depth int) {
	switch direction {
	case "U":
		digger.vertical += depth
	case "D":
		digger.vertical -= depth
	case "L":
		digger.horizontal -= depth
	case "R":
		digger.horizontal += depth
	}
}

func p1ProcessLines(lines []string) (int, error) {
	area := 0
	steps := 1
	digger := &Digger{vertical: 0, horizontal: 0}

	for _, line := range lines {
		splitLines := strings.Split(line, " ")
		direction, depth, _ := splitLines[0], splitLines[1], splitLines[2]
		iDepth, err := strconv.Atoi(depth)
		if err != nil {
			return 0, err
		}

		previousLocation := &Digger{vertical: digger.vertical, horizontal: digger.horizontal}
		// fmt.Printf("Previous Location (%d, %d)\n", previousLocation.vertical, previousLocation.horizontal)

		digger.Move(direction, iDepth)
		// fmt.Printf("Location (%d, %d)\n", digger.vertical, digger.horizontal)

		area += (previousLocation.vertical*digger.horizontal - previousLocation.horizontal*digger.vertical) + iDepth
		// fmt.Printf("Area: %d\n", area)
		steps += 1
	}

	// This took me a while to figure out, but this was really helpful: https://advent-of-code.xavd.id/writeups/2023/day/18/
	// The key here is to realize that the "area" of the shape after shoelace isn't including the space of the perimeter.
	// return (area / 2) + (steps / 2) - 1, nil

	// No idea why this works with changes on line 48
	// https://github.com/bsadia/aoc_goLang/blob/53ed198e644324d559a366a17c712e1b8c6bb4fe/day18/main.go
	// Note: Shoelace is only for 2D space.  So we're not finding area - We're finding VOLUME
	return (area / 2) + 1, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
}

func main() {
	lines := aoc.ReadLines("input.txt")

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
