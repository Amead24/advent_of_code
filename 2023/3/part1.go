package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// used for testing
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
	// // println(schematic)
	// lines := strings.Split(schematic, "\n")

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

	heatMap := make([][]bool, len(lines))
	for i := range heatMap {
		heatMap[i] = make([]bool, len(lines[0]))
	}
	// fmt.Printf("%v\n", heatMap)

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '*', '#', '$', '+', '/', '=', '@', '&', '%', '-':
				if 0 <= (y - 1) {
					if 0 <= (x - 1) {
						heatMap[y-1][x-1] = true
						heatMap[y][x-1] = true // left
					}

					heatMap[y-1][x] = true

					if (x + 1) < len(lines) {
						heatMap[y-1][x+1] = true
						heatMap[y][x+1] = true // right
					}
				}

				heatMap[y][x] = true

				if (y + 1) < len(lines) {
					if 0 <= (x - 1) {
						heatMap[y+1][x-1] = true
					}

					heatMap[y+1][x] = true

					if (x + 1) < len(lines) {
						heatMap[y+1][x+1] = true
					}
				}
			}
		}
		// fmt.Printf("(0, %d): %s\n", y, line)
		// fmt.Printf("(0, %d): %v\n", y, heatMap[y])
	}

	// for _, row := range heatMap {
	// 	fmt.Println(row)
	// }

	sum := 0
	for y, line := range heatMap { // can run in parallel now
		fmt.Printf("%d: %v\n", y, line)
		for idx := 0; idx < len(line); idx++ {
			// fmt.Printf("Before: (%d, %d)\n", y, idx)
			if line[idx] {
				// first check if a number is there
				_, err := strconv.Atoi(string(lines[y][idx]))
				if err == nil { // skip if not
					start, stop := idx, idx
					start_breaker, stop_breaker := false, false
					// fmt.Printf("Heatmap says go: (%d, %d)\n", idx, y)
					// fmt.Printf("%s\n", lines[y])
					// fmt.Printf("%v\n", line)
					for {
						if !start_breaker {
							if 0 <= (start - 1) {
								// fmt.Printf("Testing: %s\n", string(lines[y][start-1]))
								_, err := strconv.Atoi(string(lines[y][start-1]))
								if err == nil {
									// fmt.Printf("back one: %d\n", i)
									start -= 1
								} else {
									start_breaker = true
								}
							} else {
								start_breaker = true
							}
						}

						if !stop_breaker {
							if (stop + 1) < len(line) {
								_, err := strconv.Atoi(string(lines[y][stop+1]))
								if err == nil {
									// fmt.Printf("forward one: %d\n", i)
									stop += 1
								} else {
									stop_breaker = true
								}
							} else {
								stop_breaker = true
							}
						}

						// fmt.Printf("Looper: idx=%d; start=%d; stop=%d\n", idx, start, stop)
						if start_breaker && stop_breaker {
							fmt.Printf("Breaking: start=%d; stop=%d;\n", start, stop)
							break
						} else {
							fmt.Printf("Looping: idx=%d; start=%d; stop=%d\n", idx, start, stop)
						}
					}

					num, err := strconv.Atoi(string(lines[y][start : stop+1]))
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Printf("Counting Number: %s\n", lines[y][start:stop+1])
						idx = stop
						sum += num
					}
				}

			}
			// fmt.Printf("After: (%d, %d)\n", y, idx)
		}
		// fmt.Printf("%d: %v\n", y, line)
		fmt.Printf("Final part 1: %d\n", sum)
	}
}
