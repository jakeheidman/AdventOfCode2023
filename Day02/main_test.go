package main

import "testing"

func TestPart1(t *testing.T) {
	got := Day2Part1("test_input1.txt")
	want := 8
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
