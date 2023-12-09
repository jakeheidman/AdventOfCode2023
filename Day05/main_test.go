package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 35
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCreateConversion(t *testing.T) {
	got := createConversion("50 98 2")
	want := new(conversion)
	want.destStart = 50
	want.sourceStart = 98
	want.length = 2
	if got.destStart != want.destStart {
		t.Errorf("dest error: got %d want %d", got.destStart, want.destStart)
	}
	if got.sourceStart != want.sourceStart {
		t.Errorf("source error: got %d want %d", got.sourceStart, want.sourceStart)
	}
	if got.length != want.length {
		t.Errorf("length error: got %d want %d", got.length, want.length)
	}
}

func TestTranslateSeed(t *testing.T) {
	c := createConversion("50 98 2")
	seed := 99
	got := c.translateSeed(seed)
	want := 51
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
