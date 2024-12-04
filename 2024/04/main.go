package main

import (
	"fmt"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			// if the loop cuts short, and you're condition to start checking is 'X'
			//  then you have to check both normal and backwards spellings
			if lines[i][j] == 'X' {
				// horizontal
				if j - 3 >= 0 { if lines[i][j-3:j] == "SAM" { total++ }}
				if j + 3 < len(lines[i]) { if lines[i][j:j+4] == "XMAS" { total++ }}

				// vetical
				if i - 3 >= 0 { if lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' { total++ }}
				if i + 3 < len(lines) { if lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' { total++ }}

				// diagonal - looking 'NE' and then 'NW'
				if j - 3 >= 0 {
					if i - 3 >= 0 { if lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' { total++ }}
					if i + 3 < len(lines) { if lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' { total++ }}
				}

				// diagonal - looking 'SE' and then 'SW'
				if j + 3 < len(lines[i]) {
					if i - 3 >= 0 { if lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' { total++ }}
					if i + 3 < len(lines) { if lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' { total++ }}
				}
			}
		}
	}
	return total, nil
}

func p2ProcessLines(lines []string) (int, error) {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'A' {
				// fmt.Printf("(%d, %d)\n", i, j)
				if i - 1 >= 0 && i + 1 < len(lines) && j - 1 >= 0 && j + 1 < len(lines[i]) {
					if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' && lines[i+1][j+1] == 'S' { total++ }
					if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' && lines[i+1][j+1] == 'M' { total++ }
					if lines[i-1][j-1] == 'M' && lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' && lines[i+1][j+1] == 'S' { total++ }
					if lines[i-1][j-1] == 'S' && lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M' && lines[i+1][j+1] == 'M' { total++ }
				}
			}
		}
	}
	return total, nil
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
