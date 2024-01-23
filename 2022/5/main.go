package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (string, error) {
	positions := make(map[int]string)
	stacks := map[string]*aoc.Stack[string]{}
	for i, line := range lines {
		if line == "" {
			for k, r := range lines[i-1] {
				if unicode.IsDigit(r) {
					positions[k] = string(r)
				}
			}
			fmt.Printf("positions: %v\n", positions)

			for j := i - 2; j >= 0; j-- {
				for k, r := range lines[j] {
					if unicode.IsUpper(r) {
						pos := positions[k]
						if stacks[pos] == nil {
							stacks[pos] = new(aoc.Stack[string])
						}
						stacks[pos].Push(string(r))
					}
				}
			}
		} else if strings.HasPrefix(line, "move") {
			reMove := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
			movements := reMove.FindStringSubmatch(line)
			if len(movements) != 4 {
				log.Fatalln("Mismatch on movements")
			}
			loopCounterStr, src, dst := movements[1], movements[2], movements[3]
			loopCounter, _ := strconv.Atoi(loopCounterStr)

			for x := 0; x < loopCounter; x++ {
				stacks[dst].Push(stacks[src].Pop())
			}
		} else {
			continue
		}
	}

	// go doesn't guarantee any order....
	var sortedKeys []string
	for key := range stacks {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	var answer []string
	for _, key := range sortedKeys {
		answer = append(answer, stacks[key].Pop())
	}

	return strings.Join(answer, ""), nil
}

func p2ProcessLines(lines []string) (string, error) {
	positions := make(map[int]string)
	stacks := map[string]*aoc.Stack[string]{}
	for i, line := range lines {
		if line == "" {
			for k, r := range lines[i-1] {
				if unicode.IsDigit(r) {
					positions[k] = string(r)
				}
			}
			fmt.Printf("positions: %v\n", positions)

			for j := i - 2; j >= 0; j-- {
				for k, r := range lines[j] {
					if unicode.IsUpper(r) {
						pos := positions[k]
						if stacks[pos] == nil {
							stacks[pos] = new(aoc.Stack[string])
						}
						stacks[pos].Push(string(r))
					}
				}
			}
		} else if strings.HasPrefix(line, "move") {
			reMove := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
			movements := reMove.FindStringSubmatch(line)
			if len(movements) != 4 {
				log.Fatalln("Mismatch on movements")
			}
			loopCounterStr, src, dst := movements[1], movements[2], movements[3]
			loopCounter, _ := strconv.Atoi(loopCounterStr)

			crateMover9001 := aoc.Stack[string]{}
			for x := 0; x < loopCounter; x++ {
				crateMover9001.Push(stacks[src].Pop())
			}

			for !crateMover9001.Empty() {
				stacks[dst].Push(crateMover9001.Pop())
			}
		} else {
			continue
		}
	}

	// go doesn't guarantee any order....
	var sortedKeys []string
	for key := range stacks {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	var answer []string
	for _, key := range sortedKeys {
		answer = append(answer, stacks[key].Pop())
	}

	return strings.Join(answer, ""), nil
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 1 Error: %e\n", err)
	}

	fmt.Printf("Part 1 Sum = %s\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 2 Error: %e\n", err)
	}

	fmt.Printf("Part 2 Sum = %s\n", sum)
}
