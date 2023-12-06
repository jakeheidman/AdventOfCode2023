package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type card struct {
	card_num        int
	winning_numbers []int
	player_numbers  []int
}

func (c *card) getPoints() int {
	numberOfWinners := 0
	for _, cardNum := range c.player_numbers {
		for _, winningNum := range c.winning_numbers {
			if cardNum == winningNum {
				numberOfWinners += 1
				break
			} else if cardNum < winningNum {
				break
			}
		}
	}
	if numberOfWinners == 0 {
		return 0
	} else {
		return int(math.Pow(2, float64(numberOfWinners-1)))
	}
}

func CreateCard(line string) *card {
	split := strings.Split(line, "|")
	cardProperties := split[0]
	cardPlayerNumbers := split[1]
	newCard := &card{}
	setCardProperties(newCard, cardProperties)
	setCardPlayerNumbers(newCard, cardPlayerNumbers)
	return newCard
}

// Sets card_num and winning_numbers of a card
// input: card, string -> Card   1: 13  4 61 82 80 41 31 53 50  2
func setCardProperties(newCard *card, properties string) {
	split := strings.Split(properties, ":") //["Card   1", " 13  4 61 82 80 41 31 53 50  2 "]
	cardNum := strings.Replace(split[0], "Card", "", 1)
	cardNum = strings.TrimSpace(cardNum)
	newCard.card_num, _ = strconv.Atoi(cardNum)
	winningNumbers := strings.TrimSpace(split[1])
	winningNumbers = strings.Replace(winningNumbers, "  ", " ", -1)
	numSlice := strings.Split(winningNumbers, " ")
	for _, num := range numSlice {
		winningNum, _ := strconv.Atoi(num)
		newCard.winning_numbers = append(newCard.winning_numbers, winningNum)
	}
	sort.Ints(newCard.winning_numbers)
}

func setCardPlayerNumbers(card *card, cardPlayerNumbers string) {
	cardPlayerNumbers = strings.TrimSpace(cardPlayerNumbers)
	cardPlayerNumbers = strings.Replace(cardPlayerNumbers, "  ", " ", -1)
	playerNums := strings.Split(cardPlayerNumbers, " ")
	for _, num := range playerNums {
		n, _ := strconv.Atoi(num)
		card.player_numbers = append(card.player_numbers, n)
	}
	sort.Ints(card.winning_numbers)
}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)
	var sum int
	for _, line := range input {
		card := CreateCard(line)
		sum += card.getPoints()
	}
	return sum
}

func main() {
	fmt.Println(Part1("input.txt"))
}
