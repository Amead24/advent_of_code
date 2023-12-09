package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func diff(slice []int, start int) (bool, []int, int) {
	abort := true
	// init an array full of zeros
	newRow := make([]int, len(slice))
	var lastItem int

	// now find the difference
	for i := len(slice) - 1 - start; i > 0; i-- {
		newRow[i-1] = slice[i] - slice[i-1]
		if newRow[i-1] != 0 {
			abort = false
		}

		if i == (len(slice) - 1 - start) {
			// first iteration, pubish value
			lastItem = newRow[i-1]
		}
	}

	return abort, newRow, lastItem
}

func mapInt(slice []string) []int {
	// splitLine := strings.Fields(slice)
	sliceInts := make([]int, len(slice))
	for i, c := range slice {
		num, err := strconv.Atoi(c)
		if err == nil {
			sliceInts[i] = num
		}
	}
	return sliceInts
}

func readLines() []string {
	var lines []string

	file, err := os.Open("./2023/9/input.txt")
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
	lines := readLines()

	sum := 0
	// oasisReadings := make([][]int, len(lines))
	for i, line := range lines {
		splitLine := mapInt(strings.Fields(line))
		lastItems := []int{splitLine[len(splitLine)-1]}

		allZeros, row, lastItem := diff(splitLine, 0)
		lastItems = append(lastItems, lastItem)
		// fmt.Printf("Row: (%d, 0) => %v LastItem: %d\n", i, row, lastItem)

		counter := 1
		for !allZeros {
			allZeros, row, lastItem = diff(row, counter)
			lastItems = append(lastItems, lastItem)
			// fmt.Printf("Row: (%d, %d) => %v LastItem: %d\n", i, counter, row, lastItem)
			counter++
		}

		final := lastItems[len(lastItems)-1]
		// fmt.Printf("Items: %v\n", lastItems)
		for i := len(lastItems) - 1; i >= 0; i-- {
			final += lastItems[i]
			// fmt.Printf("final: %d\n", final)
		}
		// fmt.Printf("final: %d\n", final)

		// 2,098,530,125 -- don't need Math.Abs()
		// 2,113,680,473 -- too high
		// 688,407,466 - too low
		sum += final
		fmt.Printf("Reading[%d] - Sum: %d\n", i, sum)
	}
}
