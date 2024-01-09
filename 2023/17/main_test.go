package main

import (
	"testing"
)

func TestP1ProcessLines(t *testing.T) {
	testRows := []string{
		"2413432311322",
		"3215453535622",
		"3255245654253",
		"3446585845451",
		"4546657867535",
		"1438598798453",
		"4457876987765",
		"3637877979652",
		"4654967986886",
		"4564679986452",
		"1224686865562",
		"2546548887734",
		"4322674655532",
	}

	want := 102
	got, err := p1ProcessLines(testRows)
	if err != nil {
		t.Errorf("processLines() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}

func TestP2ProcessLines(t *testing.T) {
	testRows := []string{}

	want := 46
	got, err := p2ProcessLines(testRows)
	if err != nil {
		t.Errorf("processLines() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
