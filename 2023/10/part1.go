package main

import (
	"bufio"
	"fmt"
	"log"
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
					pipe: "|",
				}
			}
		}
		// fmt.Printf("newLine: %v", newLine)
	}
	fmt.Printf("Staring at: %v\n", start)

	counter := 1
	// cant figure out a boostrapping
	previous := start
	next := Coordinates{i: 37, ii: 55, pipe: "|"}
	for !Equal(start, next) {
		fmt.Printf("c: %d; (%d, %d) @ %s\n", counter, next.i, next.ii, next.pipe)
		tmp := findNext(previous, next, maze)
		previous = next
		next = tmp
		counter++
	}

	fmt.Printf("Start: %v, Next: %v, Counter-Halved: %d / 2 = %d\n", start, next, counter, counter/2)
}
