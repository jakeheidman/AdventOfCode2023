package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

func convertToSchematic(input []string) [][]byte {
	schematic := [][]byte{}
	for _, line := range input {
		row := []byte(line)
		schematic = append(schematic, row)
	}
	return schematic
}

func checkForAdjacentSymbols(x int, y int, schematic [][]byte) bool {
	minX := max(x-1, 0)
	minY := max(y-1, 0)
	maxX := min(x+1, len(schematic[0])-1)
	maxY := min(y+1, len(schematic)-1)

	if isSymbol(schematic[minY][minX]) {
		return true
	}
	if isSymbol(schematic[minY][x]) {
		return true
	}
	if isSymbol(schematic[minY][maxX]) {
		return true
	}
	if isSymbol(schematic[y][minX]) {
		return true
	}
	if isSymbol(schematic[y][maxX]) {
		return true
	}
	if isSymbol(schematic[maxY][minX]) {
		return true
	}
	if isSymbol(schematic[maxY][x]) {
		return true
	}
	if isSymbol(schematic[maxY][maxX]) {
		return true
	}
	return false
}

func isSymbol(b byte) bool {
	//symbols := []byte{'$', '&', '-', '*', '%', '/', '@', '=', '+'}
	symbols := "$@&*/=-%+#"
	return strings.Contains(symbols, string(b))
}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)
	schematic := convertToSchematic(input)
	var sum int
	var includedNums []int
	for y := 0; y < len(schematic); y++ { //for each row
		var curNumber int
		var isPartNumber bool
		for x := 0; x < len(schematic[0]); x++ { //for each cell in a row
			c := schematic[y][x]
			isDigit := c >= '0' && c <= '9'
			if isDigit && curNumber == 0 { //start of a number (ie 6 in 642.#)
				curNumber, _ = strconv.Atoi(string(c))
				if checkForAdjacentSymbols(x, y, schematic) {
					isPartNumber = true
				}
			} else if isDigit && curNumber > 0 { //non-starting digit of a number (ie 4 or 2 in 642..)
				n, _ := strconv.Atoi(string(c))
				curNumber = curNumber*10 + n
				if checkForAdjacentSymbols(x, y, schematic) {
					isPartNumber = true
				}
			} else if !isDigit && curNumber > 0 { //just finished a number (ie . in 642.#)
				if isPartNumber {
					sum += curNumber
					includedNums = append(includedNums, curNumber)
				}
				curNumber = 0
				isPartNumber = false
			} else { //not a number or directly after a number (ie # in 642.#)
				curNumber = 0
				isPartNumber = false
			}
		}
		if isPartNumber && curNumber > 0 { //needed when a row ends on a valid part number
			sum += curNumber
		}
	}
	return sum
}

func main() {
	fmt.Println(Part1("input.txt"))
}
