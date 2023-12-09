package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type almanacMap struct {
	maps []conversion
}

type conversion struct {
	sourceStart int
	destStart   int
	length      int
}

func (c *conversion) translateSeed(sourceSeed int) int {
	if !(helpers.Between(c.sourceStart, c.sourceStart+c.length, sourceSeed)) { //not in map
		return sourceSeed
	} else {
		diff := sourceSeed - c.sourceStart
		return c.destStart + diff
	}

}

func Part1(filename string) int {
	return 0
}

func getSeeds(line string) []int {
	var s []int
	line = strings.Trim(line, "seeds: ")
	seeds := strings.Split(line, " ")
	for _, seed := range seeds {
		seedNum, _ := strconv.Atoi(seed)
		s = append(s, seedNum)
	}
	return s

}

func createConversion(line string) *conversion {
	numbers := strings.Split(line, " ")
	c := new(conversion)
	c.destStart, _ = strconv.Atoi(numbers[0])
	c.sourceStart, _ = strconv.Atoi(numbers[1])
	c.length, _ = strconv.Atoi(numbers[2])
	return c
}

func main() {
	fmt.Println(Part1("input.txt"))
}
