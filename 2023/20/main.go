package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type FlipFlop struct {
	// if input == high => ignore
	// elif input == low => {
	// if state == off => send high && state := on
	// else => send low && state := off
	// }
	lastPulse bool
}

type Conjuction struct {
	lastPulse bool // low = false, high = true
}

func p1ProcessLines(lines []string) (int, error) {
	broadcast := []string{}
	flipFlops := [][]string{}
	conjuctions := [][]string{}

	for _, line := range lines {
		if strings.HasPrefix(line, "broadcaster") {
			broadcast = append(broadcast, line)
		} else if strings.HasPrefix(line, "%") {
			reFF := regexp.MustCompile("%([A-Za-z]{2}) -> (.*)$")
			flipFlops = append(flipFlops, reFF.FindStringSubmatch(line))
		} else if strings.HasPrefix(line, "&") {
			reCJ := regexp.MustCompile("&([A-Za-z]{2}) -> (.*)$")
			conjuctions = append(conjuctions, reCJ.FindStringSubmatch(line))
		}
	}

	fmt.Printf("bc: %v\nff: %v\ncj: %v\n", broadcast, flipFlops, conjuctions)

	return 0, nil
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
