package main

import (
	"fmt"
	"math"
)

func main() {
	time := [...]int{40, 82, 84, 92}
	distances := [...]int{233, 1011, 1110, 1487}

	// Brute Force
	// totalPossibleWins := 1
	// for i, t := range time {
	// 	raceWins := 0
	// 	for ii := 1; ii < t; ii++ {
	// 		distance := ii * (t - ii)
	// 		if distance > distances[i] {
	// 			raceWins++
	// 		}
	// 	}
	// 	totalPossibleWins *= raceWins
	// 	fmt.Printf("Wins this race %d; Total: %d\n", raceWins, totalPossibleWins)
	// }

	// But this is actually just solving for a range in a parabola
	// distance = x(t-x)  becomes  0 = -x^2 + tx - distance
	quad := func(a float64, b, c int) (float64, float64) {
		B, C := float64(b), float64(c)
		return (-B - math.Sqrt(B*B-4*a*C)) / (2 * a), (-B + math.Sqrt(B*B-4*a*C)) / (2 * a)
	}

	raceWins := 1
	for i, t := range time {
		min, max := quad(-1, t, -distances[i])

		// Then when we know the min and max distance we need to find the diff
		raceWins *= int(math.Abs(math.Floor(max) - math.Floor(min)))
		fmt.Printf("After race[%d], Total race wins: %d\n", i, raceWins)
	}
}
