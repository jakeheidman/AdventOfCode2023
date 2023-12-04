package main

import (
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	got := Day2Part1("test_input1.txt")
	want := 8
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func compareGames(g1 *game, g2 *game) (string, bool) {
	if g1.id != g2.id {
		return "id", false
	}
	if g1.red != g2.red {
		return ("red: " + strconv.Itoa(g1.red) + " vs " + strconv.Itoa(g2.red)), false
	}
	if g1.green != g2.green {
		return "green", false
	}
	if g1.blue != g2.blue {
		return "blue", false
	}
	return "", true
}

func TestLineToGame(t *testing.T) {
	got := LineToGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	want := &game{id: 1, red: 4, green: 2, blue: 6}
	if err, state := compareGames(got, want); !state {
		t.Errorf("field did not match: " + err)
	}
}
