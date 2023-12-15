package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type hand struct {
	cards    string
	bet      int
	strength handStrength
}

func (h *hand) calculateHandStrength() handStrength {
	charTracker := getCardTracker(h.cards)
	//5 of a kind
	if len(charTracker) == 1 {
		return FiveOfKind
	}
	if len(charTracker) == 2 { //4 of a kind or a full house
		for _, k := range charTracker {
			if k == 1 || k == 4 {
				return FourOfKind
			} else {
				return FullHouse
			}
		}
	}

	pairCount := 0
	for _, v := range charTracker { //since 5, 4 of kind, full house branch paths are already resolved, values can only be 3, 2, 1
		if v == 3 { //3 of a kind
			return ThreeOfKind
		}
		if v == 2 {
			pairCount++
		}
	}
	switch pairCount {
	case 2:
		return TwoPair
	case 1:
		return OnePair
	default:
		return HighCard
	}
}

func (h *hand) calculateJokerHandStrength() handStrength {
	s := h.calculateHandStrength()
	h.strength = s
	charTracker := getCardTracker(h.cards)
	numJokers, ok := charTracker['J']
	if !ok { //no jokers, calculate normally
		return h.calculateHandStrength()
	} else {
		return jokerStrengthConversion(h.strength, numJokers)
	}

}

// custom type so that we can custom sort the hands
type hands []hand

func (h hands) Less(i, j int) bool {
	if h[i].strength == h[j].strength {
		for card := 0; card < 5; card++ {
			if h[i].cards[card] != h[j].cards[card] {
				return cardValue(h[i].cards[card]) < cardValue(h[j].cards[card])
			}
		}
		fmt.Println("error, Less function reaching an impossible case")
		return false //should never be reached
	} else {
		return h[i].strength < h[j].strength
	}
}

func (h hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hands) Len() int {
	return len(h)
}

func getCardTracker(cards string) map[rune]int {
	charTracker := make(map[rune]int)
	for _, c := range cards {
		charTracker[c]++
	}
	return charTracker
}

func cardValue(card byte) int {

	switch string(card) {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "T":
		return 10
	case "J":
		return 1
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		return -1
	}
}

type handStrength int

const (
	FiveOfKind  handStrength = 7
	FourOfKind  handStrength = 6
	FullHouse   handStrength = 5
	ThreeOfKind handStrength = 4
	TwoPair     handStrength = 3
	OnePair     handStrength = 2
	HighCard    handStrength = 1
)

func createHand(cards string, bet int) *hand {
	h := new(hand)
	h.cards = cards
	h.bet = bet
	s := h.calculateHandStrength()
	h.strength = s
	return h
}

func jokerStrengthConversion(strength handStrength, numJokers int) handStrength {
	if numJokers == 0 {
		return strength
	}
	if numJokers >= 4 {
		return FiveOfKind
	}
	if numJokers == 3 {
		if strength == FullHouse {
			return FiveOfKind
		}
		if strength == ThreeOfKind {
			return FourOfKind
		}
	}
	if numJokers == 2 {
		if strength == FullHouse {
			return FiveOfKind
		}
		if strength == ThreeOfKind { //impossible
			fmt.Println("impossible num jokers 2 3 of kind case reached")
			return FourOfKind
		}
		if strength == TwoPair {
			return FourOfKind
		}
		if strength == OnePair {
			return ThreeOfKind
		}
		fmt.Println("num jokers 2 base case incorrectly reached")
	}
	if numJokers == 1 {
		if strength == FourOfKind {
			return FiveOfKind
		}
		if strength == ThreeOfKind {
			return FourOfKind
		}
		if strength == TwoPair {
			return FullHouse
		}
		if strength == OnePair {
			return ThreeOfKind
		}
		if strength == HighCard {
			return OnePair
		}
	}
	return strength
}

func createJokerHand(cards string, bet int) *hand {
	h := new(hand)
	h.cards = cards
	h.bet = bet
	s := h.calculateJokerHandStrength()
	h.strength = s
	return h
}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)
	var h []hand
	for _, line := range input {
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		newHand := createHand(split[0], bet)
		h = append(h, *newHand)
	}
	sort.Sort(hands(h))
	totalHandValues := 0
	for rank, hand := range h {
		totalHandValues += (hand.bet * (rank + 1))
	}

	return totalHandValues
}

func Part2(filename string) int {
	input := helpers.ParseInput(filename)
	var h []hand
	for _, line := range input {
		split := strings.Split(line, " ")
		bet, _ := strconv.Atoi(split[1])
		newHand := createJokerHand(split[0], bet)
		h = append(h, *newHand)
	}
	sort.Sort(hands(h))
	totalHandValues := 0
	for rank, hand := range h {
		totalHandValues += (hand.bet * (rank + 1))
	}

	return totalHandValues
}

func main() {
	fmt.Println(Part2("input.txt"))
}
