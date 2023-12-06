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

func mapInt(slice []string) []uint64 {
	sliceInts := make([]uint64, len(slice))
	// splitLine := strings.Fields(slice[i])
	for i, x := range slice {
		num, err := strconv.ParseUint(x, 10, 64)
		if err == nil {
			sliceInts[i] = num
		}
	}
	return sliceInts
}

func readInput() ([]uint64, map[string][][]uint64) {
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
	listOfMaps := make(map[string][][]uint64, 7)
	for i := 2; i < len(lines); i++ {
		if lines[i] != "" {
			if strings.Contains(lines[i], "map:") {
				mapName = lines[i][:len(lines[i])-5]
			} else {
				listOfMaps[mapName] = append(listOfMaps[mapName], mapInt(strings.Fields(lines[i])))
			}
		}
	}
	return mapInt(strings.Fields(lines[0][8:])), listOfMaps
}

func main() {
	// Your code here
	initSeeds, maps := readInput()

	// totalSeeds := 0
	// for i := 1; i < len(initSeeds); i += 2 {
	// 	totalSeeds += int(initSeeds[i])
	// }
	// fmt.Printf("Total Seeds: %d\n", totalSeeds)

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

	var lowest uint64 = math.MaxUint64
	for i := 0; i < len(initSeeds); i += 2 {
		// part 2 - Seeds come in pairs, an init and a range:
		// so we're going to iterate over all ~3,008,511,937
		for ii := uint64(initSeeds[i]); ii < uint64(initSeeds[i])+uint64(initSeeds[i+1]); ii++ {
			// initially setting seed and then updating it over each map
			// If the mapping isn't found default back to seed
			seed := ii

			for _, mapName := range orderOfOperations {
				for _, lookups := range maps[mapName] {
					src, dst, rangeLength := uint64(lookups[1]), uint64(lookups[0]), uint64(lookups[2])
					if src <= seed && seed < (src+rangeLength) {
						if src >= dst {
							// if src is greater than dst
							seed = dst + (seed - src)
						} else {
							// else dst greater than src
							seed = seed + (dst - src)
						}

						break
					}
				}
			}

			if seed < lowest {
				lowest = seed
			}
		}

		// 31,098,094 -- too high
		fmt.Printf("Finished seed[%d] - lowest so far: %d\n", i, lowest)
	}
}
