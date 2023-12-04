package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

func runeIsInt(char rune) bool {
	return char >= '0' && char <= '9'
}

func GetCalibrationValue(line string) int {
	var first_digit, last_digit rune
	for _, char := range line {
		if runeIsInt(char) {
			if first_digit == 0 {
				first_digit = char
			}
			last_digit = char

		}
	}
	number, _ := strconv.Atoi(string(first_digit) + string(last_digit))
	return number
}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)
	var sum int
	for _, line := range input {
		calibrationValue := GetCalibrationValue(line)
		sum += calibrationValue
	}
	return sum
}

func in(haystack []byte, needle byte) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}

func byteIsInt(b byte) bool {
	return b >= '0' && b <= '9'
}

// example input -> two1nine
// example output -> 2
func getWordNumber(restOfLine string) (int, error) {
	numberTranslation := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}
	numberWord := ""
	for i := 0; i < len(restOfLine); i++ {
		numberWord = numberWord + string(restOfLine[i])
		completeWord, status := subStrInElems(numberWord)
		if status == 2 { //numberWord is a valid spelling of a number
			return numberTranslation[completeWord], nil
		}
		if status == 0 { //numberWord is invalid
			return 0, errors.New("not a valid spelling of a number")
		}
		//numberWord is a substring of a number spelling

	}
	return 0, errors.New("function reached the end without resolving the line")
}

// searches haystacks to see if needle is a substring in one.
// returns the haystack, as well as a code that represents if it is
// a substring (1), or the entire haystack (2). 0 means not found
func subStrInElems(needle string) (string, int) {
	haystacks := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, haystack := range haystacks {
		if haystack == needle {
			return haystack, 2
		}
		if strings.Contains(haystack, needle) {
			return haystack, 1
		}
	}
	return "", 0
}

func GetUpdatedCalibrationValue(line string) int {

	starterChars := []byte{'o', 't', 'f', 's', 'e', 'n',
		'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var first_num, last_num int
	for i := 0; i < len(line); i++ {
		c := line[i]
		if byteIsInt(c) {
			if first_num == 0 {
				first_num, _ = strconv.Atoi(string(c))
			}
			last_num, _ = strconv.Atoi(string(c))
		} else {
			if in(starterChars, c) {
				validNumber, err := getWordNumber(line[i:])
				if err == nil {
					if first_num == 0 {
						first_num = validNumber
					}
					last_num = validNumber
				}
			}
		}
	}

	return first_num*10 + last_num
}

func Part2(filename string) int {
	input := helpers.ParseInput(filename)
	var sum int
	for _, line := range input {
		sum += GetUpdatedCalibrationValue(line)
	}
	return sum
}

func main() {
	fmt.Println(Part1("input.txt"))
	fmt.Println(Part2("input.txt"))
}
