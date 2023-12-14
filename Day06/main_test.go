package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

}

func TestNumRecordPermutations(t *testing.T) {
	r := makeRace(7, 9)
	got := r.numRecordPermutations()
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
