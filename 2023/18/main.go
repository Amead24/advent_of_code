package main

import (
	"fmt"
	"regexp"
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
	steps := 0
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

		area += (previousLocation.vertical*digger.horizontal - previousLocation.horizontal*digger.vertical) // + iDepth
		// fmt.Printf("Area: %d\n", area)

		steps += iDepth
	}

	// This is the best explanation for area vs. what we're doing here
	// https://www.reddit.com/r/adventofcode/comments/18lg2we/comment/kdyomrm/?utm_source=share&utm_medium=web2x&context=3
	// Think of what we drew as a single 2D line.  However the digger is actually creating a 1m cube
	// at each dot that we need to keep in mind, and then hollowing out all the insides.
	// Once we enter this 3D space, we need to add 0.5m to the outside of the perimeter to get the full boundry.
	// The edge case comes in the corners, given we're always doing a square/polygon - You get 4 convex corners (3/4 tile each).
	// But because we're only adding 2/4 for each line, you get 1/4 left over x 4 corners - You need to add one more.
	// return (area / 2) + 1, nil

	// Put another way: Picks relates  A, I, and B with the formula A = I + B/2 - 1
	// We know A with Shoelace and B is the number of steps/perimeter.  Solving for I gives: I = A - B/2 + 1
	i := (area / 2) - (steps / 2) + 1

	// But the question is actually asking for: I + B
	// that is the number of points inside the perimeter + the number of points on the perimeter
	return i + steps, nil
}

func p2ProcessLines(lines []string) (int, error) {
	area := 0
	// steps := 0
	digger := &Digger{vertical: 0, horizontal: 0}

	dirs := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}

	for _, line := range lines {
		reInstructions := regexp.MustCompile(`#(\w{5})(\d)`)
		things := reInstructions.FindStringSubmatch(line)
		strHex, direction := things[1], things[2]

		hex, _ := strconv.ParseInt(strHex, 16, 64)

		previousLocation := &Digger{vertical: digger.vertical, horizontal: digger.horizontal}
		// fmt.Printf("Previous Location (%d, %d)\n", previousLocation.vertical, previousLocation.horizontal)

		digger.Move(dirs[direction], int(hex))
		// fmt.Printf("Location (%d, %d)\n", digger.vertical, digger.horizontal)

		area += (previousLocation.vertical*digger.horizontal - previousLocation.horizontal*digger.vertical) + int(hex)
		// fmt.Printf("Area: %d\n", area)
		// steps += iDepth
	}

	return (area / 2) + 1, nil
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
