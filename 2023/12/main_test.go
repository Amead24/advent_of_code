package main

import (
	"testing"
)

func TestIsValidArrangement(t *testing.T) {
	testCases := []struct {
		str          string
		arrangements []int
		want         bool
	}{
		{"#.#.###", []int{1, 1, 3}, true},
		{".#...#....###.", []int{1, 1, 3}, true},
		{".#.###.#.######", []int{1, 3, 1, 6}, true},
		{"####.#...#...", []int{4, 1, 1}, true},
		{"#....######..#####.", []int{1, 6, 5}, true},
		{".###.##....#", []int{3, 2, 1}, true},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		got := isValidArrangement(tc.str, tc.arrangements)
		if got != tc.want {
			t.Errorf("isValidArrangement(%s, %v) = %t; want %t", tc.str, tc.arrangements, got, tc.want)
		}
	}
}

func TestP1ProcessLines(t *testing.T) {
	testCases := []struct {
		input []string
		want  int
	}{
		{[]string{"???.### 1,1,3"}, 1},
		{[]string{".??..??...?##. 1,1,3"}, 4},
		{[]string{"?#?#?#?#?#?#?#? 1,3,1,6"}, 1},
		{[]string{"????.#...#... 4,1,1"}, 1},
		{[]string{"????.######..#####. 1,6,5"}, 4},
		{[]string{"?###???????? 3,2,1"}, 10},
	}

	for _, tc := range testCases {
		got, _ := p1ProcessLines(tc.input)
		if got != tc.want {
			t.Errorf("p1() = %d; want %d", got, tc.want)
		}
	}
}
