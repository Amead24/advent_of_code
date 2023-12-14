package main

import (
	"fmt"

	"github.com/amead24/advent_of_code/aoc"
)

type Mirror struct {
	rows []string
}

func (m Mirror) findReflection() (int, error) {
	// need to study - had to copy this, could not figure out for the life of me

	// why len(m.rows) - 1 ?!
	for i := 0; i < len(m.rows)-1; i++ {
		isMirrorIdx := true
		// ex: len() == 7, min(3, 7-3-2) => j:=0; j<=2; j++ which only is going from 3 out 2
		// this is a sort of sliding scale that shrinks with each iteration
		for j := 0; j <= min(i, len(m.rows)-i-2); j++ {
			if m.rows[i-j] != m.rows[i+j+1] {
				isMirrorIdx = false
				break
			}
		}

		if isMirrorIdx {
			return i + 1, nil
		}
	}

	return 0, fmt.Errorf("Can't find reflection point.")
}

func transpose(rows []string) []string {
	columns := make([]string, len(rows[0]))
	for i := range columns {
		for j := range rows {
			columns[i] += string(rows[j][i])
		}
	}

	return columns
}

func p1ProcessLines(lines []string) (int, error) {
	var mirrors []Mirror
	start := 0
	for i, line := range lines {
		if line == "" || (len(lines)-1) == i {
			mirrors = append(mirrors, Mirror{rows: lines[start:i]})
			start = i + 1
		}
	}

	sum := 0
	for i, mirror := range mirrors {
		pos, err := mirror.findReflection()
		if err != nil {
			pos, err = Mirror{rows: transpose(mirror.rows)}.findReflection()
			sum += pos
		} else {
			sum += pos * 100
		}
		fmt.Printf("Mirror[%d] - Sum: %d\n", i, sum)
	}

	return sum, nil
}

func main() {
	lines := aoc.ReadLines("./2023/13/input.txt")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Error: %e\n", err)
	}

	fmt.Printf("Sum = %d\n", sum)

}
