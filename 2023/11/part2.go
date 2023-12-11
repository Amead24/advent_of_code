package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Galaxy struct {
	x int
	y int
}

func distance(g1, g2 Galaxy) int {
	dx := math.Abs(float64(g1.x - g2.x))
	dy := math.Abs(float64(g1.y - g2.y))
	// return math.Sqrt(float64(dx*dx + dy*dy))
	return int(dy + dx)
}

func expandUniverse(lines []string) []string {
	// column spacing
	for i := 0; i < len(lines[0]); i++ {
		needsColumnExpansion := true
		for ii := 0; ii < len(lines); ii++ {
			if lines[ii][i] == '#' {
				needsColumnExpansion = false
				break
			}
		}

		if needsColumnExpansion {
			for ii := 0; ii < len(lines); ii++ {
				lines[ii] = lines[ii][:i] + "@" + lines[ii][i:]
			}
			i++
		}
	}

	// row spacing
	for i := 0; i < len(lines); i++ {
		// fmt.Printf("Before:\n%v\n", lines)
		if !strings.Contains(lines[i], "#") {
			newRow := strings.Repeat("@", len(lines[i]))
			lines = append(lines, newRow)
			copy(lines[i+1:], lines[i:])
			lines[i] = newRow
			i++
		}
		// fmt.Printf("After:\n%v\n", lines)
	}
	return lines
}

func readLines(path string) []string {
	var lines []string

	file, err := os.Open(path)
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

	return lines
}

func main() {
	universe := readLines("./2023/11/input.txt")
	universe = expandUniverse(universe)

	for i := 0; i < len(universe); i++ {
		fmt.Printf("%v\n", universe[i])
	}

	var galaxies []Galaxy
	for i := 0; i < len(universe); i++ {
		if strings.Contains(universe[i], "#") {
			for ii := 0; ii < len(universe[i]); ii++ {
				if universe[i][ii] == '#' {
					galaxies = append(galaxies, Galaxy{x: i, y: ii})
				}
			}
		}
	}

	fmt.Printf("Galaxies: %v\n", galaxies)

	sum := 0
	counter := 0
	for i := 0; i < len(galaxies); i++ {
		for ii := i + 1; ii < len(galaxies); ii++ {
			dist := 0
			for x := galaxies[i].x; x < galaxies[ii].x; x++ {
				dist += 1
				if universe[x][galaxies[i].y] == '@' {
					dist += 9
				}
				fmt.Printf("Walking Down: (%d, %d) => %d\n", x, galaxies[i].y, dist)
			}

			for y := min(galaxies[i].y, galaxies[ii].y); y <= max(galaxies[i].y, galaxies[ii].y); y++ {
				dist += 1
				if universe[galaxies[ii].x][y] == '@' {
					dist += 9
				}
				fmt.Printf("Walking Left/Right: (%d, %d) => %d\n", galaxies[ii].x, y, dist)
			}

			sum += dist
			// 717879164676 -- too high
			counter++
			fmt.Printf("C[%d] - Distance: %v <=> %v = %d; Running Sum: %d\n", counter, galaxies[i], galaxies[ii], dist, sum)
		}
	}
}
