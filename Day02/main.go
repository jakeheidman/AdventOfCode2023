package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type game struct {
	id    int
	red   int
	blue  int
	green int
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func LineToGame(line string) *game {
	splitOnGame := strings.Split(line, ": ") //Game 1
	gameSubstring := splitOnGame[0]
	ballsSubstring := splitOnGame[1]
	ballsSubstring = strings.Replace(ballsSubstring, ";", ",", -1)
	var red, green, blue int
	gameID, _ := strconv.Atoi(strings.TrimPrefix(gameSubstring, "Game "))
	ballDraws := strings.Split(ballsSubstring, ", ")
	for _, draw := range ballDraws {
		numBalls, _ := strconv.Atoi(strings.Split(draw, " ")[0])
		if strings.Contains(draw, "blue") {
			blue = max(blue, numBalls)
		} else if strings.Contains(draw, "red") {
			red = max(red, numBalls)
		} else if strings.Contains(draw, "green") {
			green = max(green, numBalls)
		} else {
			log.Fatal("error when trying to decipher single draw: " + draw)
		}
	}
	game := game{id: gameID,
		red:   red,
		blue:  blue,
		green: green}
	return &game
}

func isGameValid(game *game) bool {
	return game.red <= 12 && game.green <= 13 && game.blue <= 14
}

func Day2Part1(filename string) int {
	input := helpers.ParseInput(filename)
	var sum int
	for _, line := range input {
		game := LineToGame(line)
		if isGameValid(game) {
			sum += game.id
		}
	}
	return sum
}

func main() {
	fmt.Println(Day2Part1("input.txt"))
}
