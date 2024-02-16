package main

import (
	"fmt"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	// start by counting everything on the perimeter
	sum := (len(lines) * 2) + (len(lines[0]) * 2) - 4

	// you fool - you rotate the board, this won't work
	seen := map[string]bool{}

	board := lines
	for rotation := 1; rotation <= 4; rotation++ {
		watchColumns := make([]int, len(lines[0]))
		for i := 0; i < len(lines[0]); i++ {
			watchColumns[i] = i
		}

		for i := 1; i < len(board); i++ {
			tmpIndiciesToWatch := []int{}

			for j := range watchColumns {
				index := fmt.Sprintf("(%d-%d)", i, j)
				if board[i][j] > board[i-1][j] && !seen[index] {
					fmt.Printf("tree (%d, %d) counted\n", i, j)
					sum++
					seen[index] = true
					tmpIndiciesToWatch = append(tmpIndiciesToWatch, j)
				}
			}

			// update the watched columns with those that are still visible
			watchColumns = tmpIndiciesToWatch
			if len(watchColumns) == 0 {
				// and if they're all too tall, rotate the board and move on
				break
			}
		}

		board = aoc.RotateLines(board)
		fmt.Printf("Rotated Board: %v\n", board)
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
		fmt.Errorf("Part 1 Error: %e\n", err)
	}

	fmt.Printf("Part 1 Sum = %d\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 2 Error: %e\n", err)
	}

	fmt.Printf("Part 2 Sum = %d\n", sum)
}
