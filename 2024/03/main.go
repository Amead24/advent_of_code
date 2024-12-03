package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func p1ProcessLines(lines []string) (int, error) {
	validCalcs := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
	newishSrc := validCalcs.FindAllStringSubmatch(lines[0], -1)

	total := 0
	for _, match := range newishSrc {
		fmt.Println(match)
		values := aoc.MapInt(strings.Split(match[1], ","))
		total += values[0] * values[1]
	}
	return total, nil
}

func p2ProcessLines(lines []string) (int, error) {
	// go doesn't support negative lookbacks
	//validCalcs := regexp.MustCompile(`(?<!don't\(\).*?)mul\((\d,\d)\)|do\(\).*?mul\((\d,\d)\)`)
	//newishSrc := validCalcs.FindAllStringSubmatch(lines[0], -1)
	println("------")

	total := 0
	validCalcs := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
	newishSrc := validCalcs.FindAllStringIndex(lines[0], -1)

	// somehow i think it'd make sense to track the last before a do()
	//  and create substrings that then do the matches
	// xmul(1,2)dont()mul(3,3)mul(5,5)do()mul(8,8)

	isEnabled := true
	lastEnabledMatchIndex := 0
	substring := lines[0][lastEnabledMatchIndex:newishSrc[0][1]]
	for i := 0; i < len(newishSrc); i = i + 1 {
		// first - create a substring from everything to the start until the end of the first match

		// now we'll consider everything valid until we hit a dont()
		if regexp.MustCompile(`don't\(\)`).Match([]byte(substring)) {
			// and then disable calcs
			isEnabled = false
		}

		// until we hit a do()
		if regexp.MustCompile(`do\(\)`).Match([]byte(substring)) {
			isEnabled = true
		}

		if isEnabled { // in which case we add it to the total
			match := aoc.MapInt(strings.Split(validCalcs.FindStringSubmatch(substring)[1], ","))
			total += match[0] * match[1]
		}

		// now update what is considered the end of the substring index, based on what we just checked
		lastEnabledMatchIndex = newishSrc[i][1]

		// and grow the substring to match the end of the last index to the end of the next one
		//  unless it's the last match, go to the very end
		if i == len(newishSrc) - 1 {
			substring = lines[0][lastEnabledMatchIndex:len(lines[0])]
		} else {
			substring = lines[0][lastEnabledMatchIndex:newishSrc[i+1][1]]
		}
	}
	return total, nil
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	sum, err := p1ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 1 Error: %e\n", err)
	}

	fmt.Printf("Part 1 Sum = %d\n", sum)

	sum, err = p2ProcessLines(lines)
	if err != nil {
		fmt.Errorf("Part 2 Error: %e\n", err)
	}

	fmt.Printf("Part 2 Sum = %d\n", sum)
}
