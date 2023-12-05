package main

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 4361 //4361
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("test_input.txt")
	want := 467835
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2b(t *testing.T) {
	got := Part2("test_input2.txt")
	want := 158 * 141
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
