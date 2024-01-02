package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	part := Part{x: 787, m: 2655, a: 1222, s: 2876}
	got := filter("{m>568:A,R}", part)
	want := "A"
	if want != got {
		t.Errorf("Failure on: %s != %s\n", want, got)
	}
}

func TestNewPart(t *testing.T) {
	got := NewPart("{x=787,m=2655,a=1222,s=2876}")
	want := Part{x: 787, m: 2655, a: 1222, s: 2876}
	if want != got {
		t.Errorf("Failure on: %v != %v\n", want, got)
	}
}

func TestP1ProcessLines(t *testing.T) {
	testRows := []string{
		"px{a<2006:qkq,m>2090:A,rfg}",
		"pv{a>1716:R,A}",
		"lnx{m>1548:A,A}",
		"rfg{s<537:gd,x>2440:R,A}",
		"qs{s>3448:A,lnx}",
		"qkq{x<1416:A,crn}",
		"crn{x>2662:A,R}",
		"in{s<1351:px,qqz}",
		"qqz{s>2770:qs,m<1801:hdj,R}",
		"gd{a>3333:R,R}",
		"hdj{m>838:A,pv}",
		"{x=787,m=2655,a=1222,s=2876}",
		"{x=1679,m=44,a=2067,s=496}",
		"{x=2036,m=264,a=79,s=2244}",
		"{x=2461,m=1339,a=466,s=291}",
		"{x=2127,m=1623,a=2188,s=1013}",
	}

	want := 19114
	got, err := p1ProcessLines(testRows)
	if err != nil {
		t.Errorf("processLines() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}

// func TestP2ProcessLines(t *testing.T) {
// 	testRows := []string{}

// 	want := 46
// 	got, err := p2ProcessLines(testRows)
// 	if err != nil {
// 		t.Errorf("processLines() error = %v", err)
// 		return
// 	}
// 	if got != want {
// 		t.Errorf("processLines() = %v, want %v", got, want)
// 	}
// }
