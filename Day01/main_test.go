package main

import "testing"

func TestDay01(t *testing.T) {
	got := Part1("test_input.txt")
	want := 142
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
