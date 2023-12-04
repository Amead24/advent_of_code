package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isGear(lines []string, y, x int) (bool, [][]int) {
	// calc up and down (y) before left and right (x)
	if lines[y][x] != '*' {
		return false, nil
	}

	gearCount := 0
	var gearIndicies [][]int

	// if 0 <= (y - 1) { // top row exists -- should be handled, but input has none
	if isDigit(lines[y-1][x-1]) { // and top left is digit
		gearIndicies = append(gearIndicies, []int{y - 1, x - 1})
		gearCount += 1

		// fmt.Printf("xx: (%d, %d)", x, y)
		if !isDigit(lines[y-1][x]) && isDigit(lines[y-1][x+1]) {
			// top left was a number, let's skip the middle number and jump ahead to the top right
			// if (x + 1) < len(lines[y]) { // not the end
			// if  {
			gearIndicies = append(gearIndicies, []int{y - 1, x + 1})
			gearCount += 1
			// }
			// }
		} // else [1,2, ??] // parser will catch it
	} else {
		// ughhh but you also have to be sure it's not the same number...
		// the top left is not a number, so let's check it's sibling
		// if it was a number then the parser later will determine start:stop
		if isDigit(lines[y-1][x]) {
			gearIndicies = append(gearIndicies, []int{y - 1, x})
			gearCount += 1
		} else {
			// if the middle was a number, we'll skip the top right
			// else we need to check if it's a standalone number
			// if (x + 1) < len(lines[y]) { // not the end
			if isDigit(lines[y-1][x+1]) {
				gearIndicies = append(gearIndicies, []int{y - 1, x + 1})
				gearCount += 1
			}
			// }
		}
	}
	// }

	// i should handled the left and right side but they don't exist
	// mid left
	if isDigit(lines[y][x-1]) {
		gearIndicies = append(gearIndicies, []int{y, x - 1})
		gearCount += 1
	}

	// mid right
	if isDigit(lines[y][x+1]) {
		gearIndicies = append(gearIndicies, []int{y, x + 1})
		gearCount += 1
	}

	// if (y + 1) < len(lines) { // bottom row exists
	// 	if 0 <= (x - 1) { // and left column exists
	if isDigit(lines[y+1][x-1]) {
		gearIndicies = append(gearIndicies, []int{y + 1, x - 1})
		gearCount += 1

		if !isDigit(lines[y+1][x]) && isDigit(lines[y+1][x+1]) {
			// the middle wasn't a number, skip ahead to bot righr
			// if (x + 1) < len(lines[y]) { // and right column exists
			// if  {
			gearIndicies = append(gearIndicies, []int{y + 1, x + 1})
			gearCount += 1
			// }
			// }
		} // else [1, 2, ??] - parser will catch it
	} else {
		// bot mid
		if isDigit(lines[y+1][x]) {
			gearIndicies = append(gearIndicies, []int{y + 1, x})
			gearCount += 1
		} else {
			// bottom left and middle wasn't a number, check if the bot right is
			// if (x + 1) < len(lines[y]) { // not the end
			if isDigit(lines[y+1][x+1]) {
				gearIndicies = append(gearIndicies, []int{y + 1, x + 1})
				gearCount += 1
			}
			// }
		}
	}
	// }

	if gearCount == 2 {
		// fmt.Printf("Found gear: (%d, %d)\n", x, y)
		return true, gearIndicies
	} else {
		return false, nil
	}
}

func calcGearRatio(lines []string, positions [][]int) int {
	// move the start:stop logic here - don't need heat map anymore
	totalGearRatio := 1
	for _, position := range positions {
		// position => [y, x]
		// so given numbers are only left to right - we only care for the x axis
		start, stop := position[1], position[1]
		start_breaker, stop_breaker := false, false
		// fmt.Printf("Starting gear at: (%d, %d)\n", start, stop)
		for {
			if !start_breaker {
				if 0 <= (start - 1) {
					if isDigit(lines[position[0]][start-1]) {
						start--
					} else {
						start_breaker = true
					}
				} else {
					start_breaker = true
				}
			}

			if !stop_breaker {
				if (stop + 1) < len(lines[position[0]]) {
					if isDigit(lines[position[0]][stop+1]) {
						stop++
					} else {
						stop_breaker = true
					}
				} else {
					stop_breaker = true
				}
			}

			// fmt.Printf("Looper: idx=%d; start=%d; stop=%d\n", idx, start, stop)
			if start_breaker && stop_breaker {
				// fmt.Printf("Breaking: start=%d; stop=%d;\n", start, stop)
				break
			} else {
				//fmt.Printf("Looping: start=%d; stop=%d\n", start, stop)
				continue
			}
		}

		fmt.Printf("Counting Number: %s\n", lines[position[0]][start:stop+1])
		num, err := strconv.Atoi(string(lines[position[0]][start : stop+1]))
		if err == nil {
			totalGearRatio *= num
		} else {
			log.Fatal(err)
		}
	}

	return totalGearRatio
}

func main() {
	// // Test Start
	// schematic := "" +
	// 	"467..114..\n" +
	// 	"...*......\n" +
	// 	"..35..633.\n" +
	// 	"......#...\n" +
	// 	"617*......\n" +
	// 	".....+.58.\n" +
	// 	"..592.....\n" +
	// 	"......755.\n" +
	// 	"...$.*....\n" +
	// 	".664.598.."
	// println(schematic)
	// lines := strings.Split(schematic, "\n")
	// // Test Stop

	// Final Start
	var lines []string
	file, err := os.Open("./2023/3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from file:", err)
	}
	// Final Stop

	// removing references to heat map for this one
	var sum int
	for y, line := range lines {
		for x := range line {
			if gear, positions := isGear(lines, y, x); gear {
				fmt.Printf("positions: %v\n", positions)
				sum += calcGearRatio(lines, positions)
			}
		}
	}
	fmt.Printf("Final part 2: %d\n", sum)
}
