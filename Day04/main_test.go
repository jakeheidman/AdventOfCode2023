package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("test_input.txt")
	want := 13
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCreateCard(t *testing.T) {
	got := CreateCard("Card 1: 41 48 83 | 83 86  6 31 17  9")
	winning_numbers := make([]int, 3)
	winning_numbers[0] = 41
	winning_numbers[1] = 48
	winning_numbers[2] = 83
	player_numbers := make([]int, 6)
	player_numbers[0] = 83
	player_numbers[1] = 86
	player_numbers[2] = 6
	player_numbers[3] = 31
	player_numbers[4] = 17
	player_numbers[5] = 9
	want := &card{card_num: 1, winning_numbers: winning_numbers, player_numbers: player_numbers}
	if got.card_num != want.card_num {
		fmt.Println(got)
		fmt.Println(want)
		t.Errorf("incorrect card num got %d want %d", got.card_num, want.card_num)
	}
	// if got.player_numbers != want.winning_numbers {
	// 	t.Errorf("players numbers do not match")
	// }
}
