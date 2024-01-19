package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type Module struct {
	Name         string
	Type         string // "flipflop" or "conjunction"
	State        string // "on", "off", or specific state for conjunction modules
	Destinations []string
	InputStates  map[string]string // Only for conjunction modules to track inputs
}

func processModule(module *Module, pulse string, lowPulses, highPulses *int) (string, []string) {
	switch module.Type {
	case "flipflop":
		if pulse == "low" {
			if module.State == "off" {
				*highPulses++
				module.State = "on"
				return "high", module.Destinations
			} else {
				*lowPulses++
				module.State = "off"
				return "low", module.Destinations
			}
		}
		// High pulse is ignored for flip-flop
		return "", nil

	case "conjunction":
		// Update the state for the incoming pulse
		module.InputStates[module.Name] = pulse

		// Check if all inputs are high
		allHigh := true
		for _, inputState := range module.InputStates {
			if inputState != "high" {
				allHigh = false
				break
			}
		}

		if allHigh {
			*lowPulses++
			return "low", module.Destinations
		} else {
			*highPulses++
			return "high", module.Destinations
		}
	}
	return "", nil
}

func p1ProcessLines(lines []string) (int, error) {
	// start lowPulse as 1 for the initial button click
	lowPulses, highPulses := 1, 0

	modules := make(map[string]*Module)
	for _, line := range lines {
		splitLine := strings.Split(line, " -> ")
		left, right := splitLine[0], splitLine[1]
		if strings.HasPrefix(left, "%") {
			modules[left[1:]] = &Module{left[1:], "flipflop", "off", aoc.Map(strings.Split(right, ","), strings.TrimSpace), map[string]string{}}
		} else if strings.HasPrefix(left, "%") {
			modules[left[1:]] = &Module{left[1:], "conjunction", "n/a", aoc.Map(strings.Split(right, ","), strings.TrimSpace), map[string]string{}}
		} else {
			modules["broadcaster"] = &Module{"broadcaster", "flipflop", "off", aoc.Map(strings.Split(right, ","), strings.TrimSpace), map[string]string{}}
		}
	}

	// using a stack with lists in reverse feels like it's similiar to the recursive approach?
	// that is, it should solve the first node in the destinations
	// all the way before moving on to the second destinations
	// aka depth-first traversal
	stack := aoc.Stack[string]{}
	stack.Push("broadcaster")
	pulseToSend, destinations := processModule(modules["broadcaster"], "low", &lowPulses, &highPulses)
	for !stack.Empty() {
		currentModuleName := stack.Pop()
		currentModule := modules[currentModuleName]
		fmt.Printf("%s -", currentModuleName)

		// Only push destinations onto the stack if a valid pulse is returned
		if pulseToSend != "" {
			for _, dest := range destinations {
				stack.Push(dest)
			}
		}

		pulseToSend, destinations = processModule(currentModule, pulseToSend, &lowPulses, &highPulses)
		fmt.Printf("%s-> %v\n", pulseToSend, destinations)
		fmt.Printf("Stack: %v\n", stack)
	}

	return lowPulses * 1000 * highPulses * 1000, nil
}

func p2ProcessLines(lines []string) (int, error) {
	return 0, nil
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
