package main // use the same package name as your code

import "fmt"

func printTestResult(testName string, got, want int) {
	if got == want {
		fmt.Printf("PASSED: %s\n", testName)
	} else {
		fmt.Printf("FAILED: %s - Got %d, Want %d\n", testName, got, want)
	}
}

func testIsGear() {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	test1, positions1 := isGear(lines, 3, 1)           // Position of first '*'
	printTestResult("TestIsGear1", len(positions1), 2) // Expect 2 adjacent numbers

	test2, positions2 := isGear(lines, 7, 9)           // Position of second '*'
	printTestResult("TestIsGear2", len(positions2), 2) // Expect 2 adjacent numbers

	test3, _ := isGear(lines, 0, 0)              // Position of '4' in first line
	printTestResult("TestIsGear3", test3, false) // Expect false (not a gear)
}

testIsGear()

func testCalcGearRatio() {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	// Assuming positions are correctly identified
	positions1 := [][]int{{0, 0}, {2, 2}} // Gear 1
	ratio1 := calcGearRatio(lines, positions1)
	printTestResult("TestCalcGearRatio1", ratio1, 467*35) // 467 x 35

	positions2 := [][]int{{9, 2}, {9, 5}} // Gear 2
	ratio2 := calcGearRatio(lines, positions2)
	printTestResult("TestCalcGearRatio2", ratio2, 664*598) // 664 x 598
}
