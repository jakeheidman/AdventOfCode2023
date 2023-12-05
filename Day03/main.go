package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type schematic_number struct {
	value     int
	start_idx int
	end_idx   int
}

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

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func getNumbersFromRow(row []byte) []schematic_number {
	var schematic_numbers []schematic_number
	var buildingNumber bool
	curNumber := 0
	var idx_start int
	for x := 0; x < len(row); x++ {
		cell := row[x]
		if isDigit(cell) && !buildingNumber {
			buildingNumber = true
			curNumber, _ = strconv.Atoi(string(cell))
			idx_start = x
		} else if isDigit(cell) && buildingNumber {
			cellValue, _ := strconv.Atoi(string(cell))
			curNumber = curNumber*10 + cellValue
		} else if !isDigit(cell) && buildingNumber {
			new_schematic_num := schematic_number{value: curNumber, start_idx: idx_start, end_idx: x - 1}
			schematic_numbers = append(schematic_numbers, new_schematic_num)
			idx_start = 0
			buildingNumber = false
			curNumber = 0
		}
	}
	if buildingNumber {
		new_schematic_num := schematic_number{value: curNumber, start_idx: idx_start, end_idx: len(row) - 1}
		schematic_numbers = append(schematic_numbers, new_schematic_num)
	}
	return schematic_numbers
}

func getValidPartNumbers(nums []schematic_number, x int) []int {
	var validPartNums []int
	for _, schemNum := range nums {
		if (x-1 >= schemNum.start_idx && x-1 <= schemNum.end_idx) || (x >= schemNum.start_idx && x <= schemNum.end_idx) || (x+1 >= schemNum.start_idx && x+1 <= schemNum.end_idx) {
			validPartNums = append(validPartNums, schemNum.value)
		}
	}
	return validPartNums
}

func Part2(filename string) int {
	input := helpers.ParseInput(filename)
	schematic := convertToSchematic(input)
	savedNumberInfo := make([][]schematic_number, len(schematic))
	const GEAR = '*'
	for y := 0; y < len(schematic); y++ {
		savedNumberInfo[y] = getNumbersFromRow(schematic[y])
	}
	var sum int
	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[y]); x++ {
			cell := schematic[y][x]
			if cell == GEAR {
				var validPartsAbove, validPartsBelow []int
				if y > 0 {
					validPartsAbove = getValidPartNumbers(savedNumberInfo[y-1], x)
				}
				if y < len(schematic)-1 {
					validPartsBelow = getValidPartNumbers(savedNumberInfo[y+1], x)
				}
				var leftNumBuilder, rightNumBuilder string
				var leftNum, rightNum int
				var leftExists, rightExists bool
				if x > 0 && isDigit(schematic[y][x-1]) { //calculate left part number
					leftExists = true
					for xIter := x - 1; xIter > 0; xIter-- {
						leftCell := schematic[y][xIter]
						if !isDigit(leftCell) {
							break
						}
						leftNumBuilder = string(leftCell) + leftNumBuilder
					}
					leftNum, _ = strconv.Atoi(leftNumBuilder)
				}
				if x < len(schematic[y])-1 && isDigit(schematic[y][x+1]) { //calculate right part number
					rightExists = true
					for xIter := x + 1; xIter < len(schematic[y]); xIter++ {
						rightCell := schematic[y][xIter]
						if !isDigit(rightCell) {
							break
						}
						rightNumBuilder = rightNumBuilder + string(rightCell)
					}
					rightNum, _ = strconv.Atoi(rightNumBuilder)
				}
				allValidPartNumbers := append(validPartsAbove, validPartsBelow...)
				if leftExists {
					allValidPartNumbers = append(allValidPartNumbers, leftNum)
				}
				if rightExists {
					allValidPartNumbers = append(allValidPartNumbers, rightNum)
				}
				if len(allValidPartNumbers) == 2 {
					sum += allValidPartNumbers[0] * allValidPartNumbers[1]
				}

			}
		}
	}
	return sum
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
