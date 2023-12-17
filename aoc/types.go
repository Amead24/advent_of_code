package aoc

import "strconv"

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
