package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type Part struct {
	x int
	m int
	a int
	s int
}

func NewPart(str string) Part {
	var newPart Part
	re := regexp.MustCompile(`(\w)=([0-9]+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		key := match[1]
		value, err := strconv.Atoi(match[2])
		if err != nil {
			// Handle the error according to your needs
			fmt.Printf("Error converting %v to integer: %v\n", match[2], err)
			continue
		}

		switch key {
		case "x":
			newPart.x = value
		case "m":
			newPart.m = value
		case "a":
			newPart.a = value
		case "s":
			newPart.s = value
		}
	}
	return newPart
}

func filter(workflow string, part Part) string {
	conditions := strings.Split(workflow[1:len(workflow)-1], ",")
	conditionRegex := regexp.MustCompile(`([xmas])\s*([><])\s*([0-9]+)\s*:\s*([a-zA-Z]+)`)
	for _, condition := range conditions {
		if conditionRegex.MatchString(condition) {
			matches := conditionRegex.FindStringSubmatch(condition)
			// fmt.Printf("Matches: %v\n", matches)
			field, operator, valueStr, returnStr := matches[1], matches[2], matches[3], matches[4]
			value, _ := strconv.Atoi(valueStr)
			switch field {
			case "x":
				if operator == ">" {
					if part.x > value {
						return returnStr
					}
				} else {
					if part.x < value {
						return returnStr
					}
				}
			case "m":
				if operator == ">" {
					if part.m > value {
						return returnStr
					}
				} else {
					if part.m < value {
						return returnStr
					}
				}
			case "a":
				if operator == ">" {
					if part.a > value {
						return returnStr
					}
				} else {
					if part.a < value {
						return returnStr
					}
				}
			case "s":
				if operator == ">" {
					if part.s > value {
						return returnStr
					}
				} else {
					if part.s < value {
						return returnStr
					}
				}
			}
		}
	}

	// else return the last key in the conditions
	// fmt.Printf("conditions: %v\n", conditions)
	return conditions[len(conditions)-1]
}

func p1ProcessLines(lines []string) (int, error) {
	var parts []Part
	workflows := make(map[string]string)

	partsRegex := regexp.MustCompile(`^\{(.*?)\}.*$`)
	workflowsRegex := regexp.MustCompile(`^([a-zA-Z]+)(\{.*?\}).*$`)

	for _, line := range lines {
		if line != "" {
			if matches := workflowsRegex.FindStringSubmatch(line); matches != nil {
				workflows[matches[1]] = matches[2]
			} else if matches := partsRegex.FindStringSubmatch(line); matches != nil {
				parts = append(parts, NewPart(matches[1]))
			}
		}
	}

	sum := 0
	for _, part := range parts {
		partAccepted := filter(workflows["in"], part)
	inner:
		for {
			// fmt.Printf("Part: %v, PA: %s\n", part, partAccepted)
			switch partAccepted {
			case "A":
				sum += part.x + part.a + part.m + part.s
				break inner
			case "R":
				break inner
			default:
				partAccepted = filter(workflows[partAccepted], part)
			}
		}
	}

	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	var parts []Part
	workflows := make(map[string]string)

	partsRegex := regexp.MustCompile(`^\{(.*?)\}.*$`)
	workflowsRegex := regexp.MustCompile(`^([a-zA-Z]+)(\{.*?\}).*$`)

	for _, line := range lines {
		if line != "" {
			if matches := workflowsRegex.FindStringSubmatch(line); matches != nil {
				workflows[matches[1]] = matches[2]
			} else if matches := partsRegex.FindStringSubmatch(line); matches != nil {
				parts = append(parts, NewPart(matches[1]))
			}
		}
	}

	sum := 0
	for _, part := range parts {
		partAccepted := filter(workflows["in"], part)
	inner:
		for {
			// fmt.Printf("Part: %v, PA: %s\n", part, partAccepted)
			switch partAccepted {
			case "A":
				sum += part.x + part.a + part.m + part.s
				break inner
			case "R":
				break inner
			default:
				partAccepted = filter(workflows[partAccepted], part)
			}
		}
	}

	return sum, nil
}

func main() {
	lines := aoc.ReadLines("./2023/19/input.txt")

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
