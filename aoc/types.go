package aoc

import "strconv"

// converts a list of strings to a list of ints
func MapInt(slice []string) []int {
	sliceInts := make([]int, len(slice))
	// splitLine := strings.Fields(slice[i])
	for i, x := range slice {
		num, err := strconv.Atoi(x)
		if err == nil {
			sliceInts[i] = num
		}
	}
	return sliceInts
}

// get the abs from ints (stdlib only supports floats)
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
