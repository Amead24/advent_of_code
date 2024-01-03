package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Coordinates struct {
	pipe string
	i    int
	ii   int
}

func Equal(left, right Coordinates) bool {
	return left.i == right.i && left.ii == right.ii
}

func findNext(previous, now Coordinates, maze [][]string) Coordinates {
	var next Coordinates
	switch now.pipe {
	case "|":
		if previous.i > now.i { // flow is moving "up"
			next = Coordinates{i: now.i - 1, ii: now.ii}
		} else { // flow is moving "down"
			next = Coordinates{i: now.i + 1, ii: now.ii}
		}
	case "F":
		if previous.i > now.i { // flow is moving "up"
			next = Coordinates{i: now.i, ii: now.ii + 1}
		} else { // flow is moving "left and down" - The fact it's here means it's already moved left
			next = Coordinates{i: now.i + 1, ii: now.ii}
		}
	case "-":
		if previous.ii > now.ii { // flow is moving "right to left"
			next = Coordinates{i: now.i, ii: now.ii - 1}
		} else { // flow is moving left to right
			next = Coordinates{i: now.i, ii: now.ii + 1}
		}
	case "7":
		if previous.i > now.i { // flow is moving up
			next = Coordinates{i: now.i, ii: now.ii - 1}
		} else {
			next = Coordinates{i: now.i + 1, ii: now.ii}
		}
	case "J":
		if previous.i < now.i { // flow is moving down
			next = Coordinates{i: now.i, ii: now.ii - 1}
		} else {
			next = Coordinates{i: now.i - 1, ii: now.ii}
		}
	case "L":
		if previous.i < now.i {
			next = Coordinates{i: now.i, ii: now.ii + 1}
		} else {
			next = Coordinates{i: now.i - 1, ii: now.ii}
		}
	}
	// fmt.Printf("Maze: %v\n", maze)
	next.pipe = maze[next.i][next.ii]
	return next
}

func readLines() []string {
	var lines []string

	file, err := os.Open("./2023/10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from file:", err)
	}

	return lines
}

func main() {
	lines := readLines()

	// finding the starting position and creating the board
	var start Coordinates
	maze := make([][]string, len(lines))
	for i, line := range lines {
		maze[i] = strings.Split(line, "")
		for ii, c := range maze[i] {
			if c == "S" {
				start = Coordinates{
					i:    i,
					ii:   ii,
					pipe: "|", // puzzle "|", test "7"
				}
			}
		}
	}
	fmt.Printf("start: %v\n", start)

	area := 0  // for shoelace
	steps := 0 // for picks

	// cant figure out a boostrapping method - doing it manually
	previous := start
	next := Coordinates{i: 37, ii: 55, pipe: "|"} // 37, 55, "|" = 0, 3, "F"

	// now we walk the maze
	for !Equal(start, next) {
		fmt.Printf("Walking... %v\n", next)

		// if slices.Contains([]string{"F", "J", "L", "7"}, maze[next.i][next.ii]) {
		// 	area += next.i*previous.ii - next.ii*previous.i
		// }
		area += next.i*previous.ii - next.ii*previous.i

		// fmt.Printf("%v, seen: %v\n", next, bounds)
		tmp := findNext(previous, next, maze)

		previous = next
		next = tmp

		steps++
	}

	// one last stop to make it back to the start
	area += next.i*previous.ii - next.ii*previous.i

	// picks theorum - see notes on day 18, solved for I
	// I = A - B/2 - 1

	// First we need to find area => shoelace
	fmt.Printf("Area: %d, Steps: %d\n", area, steps)
	interiorPoints := int((math.Abs(float64(area)) / 2.0) - (float64(steps) / 2.0) + 1.0)

	// 1000 -- too high
	// 566 -- ??
	// 595 - right for someone else?!
	fmt.Printf("Total Interior: %d\n", interiorPoints)
}
