package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Map of words to their numeric values
	words := map[string]int{
		"one": 1, "eno": 1, "two": 2, "owt": 2, "three": 3, "eerht": 3,
		"four": 4, "ruof": 4, "five": 5, "evif": 5, "six": 6, "xis": 6,
		"seven": 7, "neves": 7, "eight": 8, "thgie": 8, "nine": 9, "enin": 9,
	}

	sum := 0
	file, err := os.Open("./1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Regular expressions to find numbers or number words
		reFirst := regexp.MustCompile(`(\d)`)
		reSecond := regexp.MustCompile(`(\d)`)

		firstStr := reFirst.FindString(line)
		secondStr := reSecond.FindString(reverseString(line))

		first := getNumber(firstStr, words)
		second := getNumber(secondStr, words)

		sum += (first * 10) + second
		fmt.Printf("line=%s; first=%d; second=%d; sum=%d;\n", line, first, second, sum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

// Helper function to get number from string
func getNumber(s string, words map[string]int) int {
	if val, ok := words[s]; ok {
		return val
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Error converting string to number:", err)
	}
	return num
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
