package main

import (
	"testing"
)

// func TestProcessModule(t *testing.T) {
// 	modules := map[string]Module{}
// 	modules["foo"] = Module{state: "off", desintations: "bar", modType: "flipflop", incomingPulse: "empty"}
// 	modules["bar"] = Module{state: "off", destinations: "inv", modType: "flipflop", incomingPulse: "empty"}
// }

func TestP1ProcessLines(t *testing.T) {
	testRows := []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}

	want := 32000000
	got, err := p1ProcessLines(testRows)
	if err != nil {
		t.Errorf("processLines() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}

	// testRows = []string{
	// 	"broadcaster -> a",
	// 	"%a -> inv, con",
	// 	"&inv -> b",
	// 	"%b -> con",
	// 	"&con -> output",
	// }

	// want = 11687500
	// got, err = p1ProcessLines(testRows)
	// if err != nil {
	// 	t.Errorf("processLines() error = %v", err)
	// 	return
	// }
	// if got != want {
	// 	t.Errorf("processLines() = %v, want %v", got, want)
	// }
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
