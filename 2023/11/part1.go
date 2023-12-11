package main

import (
	"bufio"
	"log"
	"os"
)

func readLines() []string {
	var lines []string

	file, err := os.Open("./2023/11/input.txt")
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
	// Your code here
	lines := readLines()
	println(lines)
}
