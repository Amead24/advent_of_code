package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	// Your code here
	lines := readLines()

	var start Coordinates
	maze := make([][]string, len(lines))
	for i, line := range lines {
		maze[i] = strings.Split(line, "")
		for ii, c := range maze[i] {
			if c == "S" {
				start = Coordinates{
					i:    i,
					ii:   ii,
					pipe: "7", // puzzle "|"
				}
			}
		}
		// fmt.Printf("newLine: %v", newLine)
	}
	// fmt.Printf("Staring at: %v\n", start)

	bounds := make(map[int][]int)
	bounds[start.i] = []int{start.ii}

	maze[start.i][start.ii] = "P"

	// cant figure out a boostrapping method - doing it manually
	previous := start
	next := Coordinates{i: 0, ii: 3, pipe: "F"} // 37, 55, "|"
	for !Equal(start, next) {
		maze[next.i][next.ii] = "P"

		if seen, keyExists := bounds[next.i]; keyExists {
			seen = append(seen, next.ii)
			bounds[next.i] = seen // go returns a copy with maps?!
		} else {
			bounds[next.i] = []int{next.ii}
		}

		// fmt.Printf("%v, seen: %v\n", next, bounds)
		tmp := findNext(previous, next, maze)
		previous = next
		next = tmp
	}

	sum := 0
	for row, seen := range bounds {
		// sorting the x coords "left to right" so that you can identify gaps
		sort.Ints(seen)

		// if len(seen) % 2 != 0 - idk
		// else - use this method??
		// https://en.wikipedia.org/wiki/Point_in_polygon

		// no change for len(seen) to be one => Sum(Xodd - Xeven)
		for i := len(seen) - 1; i > 0; i-- {
			gap := seen[i] - seen[i-1] - 1
			if (gap > 1) && (i%2 == 0) {
				sum += gap
				fmt.Printf("Row[%d] Gap between: %d & %d; Sum: %d\n", row, seen[i], seen[i-1], sum)
			}
		}
		fmt.Printf("Row[%d] Seen: %d => Sum: %d\n", row, seen, sum)
	}

	// gravity?
	// for i := len(maze) - 2; i > 0; i-- {
	// 	for ii, c := range maze[i-1] {
	// 		if (c == "P" || c == "I") && maze[i][ii] != "P" {
	// 			maze[i][ii] = "I"
	// 			sum++
	// 		}
	// 	}
	// }

	for i := 0; i < len(maze); i++ {
		fmt.Printf("%v\n", maze[i])
	}
	// 1000 -- too high
	// 566 -- ??
	// 595 - right for someone else?!
	fmt.Printf("Total Sum: %d\n", sum)
}
