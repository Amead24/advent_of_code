package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	value string
	left  string
	right string
}

func readInput() [][]string {
	var lines [][]string

	file, err := os.Open("./2023/8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_line := scanner.Text()
		new_line = strings.Replace(new_line, " = (", ", ", -1)
		new_line = strings.Replace(new_line, ")", "", -1)
		lines = append(lines, strings.Split(new_line, ", "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from file:", err)
	}

	return lines
}

func main() {
	// Your code here
	root := readInput()
	fmt.Printf("%v\n", root)

	// 1. create the tree
	// 2. traverse the tree.. until I hit ZZZ
	// 3. return count?
}
