package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// testCards := "" +
	// 	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n" +
	// 	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" +
	// 	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n" +
	// 	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n" +
	// 	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n" +
	// 	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	// cards := strings.Split(testCards, "\n")

	var cards []string

	file, err := os.Open("./2023/4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from file:", err)
	}

	var setsWinners [][]string
	var setsScratchoff [][]string
	// splitCards := strings.Split(cards, "\n")
	// for i, card := range splitCards {
	for i, card := range cards {
		dropGame := strings.SplitN(card, ":", 2)
		contents := strings.SplitN(dropGame[1], "|", 2)

		// Note: Danger - Strings.Split() leaves in blank spaces
		// while Strings.Fields() does not.
		winners := strings.Fields(strings.Trim(contents[0], " "))
		setsWinners = append(setsWinners, winners)

		numbers := strings.Fields(strings.Trim(contents[1], " "))
		setsScratchoff = append(setsScratchoff, numbers)

		fmt.Printf("Game %d: winner={%v}; scrachoff={%v}\n", i, winners, numbers)
	}

	winners := make([]int, len(cards))
	// for i := 0; i < len(splitCards); i++ {
	for i := 0; i < len(cards); i++ {
		fmt.Printf("Start winners: %v\n", winners)
		// not the fastest O(N^2) - but good enough for 200 items
		winners[i]++
		localWinners := 0
		// fmt.Printf("Testing (%d)\n:%v\n%v\n", i, setsWinners[i], setsScratchoff[i])
		for _, w := range setsWinners[i] {
			for _, s := range setsScratchoff[i] {
				if w == s {
					localWinners++
					// fmt.Printf("Winner %d: {%s==%s}, multi=%d\n", i, w, s, multiplier)
				}
			}
		}

		for w := i + 1; w < i+1+localWinners; w++ {
			winners[w] += winners[i]
		}

		fmt.Printf("Stop winners: %v\n", winners)

	}

	sum := 0
	for _, x := range winners {
		sum += x
	}

	fmt.Printf("Final part 2: %d\n", sum)
}
