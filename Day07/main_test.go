package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 6440
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCalulateHandStrength(t *testing.T) {
	h := createHand("32T3K", 765)
	h.calculateHandStrength()
	got := h.strength
	want := OnePair
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("test_input.txt")
	want := 5905
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
