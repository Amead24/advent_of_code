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

func (d Digger) dig(lagoon [][]int, direction string, depth int) {
	for i := 0; i <= depth; i++ {
		switch direction {
		case "U":
			d.vertical++
		case "D":
			d.vertical--
		case "L":
			d.horizontal--
		case "R":
			d.horizontal++
		}
	}
	lagoon[d.vertical][d.vertical] = 1
}

func p1ProcessLines(lines []string) (int, error) {
	lagoon := make([][]int, 1000)
	for i := range lagoon {
		lagoon[i] = make([]int, 1000)
	}

	// picking randomly?
	digger := Digger{vertical: 500, horizontal: 500}
	lagoon[500][500] = 1

	for _, line := range lines {
		splitLines := strings.Split(line, " ")
		direction, depth, _ := splitLines[0], splitLines[1], splitLines[2]
		iDepth, _ := strconv.Atoi(depth)
		digger.dig(lagoon, direction, iDepth)
		fmt.Printf("Lagoon:\n%v\n", lagoon)
	}

	return 405, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
}

func main() {
	lines := aoc.ReadLines("")

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
