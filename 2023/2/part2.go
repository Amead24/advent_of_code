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

	// return 0, fmt.Errorf("color %s not found", color)

	// default to 0
	return 0, nil
}

func main() {
	lines := make(chan string)
	go readLines("./2023/2/input.txt", lines)

	sum := 0
	for line := range lines {
		reGameNumber := regexp.MustCompile(`^Game (\d+):`)
		gameNumberMatches := reGameNumber.FindStringSubmatch(line)

		counts := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		hands := strings.Split(line[len(gameNumberMatches[0]):], ";")
		fmt.Printf("hands=%v\n", hands)
		for _, hand := range hands {
			hand = strings.TrimSpace(hand)
			for color := range counts {
				count, err := parseColorCount(hand, color)
				if err == nil {
					if count >= counts[color] {
						counts[color] = count
					}
				}
			}
		}

		sum += (counts["red"] * counts["blue"] * counts["green"])
		fmt.Printf("GN: %d, counts: %v\n", sum, counts)
	}

	fmt.Println("Final Part 2:", sum)
}
