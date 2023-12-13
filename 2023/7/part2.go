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
	cType string
}

func parseCardType(card string) string {
	counts := make(map[rune]int)
	for _, c := range card {
		counts[c]++
	}

	jokerCount, hasJoker := counts['1']

	switch len(counts) {
	case 1: // All the same or all jokers
		return "five-of-a-kind"
	case 2:
		// had to do two loops to avoid the '1' being first and returning early
		for _, v := range counts { // 'T55J5' vs '1111P'
			if v+jokerCount == 5 { // Four of a kind or full house with jokers
				return "five-of-a-kind"
			}
		}

		for _, v := range counts { // 'T55J5' vs '1111P'
			if (v == 4) || (v == 3 && hasJoker) { // Four of a kind with or without joker - '1KQ11'
				return "four-of-a-kind"
			}
		}

		return "full-house" // Two different cards, one must be a pair at least
	case 3:
		for _, v := range counts { // '1' will always be first -
			if (v + jokerCount) == 4 { // ex: ZXX11
				return "four-of-a-kind"
			}
		}

		for k, v := range counts { // have to worry ex: ZZ1VV
			if k != '1' {
				if (v + jokerCount) == 3 { // Three of a kind or two pair with jokers
					return "full-house"
				}
			}
		}

		for k, v := range counts { // have to worry ex: ZZ1VV
			if k != '1' {
				if (v + jokerCount) >= 3 { // Three of a kind or two pair with jokers
					return "three-of-a-kind"
				}
			}
		}

		return "two-pair"
	case 4:
		if hasJoker { // One pair with a joker
			return "three-of-a-kind"
		}

		return "one-pair"
	default: // All different cards, possibly with a joker
		if hasJoker {
			return "one-pair" // Joker makes a pair with the highest card
		}

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
		for old, new := range map[string]string{"A": "Z", "K": "Y", "Q": "X", "T": "V", "J": "1"} {
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

		buckets[cardType] = append(buckets[cardType], cardIndexPair{card: pair[0], index: i, bid: bid, cType: cardType})
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
		fmt.Printf("card: %s, type: %s, index: %d, bid: %d, sum: %d\n", combinedCards[i].card, combinedCards[i].cType, i, combinedCards[i].bid, sum)
	}

	// 251,123,159 -- too high
	// 250,569,264 -- too low
	fmt.Printf("Sum: %d", sum)
}
