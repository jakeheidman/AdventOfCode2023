package main

import "github.com/jakeheidman/AdventOfCode2023/helpers"

type hand struct {
	cards    string
	bet      int
	strength handStrength
}

func (h *hand) calculateHandStrength() {
	//5 of a kind

	//4 of a kind

	//full house

	//3 of a kind

	//2 pair

	//1 pair

	//high card

}

const (
	FiveOfKind  handStrength = 7
	FourOfKind  handStrength = 6
	FullHouse   handStrength = 5
	ThreeOfKind handStrength = 4
	TwoPair     handStrength = 3
	OnePair     handStrength = 2
	HighCard    handStrength = 1
)

type handStrength int

func createHand(cards string, bet int) *hand {
	h := new(hand)
	h.cards = cards
	h.bet = bet
	h.calculateHandStrength()
	return h
}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)

	return 0
}

func Part2(filename string) int {
	return 0
}

func main() {

}
