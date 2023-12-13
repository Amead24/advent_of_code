package aoc

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(path string) []string {
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
