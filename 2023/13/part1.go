package main

import (
	"fmt"

	"github.com/amead24/advent_of_code/aoc"
)

type Mirror struct {
	rows []string
}

func (m Mirror) findReflection() (int, error) {
	for i := 0; i < len(m.rows); i++ {
		if m.rows[i] == m.rows[i+1] {
			// indicies start at 0 so add one
			return i + 1, nil
		}
	}

	return 0, fmt.Errorf("Can't find a reflection point")
}

func transpose(rows []string) []string {
	columns := make([]string, len(rows))
	for i := 0; i < len(rows); i++ {
		col := ""
		for j := 0; j < len(rows[i]); j++ {
			col += string(rows[j][i])
		}
		columns = append(columns, col)
	}

	return columns
}

func processLines(lines []string) (int, error) {
	var mirrors []Mirror
	start := 0
	for i, line := range lines {
		if line == "" || (len(lines)-1) == i {
			mirrors = append(mirrors, Mirror{rows: lines[start:i]})
			start = i
		}
	}
	fmt.Printf("Mirrors: %v\n", mirrors)

	sum := 0
	for i, mirror := range mirrors {
		pos, err := mirror.findReflection()
		if err != nil {
			tMirror := Mirror{rows: transpose(mirror.rows)}
			fmt.Printf("tMirror: %v\n", tMirror)
			pos, err = Mirror{rows: transpose(mirror.rows)}.findReflection()
			sum += pos * 100
		} else {
			sum += pos
		}
		fmt.Printf("Mirror[%d] - Sum: %d\n", i, sum)
	}

	return sum, nil
}

func main() {
	lines := aoc.ReadLines("./2023/13/input.txt")

	sum, err := processLines(lines)
	if err != nil {
		fmt.Errorf("Error: %e\n", err)
	}

	fmt.Printf("Sum = %d\n", sum)
}
