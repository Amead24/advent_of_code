package main

import "fmt"

func main() {
	// time := [...]int{40, 82, 84, 92}
	// distances := [...]int{233, 1011, 1110, 1487}

	time := [...]uint64{40828492}
	distances := [...]uint64{233101111101487}

	var totalPossibleWins uint64 = 1
	for i, t := range time {
		var raceWins uint64 = 0
		for ii := uint64(1); ii < t; ii++ {
			distance := ii * (t - ii)
			if distance > distances[i] {
				raceWins++
			}
		}
		totalPossibleWins *= raceWins
		fmt.Printf("Wins this race %d; Total: %d\n", raceWins, totalPossibleWins)
	}
}
