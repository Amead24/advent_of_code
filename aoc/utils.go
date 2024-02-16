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

func RotateLines(lines []string) []string {
	if len(lines) == 0 || len(lines[0]) == 0 {
		return []string{}
	}

	size := len(lines)
	newLines := make([]string, size)
	for i := 0; i < size; i++ {
		var newRow string
		for j := size - 1; j >= 0; j-- {
			newRow += string(lines[j][i])
		}
		newLines[i] = newRow
	}
	return newLines
}
