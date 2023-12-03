package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// readLines reads lines from a given file and sends them through a channel.
func readLines(filePath string, lines chan<- string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}

	close(lines) // Close the channel after all lines are sent
}

func parseColorCount(hand, color string) (int, error) {
	re := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, color))
	matches := re.FindStringSubmatch(hand)
	if len(matches) > 1 {
		// FindStringSubmatch => [whole string, match]
		return strconv.Atoi(matches[1])
	}
	return 0, fmt.Errorf("color %s not found", color)
}

func main() {
	lines := make(chan string)
	go readLines("./2023/2/input.txt", lines)

	sum := 0
	for line := range lines {
		reGameNumber := regexp.MustCompile(`^Game (\d+):`)
		gameNumberMatches := reGameNumber.FindStringSubmatch(line)
		gameNumber, err := strconv.Atoi(gameNumberMatches[1])
		if err != nil {
			log.Fatalf("Error converting game number: %v", err)
		}

		possibleGame := true
		hands := strings.Split(line[len(gameNumberMatches[0]):], ";")
		fmt.Printf("hands=%v\n", hands)
		for _, hand := range hands {
			hand = strings.TrimSpace(hand)
			redCount, err := parseColorCount(hand, "red")
			if err == nil && redCount > 12 {
				possibleGame = false
				break
			}

			blueCount, err := parseColorCount(hand, "green")
			if err == nil && blueCount > 13 {
				possibleGame = false
				break
			}

			greenCount, err := parseColorCount(hand, "blue")
			if err == nil && greenCount > 14 {
				possibleGame = false
				break
			}
		}

		if possibleGame {
			sum += gameNumber
			fmt.Printf("\nAdding game %d to sum, current sum: %d\n", gameNumber, sum)
		} else {
			fmt.Printf("\nGame %d is not possible\n", gameNumber)
		}
	}

	fmt.Println("Final Part 1:", sum)
}
