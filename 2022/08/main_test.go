package main

import (
	"testing"
)

func TestP1ProcessLines(t *testing.T) {
	testRows := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	want := 21
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

	want := 0
	got, err := p2ProcessLines(testRows)
	if err != nil {
		t.Errorf("processLines() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
