package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func hash(str string) int {
	value := 0
	for _, r := range str {
		value += int(r)
		value *= 17
		value = value % 256
	}
	return value
}

func p1ProcessLines(lines []string) (int, error) {
	sum := 0
	for _, str := range strings.Split(lines[0], ",") {
		sum += hash(str)
	}
	return sum, nil
}

type Lense struct {
	label       string
	focalLength int
}

func extractBoxLabel(s string) Lense {
	re := regexp.MustCompile("^([^=-]*)(?:=([0-9]+))?")
	match := re.FindStringSubmatch(s)

	label := match[1]
	focalLength := 0 // Default value
	if len(match) > 2 && match[2] != "" {
		focalLength, _ = strconv.Atoi(match[2]) // Ignoring error for simplicity
	}

	return Lense{label: label, focalLength: focalLength}
}

func createHashMap(lines []string) map[int][]Lense {
	hashmap := make(map[int][]Lense)
	for _, str := range strings.Split(lines[0], ",") {
		operation := "="
		if strings.Index(str, "=") == -1 {
			operation = "-"
		}

		newLense := extractBoxLabel(str)
		boxNumber := hash(newLense.label)
		lenses, _ := hashmap[boxNumber]
		found := false
		for i := 0; i < len(lenses); i++ {
			if lenses[i].label == newLense.label {
				found = true
				if operation == "=" {
					lenses[i].focalLength = newLense.focalLength
				} else {
					copy(lenses[i:], lenses[i+1:])
					hashmap[boxNumber] = lenses[:len(lenses)-1]
				}
				break
			}
		}

		// CAUTION - the hashmap will return a copy of the list,
		// you need to modify the original not the copy
		if !found && operation == "=" {
			hashmap[boxNumber] = append(lenses, newLense)
		}
	}

	return hashmap
}

func p2ProcessLines(lines []string) (int, error) {
	sum := 0
	hashmap := createHashMap(lines)
	fmt.Printf("HM: %v\n", hashmap)
	for i, lenses := range hashmap {
		for j, lense := range lenses {
			sum += (i + 1) * (j + 1) * lense.focalLength
			// fmt.Printf("Lense: %v, Sum: %d\n", lense, sum)
		}
	}
	return sum, nil
}

func main() {
	lines := aoc.ReadLines("./2023/15/input.txt")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 1 Error: %e\n", err)
	}

	fmt.Printf("Part 1 Sum = %d\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 2 Error: %e\n", err)
	}

	// 99157 - to low
	fmt.Printf("Part 2 Sum = %d\n", sum)
}
