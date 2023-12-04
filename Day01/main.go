package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ParseInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func isInt(char rune) bool {
	return char >= '0' && char <= '9'
}

func GetCalibrationValue(line string) int {
	var first_digit, last_digit rune
	for _, char := range line {
		if isInt(char) {
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
	input := ParseInput(filename)
	var sum int
	for _, line := range input {
		calibrationValue := GetCalibrationValue(line)
		sum += calibrationValue
	}
	return sum
}

func Part2() int {
	return 0
}

func main() {
	fmt.Println(Part1("input1.txt"))

}
