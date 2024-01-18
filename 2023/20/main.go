package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

type Module struct {
	currentPulse string
	state        string
	destination  string
	modType      string
}

func processModule(modules map[string]Module, module Module, lowPulse, highPulse *int) (Module, error) {
	switch module.modType {
	case "flipflop":
		if module.currentPulse == "low" {
			if module.state == "off" {
				*highPulse++
				return Module{"high", "on", modules[module.destination].destination, modules[module.destination].modType}, nil
			} else {
				*lowPulse++
				return Module{"low", "off", modules[module.destination].destination, modules[module.destination].modType}, nil
			}
		}
		// else if a high pulse ignore it

	case "conjunction":
		flip := module.currentPulse
		if flip == "low" {
			flip = "high"
			*highPulse++
		} else {
			flip = "low"
			*lowPulse++
		}

		return Module{flip, "na", modules[module.destination].destination, modules[module.destination].modType}, nil
	}

	return Module{}, fmt.Errorf("unhandled module type: %s", module.modType)
}

func p1ProcessLines(lines []string) (int, error) {
	broadcast := []string{}
	modules := map[string]Module{}

	for _, line := range lines {
		if strings.HasPrefix(line, "broadcaster") {
			broadcast = strings.Split(line, " -> ")
		} else if strings.HasPrefix(line, "%") {
			split := strings.Split(line, " -> ")
			modules[strings.TrimLeft(split[0], "%")] = Module{"low", "off", strings.TrimSpace(split[1]), "flipflop"}
		} else if strings.HasPrefix(line, "&") {
			split := strings.Split(line, " -> ")
			modules[strings.TrimLeft(split[0], "&")] = Module{"low", "off", strings.TrimSpace(split[1]), "conjunction"}
		}
	}

	lowPulse, highPulse := 0, 0

	// stack := aoc.Stack{}
	// // todo: create stack based on the broadcast list

	// for !stack.empty() {
	// 	module := stack.pop()
	// 	todo: stack.push([processModule(module.desintations[::-1]]))

	// 	todo: processModule(module)
	// }

	// for _, node := range strings.Split(broadcast[1], ",") {
	// 	// queue := []Module{}

	// 	bcModule := Module{"low", "off", strings.TrimSpace(node), "flipflop"}
	// 	lowPulse++

	// 	// split destination
	// 	stack := []Module{}
	// 	for !stack {
	// 	}

	// 	// take the first off the queue (queue[0]) and then append the result to the end (queue[len(queue) - 1])
	// 	for {

	// 		fmt.Printf("Incoming: %v -> ", module)
	// 		newModule, err := processModule(modules, bcModule, &lowPulse, &highPulse)
	// 		if err != nil {
	// 			break
	// 		}
	// 		fmt.Printf("%v\n", newModule)
	// 	}

	// 	// for len(queue) != 0 {
	// 	// 	module := queue[0]
	// 	// 	queue = queue[1:]
	// 	// 	newModule, err := processModule(modules, module, &lowPulse, &highPulse)
	// 	// 	if err == nil {
	// 	// 		queue = append(queue, newModule)
	// 	// 	}
	// 	// 	fmt.Printf("%v\n", newModule)

	// 	// 	fmt.Printf("LP: %d, HP: %d\n", lowPulse, highPulse)
	// 	// }

	// }

	return lowPulse * 1000 * highPulse * 1000, nil
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
