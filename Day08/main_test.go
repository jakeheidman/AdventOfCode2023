package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 2
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart1b(t *testing.T) {
	got := Part1("test_input2.txt")
	want := 6
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// func TestPart2(t *testing.T) {
// 	got := Part2("test_input.txt")
// 	want := 5905
// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }
