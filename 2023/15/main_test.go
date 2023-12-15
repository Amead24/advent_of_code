package main

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
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
	testRows := []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}

	want := 1320
	got, err := p1ProcessLines(testRows)
	if err != nil {
		t.Errorf("processLines() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}

func TestBoxExtractLabel(t *testing.T) {
	testCases := []struct {
		input string
		want  Lense
	}{
		{"rn=1", Lense{label: "rn", focalLength: 1}},
		{"cm-", Lense{label: "cm", focalLength: 0}},
		{"qp=3", Lense{label: "qp", focalLength: 3}},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		got := extractBoxLabel(tc.input)
		if got != tc.want {
			t.Errorf("extractBoxLabel(%s) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestCreateHashMap(t *testing.T) {
	testRows := []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}

	got := createHashMap(testRows)
	want := map[int][]Lense{
		0: {Lense{"rn", 1}, Lense{"cm", 2}},
		3: {Lense{"ot", 7}, Lense{"ab", 5}, Lense{"pc", 6}},
	}

	// todo how compare structs?
	fmt.Printf("got: %v, wanted: %v\n", got, want)
}

func TestP2ProcessLines(t *testing.T) {
	testRows := []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}

	want := 145
	got, err := p2ProcessLines(testRows)
	if err != nil {
		t.Errorf("processLines() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
