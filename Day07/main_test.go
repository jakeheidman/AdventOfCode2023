package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

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
