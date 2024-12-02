package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	leftList := make([]string, len(lines))
	rightList := make([]string, len(lines))

	for _, line := range lines {
		leftAndRight := strings.Split(line, "   ")
		leftList = append(leftList, leftAndRight[0])
		rightList = append(rightList, leftAndRight[1])
	}

	leftListInt := aoc.MapInt(leftList)

	rightListInt := aoc.MapInt(rightList)

	sort.Ints(leftListInt)
	sort.Ints(rightListInt)

	total := 0
	for i := 0; i < len(leftListInt); i++ {
		total += aoc.AbsInt(rightListInt[i] - leftListInt[i])
	}

	return total, nil
}

func p2ProcessLines(lines []string) (int, error) {
	leftCounter := make(map[string]int)
	rightCounter := make(map[string]int)

	for _, line := range lines {
		leftAndRight := strings.Split(line, "   ")

		_, ok := leftCounter[leftAndRight[0]]
		if ok {
			leftCounter[leftAndRight[0]] += 1
		} else {
			leftCounter[leftAndRight[0]] = 1
		}

		_, ok = rightCounter[leftAndRight[1]]
		if ok {
			rightCounter[leftAndRight[1]] += 1
		} else {
			rightCounter[leftAndRight[1]] = 1
		}
	}

	similarityScore := 0
	for k, v := range leftCounter {
		num, _ := strconv.Atoi(k)
		similarityScore += (num * rightCounter[k]) * v
	}
	return similarityScore, nil
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
