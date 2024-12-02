package main

import (
	"fmt"
	"strings"

	"github.com/amead24/advent_of_code/aoc"
)

func isReportSafe(report []int) bool {
	isSafe := true
	reportDiffs := make([]int, len(report) - 1)
	for i := 0; i < len(report) - 1; i++ {
		reportDiffs[i] = (report[i+1] - report[i])
	}

	// Determine direction: increasing or decreasing
	isIncreasing := false
	isDecreasing := false
	for _, diff := range reportDiffs {
		if diff > 0 {
			isIncreasing = true
			break
		} else if diff < 0 {
			isDecreasing = true
			break
		}
	}

	for _, diff := range reportDiffs {
		if aoc.AbsInt(diff) < 1 || aoc.AbsInt(diff) > 3 {
			isSafe = false
			break
		}

		if isIncreasing && diff <= 0 {
			isSafe = false
			break
		}

		if isDecreasing && diff >= 0 {
			isSafe = false
			break
		}
	}

	return isSafe
}

func p1ProcessLines(lines []string) (int, error) {
	safeReports := 0
	for _, line := range lines {
		report := aoc.MapInt(strings.Split(line, " "))
		isSafe := isReportSafe(report)
		if isSafe {
			safeReports++
		}
	}

	return safeReports, nil
}

func p2ProcessLines(lines []string) (int, error) {
	safeReports := 0
	for _, line := range lines {
		report := aoc.MapInt(strings.Split(line, " "))
		if isReportSafe(report) {
			safeReports++
		} else {
			for i := 0; i < len(report); i++ {
				newReport := append([]int{}, report[:i]...)
				newReport = append(newReport, report[i+1:]...)

				if isReportSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports, nil
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
