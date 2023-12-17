package main

import (
	"fmt"

	"github.com/amead24/advent_of_code/aoc"
)

type Coord struct {
	row  int
	col  int
	tile rune
}

func walk(previous Coord, row, col int, contraption [][]rune) int {
	if col < 0 || col >= len(contraption[0]) || row < 0 || row >= len(contraption) {
		return 0
	}

	energizedTiles := 0
	now := Coord{row, col, contraption[row][col]}
	next := now

	if previous.col > now.col { // left to right
		if now.tile == '|' {
			energizedTiles += walk(now, row-1, col, contraption)
			energizedTiles += walk(now, row+1, col, contraption)
		} else if now.tile == '/' { // bend up
			next.row--
		} else if now.tile == '\\' { // bend down
			next.row++
		} else {
			next.col++
		}
	} else if previous.col < now.col { // right to left
		if now.tile == '|' {
			queue.append(now) // split ?
		} else if now.tile == '\\' { // bend up
			next.row--
		} else if now.tile == '/' { //bend down
			next.row++
		} else {
			next.col--
		}
	} else if previous.row > now.row { // going up
		if now.tile == '-' {
			queue.append(now) // split ?
		} else if now.tile == '/' { // bend right
			next.col++
		} else if now.tile == '\\' { // bend left
			next.col--
		} else {
			next.row--
		}
	} else if previous.row < now.row { // going down
		if now.tile == '-' {
			queue.append(now) // split ?
		} else if now.tile == '/' { // bend left
			next.col--
		} else if now.tile == '\\' { // bend right
			next.col++
		} else {
			next.row++
		}
	}

	return next
}

func p1ProcessLines(lines []string) (int, error) {
	energizedTiles := 0
	previous := Coord{0, 0}
	next := Coord{0, 1}
	// queue := []Coord{next}
	// for len(queue) > 0 {
	// 	tmp := next
	// 	next = walk(previous, next)
	// 	previous = tmp
	// 	energizedTiles++
	// }
	return energizedTiles, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
}

func main() {
	lines := aoc.ReadLines("./2023/16/input.txt")

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
