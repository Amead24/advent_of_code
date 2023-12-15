package main

import (
	"testing"
)

func TestP1ProcessLines(t *testing.T) {
	testRows := []string{}

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
