package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func mapInt(slice []string) []int {
	sliceInts := make([]int, len(slice))
	// splitLine := strings.Fields(slice[i])
	for i, x := range slice {
		num, err := strconv.Atoi(x)
		if err == nil {
			sliceInts[i] = num
		}
	}
	return sliceInts
}
func readInput() ([]int, map[string][][]int) {
	var lines []string

	file, err := os.Open("./2023/5/input.txt")
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

	var mapName string
	listOfMaps := make(map[string][][]int, 7)
	for i := 2; i < len(lines); i++ {
		if lines[i] != "" {
			if strings.Contains(lines[i], "map:") {
				mapName = lines[i][:len(lines[i])-5]
			} else {
				listOfMaps[mapName] = append(listOfMaps[mapName], mapInt(strings.Fields(lines[i])))
			}
		}
	}

	return mapInt(strings.Fields(lines[0][7:])), listOfMaps
}

func isBetween(lower, n, upper int) bool {
	// range is exclusive - right term?
	return lower <= n && n < upper
}

func convertLookup(n, src, dst int) int {
	if src >= dst { // 50 30 10 -> N + 30 => 80
		return dst + (n - src)
	} else { // 50 80 10 -> N + (80 - 30) => 80
		return n + (dst - src)
	}
}

func main() {
	// Your code here
	initSeeds, maps := readInput()
	// fmt.Printf("Seeds: %v\n", initSeeds)

	// for k, v := range maps {
	// 	fmt.Printf("%s:\n%v\n", k, v)
	// }

	// [...] guarantees you get a fixed size instead of slice
	orderOfOperations := [...]string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	var lowest int = math.MaxInt
	locations := make([]int, len(initSeeds))
	for i, seed := range initSeeds {
		// initially setting seed and then updating it over each map
		// If the mapping isn't found default back to seed
		location := seed
		// fmt.Printf("Starting Seed: %d\n", seed)
		for _, mapName := range orderOfOperations {
			// fmt.Printf("%s:\n", mapName)
			for _, lookups := range maps[mapName] {
				if isBetween(lookups[1], location, lookups[1]+lookups[2]) { // src <= n <= src + range
					// fmt.Printf("Lookup was: %v\n", lookups)
					// fmt.Printf("%d <= %d < %d =>", lookups[0], location, lookups[0]+lookups[2])
					location = convertLookup(location, lookups[1], lookups[0])
					// fmt.Printf(" %d\n", location)
					if location < lowest {
						lowest = location
					}

					break
				}
			}
		}
		locations[i] = location
		fmt.Printf("Finished seed[%d] - location this turn: %d\n", i, location)
		fmt.Printf("Finished seed[%d] - lowest so far: %d\n", i, lowest)
	}

	xLowest := math.MaxInt64
	for _, x := range locations {
		if x < xLowest {
			xLowest = x
		}
	}

	fmt.Printf("Lows: %v\n", locations)
	fmt.Printf("Lowest Location: %d\n", xLowest)
}
