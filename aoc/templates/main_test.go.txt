package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
		testCases := []struct {
		input string
		want  int
	}{
		{"rn=1", 30},
		{"cm-", 253},
		{"qp=3", 97},
		{"rn", 0},
		{"cm", 0},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		got := hash(tc.input)
		if got != tc.want {
			t.Errorf("Hash(%q) = %d; want %d", tc.input, got, tc.want)
		}
	}
}

func TestP1ProcessLines(t *testing.T) {
	want := 405
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
