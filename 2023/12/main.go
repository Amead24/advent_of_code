package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func isValidArrangement(str string, arrangements []int) bool {
	count, arrCounter := 0, 0
	wasContinuous := false
	for i := 0; i < len(str); i++ {
		if str[i] == '#' {
			if arrCounter == len(arrangements) {
				// the assumption for no more '.'s was wrong
				fmt.Println("ass")
				return false
			}

			wasContinuous = true
			count++
			if count > arrangements[arrCounter] {
				fmt.Println("count")
				return false
			}
		} else if wasContinuous { // first break after continuous
			arrCounter++
			if arrCounter > len(arrangements) {
				fmt.Println("arr")
				// we've gone to far - abort
				// note: len() - 1, would abort early if there was only ...'s left
				return false
			}

			// otherwise reset all the counters
			count = 0
			wasContinuous = false
		}
		// fmt.Printf("idk: %d, %d, %v\n", count, arrCounter, arrangements)
	}

	return true
}

func p1ProcessLines(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		left, right := splitLine[0], splitLine[1]
		arrangements := aoc.MapInt(strings.Split(right, ","))
		for i := 0; i < len(left); i++ {
			if left[i] == '?' {
				newStr := left[:i] + "#"
				if i+1 <= len(left) {
					newStr += left[i+1:]
				}
				fmt.Printf("neStr: %v\n", newStr)
				if isValidArrangement(newStr, arrangements) {
					sum++
				}
			}
		}
	}
	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
}

func main() {
	lines := aoc.ReadLines("./2023/12/input.txt")

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
