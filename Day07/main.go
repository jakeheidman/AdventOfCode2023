package main

type hand struct {
	cards    string
	bet      int
	strength handStrength
}

func (h *hand) calculateHandStrength() handStrength {
	charTracker := make(map[rune]int)
	for _, c := range h.cards {
		charTracker[c]++
	}
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
	s := h.calculateHandStrength()
	h.strength = s
	return h
}

func Part1(filename string) int {
	//input := helpers.ParseInput(filename)

	return 0
}

func Part2(filename string) int {
	return 0
}

func main() {

}
