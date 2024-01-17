package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/amead24/advent_of_code/aoc"
)

func isValidArrangement(str string, arrangements []int) bool {
	count, arrCounter := 0, 0

	for i := 0; i < len(str); i++ {
		if str[i] == '#' {
			count++
		} else {
			if count > 0 {
				if arrCounter >= len(arrangements) || count != arrangements[arrCounter] {
					return false
				}
				arrCounter++
				count = 0
			}
		}
	}

	if count > 0 {
		if arrCounter >= len(arrangements) || count != arrangements[arrCounter] {
			return false
		}
		arrCounter++
	}

	return arrCounter == len(arrangements)
}

type Stack struct {
	items []string
}

func (s *Stack) pop() string {
	if len(s.items) == 0 {
		return "" // Handle underflow
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *Stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) empty() bool {
	return len(s.items) == 0
}

// func p1ProcessLines(lines []string) (int, error) {
// 	sum := 0

// 	for _, line := range lines {
// 		splitLine := strings.Split(line, " ")

// 		possibilities := []string{}

// 		stack := Stack{}
// 		stack.push(splitLine[0])
// 		for !stack.empty() {
// 			springs := stack.pop()
// 			if !strings.Contains(springs, "?") {
// 				possibilities = append(possibilities, springs)
// 			} else {
// 				stack.push(strings.Replace(springs, "?", ".", 1))
// 				stack.push(strings.Replace(springs, "?", "#", 1))
// 			}
// 		}

// 		fmt.Printf("len(possibilities): %v\n", len(possibilities))

// 		arrangements := aoc.MapInt(strings.Split(splitLine[1], ","))
// 		for i := range possibilities {
// 			if isValidArrangement(possibilities[i], arrangements) {
// 				sum++
// 			}
// 		}
// 	}

// 	return sum, nil
// }

func p1ProcessLines(lines []string) (int, error) {
	sum := 0
	var wg sync.WaitGroup
	results := make(chan bool)

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		arrangements := aoc.MapInt(strings.Split(splitLine[1], ","))

		stack := Stack{}
		stack.push(splitLine[0])
		for !stack.empty() {
			springs := stack.pop()
			if strings.Contains(springs, "?") {
				stack.push(strings.Replace(springs, "?", ".", 1))
				stack.push(strings.Replace(springs, "?", "#", 1))
				continue
			}

			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				if isValidArrangement(s, arrangements) {
					results <- true
				} else {
					results <- false
				}
			}(springs)
		}
	}

	// Close the results channel when all goroutines are done.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		if result {
			sum++
		}
	}

	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	// expanding the input, we'll hook back into P1
	const RepeatCount = 5
	newLines := []string{}
	for _, line := range lines {
		splitLine := strings.Split(line, " ")

		// Repeat the left and right parts of the line
		newLeft := strings.TrimRight(strings.Repeat(splitLine[0]+"?", RepeatCount), "?")
		newRight := strings.TrimRight(strings.Repeat(splitLine[1]+",", RepeatCount), ",")

		newLines = append(newLines, newLeft+" "+newRight)
	}

	fmt.Printf("%v\n", newLines)

	return p1ProcessLines(newLines)
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
