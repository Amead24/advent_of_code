package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cardIndexPair struct {
	card  string
	index int
	bid   int
}

func parseCardType(card string) string {
	counts := make(map[rune]int)
	for _, c := range card {
		counts[c]++
	}
	// fmt.Printf("Card: %s, counts: %v", card, counts)

	if len(counts) == 1 {
		return "five-of-a-kind"
	} else if len(counts) == 2 {
		for _, v := range counts {
			if v == 4 {
				return "four-of-a-kind"
			}
		}

		// only other possibility?
		return "full-house"
	} else if len(counts) == 3 {
		for _, v := range counts {
			if v == 3 {
				return "three-of-a-kind"
			}
		}

		return "two-pair"
	} else if len(counts) == 4 {
		return "one-pair"
	} else {
		return "high-card"
	}
}

func readInput() [][]string {
	var cards [][]string

	file, err := os.Open("./2023/7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_line := strings.Fields(scanner.Text())
		for old, new := range map[string]string{"A": "Z", "K": "Y", "Q": "X", "J": "W", "T": "V"} {
			new_line[0] = strings.Replace(new_line[0], old, new, -1)
		}
		cards = append(cards, new_line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from file:", err)
	}

	return cards
}

func main() {
	cards := readInput()

	var indicies = make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		indicies[i] = i
	}

	priorityOfTypes := [...]string{
		"five-of-a-kind",
		"four-of-a-kind",
		"full-house",
		"three-of-a-kind",
		"two-pair",
		"one-pair",
		"high-card",
	}

	buckets := make(map[string][]cardIndexPair, len(priorityOfTypes))
	for i, pair := range cards {
		cardType := parseCardType(pair[0])
		bid, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatalf("Invalid index: %s", pair[1])
		}
		buckets[cardType] = append(buckets[cardType], cardIndexPair{card: pair[0], index: i, bid: bid})
	}

	for _, v := range buckets {
		// fmt.Printf("sorting: %s, len()==%d\n", k, len(v))
		sort.Slice(v, func(x, y int) bool {
			return v[x].card >= v[y].card
		})
	}

	// combine each bucket in order
	var combinedCards []cardIndexPair
	for _, priorityType := range priorityOfTypes {
		combinedCards = append(combinedCards, buckets[priorityType]...)
		// fmt.Printf("bucket: %s, len()==%d; new len()==%d\n", priorityType, len(buckets[priorityType]), len(combinedCards))
	}

	sum := 0
	for i := 0; i < len(combinedCards); i++ {
		sum += (len(combinedCards) - i) * combinedCards[i].bid
		fmt.Printf("card: %s, index: %d, bid: %d, sum: %d\n", combinedCards[i].card, i, combinedCards[i].bid, sum)
	}

	// 158820451 -- too low
	// 159219844 -- no
	// 201061398 -- too high
	fmt.Printf("Sum: %d", sum)
}
