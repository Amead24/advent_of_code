package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	var newLines [][]string
	for i := 0; i < len(lines); i++ {
		newLines = append(newLines, strings.SplitAfter(lines[i], ""))
	}

	for i := 0; i < len(newLines); i++ {
		for ii := 0; ii < len(newLines[i]); ii++ {
			for j := i; j > 0; j-- {
				if newLines[j][ii] == "O" && newLines[j-1][ii] == "." {
					newLines[j-1][ii], newLines[j][ii] = newLines[j][ii], newLines[j-1][ii]
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(newLines); i++ {
		// fmt.Printf("%v\n", newLines[i])
		for ii := 0; ii < len(newLines[0]); ii++ {
			if newLines[i][ii] == "O" {
				sum += len(newLines) - i
			}
		}
	}

	return sum, nil
}

func p2ProcessLines(lines []string) (int, error) {
	var newLines [][]string
	for i := 0; i < len(lines); i++ {
		newLines = append(newLines, strings.SplitAfter(lines[i], ""))
	}

	// for part 2 - When there's a billion cycles
	// it's reasonable to think that you need to reduce the number of cycles
	// even a ms runtime over a billion runs still takes years to run
	// thus if you can find a pattern in the cycles you can cut them out

	cycles := 1_000_000_000
	states := make(map[string]int)
	for i := 0; i < cycles; i++ {
		newState := []string{}
		for i := 0; i < len(newLines); i++ {
			line := strings.Join(newLines[i], "")
			newState = append(newState, line)
		}
		state := strings.Join(newState, "")

		// fmt.Printf("state == %v\n", state)

		if pos, stateExists := states[state]; stateExists {
			// after identifying the cycle length, we're setting i equal to the value of
			// what **is left** after dividing out all the possible cycles

			// so here we identify a cycle of 84 mod (what's left) = 66
			// meaning we have 66 loops left, so we'll set (i = 1b - 66)
			i = cycles - (cycles-i)%(i-pos)
		}

		for _, rotation := range []string{"N", "W", "S", "E"} {
			// I found that this worked on the first part
			// but on the second the rocks would get "pilled" up and a first pass, depending on direction,
			// would cause some to not move, ex: "OOO..#.#" wouldn't move until the last one because
			// the condition is only checking for O next to .
			// so let's loop through the board relative to the rotation:
			switch rotation {
			case "N":
				// moving everything north - so start north and go down
				for i := 0; i < len(newLines); i++ {
					for ii := 0; ii < len(newLines[i]); ii++ {
						for j := i; j > 0; j-- {
							if newLines[j][ii] == "O" && newLines[j-1][ii] == "." {
								newLines[j-1][ii], newLines[j][ii] = newLines[j][ii], newLines[j-1][ii]
							}
						}
					}
				}
			case "W":
				// moving everything west - so start left and go right
				for i := 0; i < len(newLines); i++ {
					for ii := 0; ii < len(newLines[i]); ii++ {
						for j := ii; j > 0; j-- {
							if newLines[i][j] == "O" && newLines[i][j-1] == "." {
								newLines[i][j-1], newLines[i][j] = newLines[i][j], newLines[i][j-1]
							}
						}
					}
				}
			case "S":
				// moving everything south - so start down and go up
				// note: could also just reverse the board and do "N"
				for i := len(newLines) - 1; i >= 0; i-- {
					for ii := 0; ii < len(newLines[i]); ii++ {
						for j := i; j < len(newLines)-1; j++ {
							if newLines[j][ii] == "O" && newLines[j+1][ii] == "." {
								newLines[j+1][ii], newLines[j][ii] = newLines[j][ii], newLines[j+1][ii]
							}
						}
					}
				}
			case "E":
				// moving everything east - so start left and go right
				for i := 0; i < len(newLines); i++ {
					for ii := len(newLines[i]) - 1; ii >= 0; ii-- {
						for j := ii; j < len(newLines[i])-1; j++ {
							if newLines[i][j] == "O" && newLines[i][j+1] == "." {
								newLines[i][j+1], newLines[i][j] = newLines[i][j], newLines[i][j+1]
							}
						}
					}
				}
			}
		}

		states[state] = i
	}

	sum := 0
	for i := 0; i < len(newLines); i++ {
		// fmt.Printf("%v\n", newLines[i])
		for ii := 0; ii < len(newLines[0]); ii++ {
			if newLines[i][ii] == "O" {
				sum += len(newLines) - i
			}
		}
	}

	return sum, nil
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Error: %e\n", err)
	}

	fmt.Printf("Sum = %d\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Error: %e\n", err)
	}

	fmt.Printf("Sum = %d\n", sum)
}
