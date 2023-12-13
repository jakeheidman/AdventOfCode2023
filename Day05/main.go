package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type almanacMap struct {
	maps []*conversion
}

func (a *almanacMap) addConversion(c *conversion) {
	a.maps = append(a.maps, c)
}

func (a *almanacMap) seedToLocation(seed int) int {
	convertedSeed := seed
	for _, c := range a.maps {

		newConvertedSeed := c.translateSeed(convertedSeed)
		if newConvertedSeed != convertedSeed {
			return newConvertedSeed
		}
	}
	return convertedSeed
}

type conversion struct {
	sourceStart int
	destStart   int
	length      int
}

func (c *conversion) translateSeed(sourceSeed int) int {
	if !(helpers.Between(c.sourceStart, c.sourceStart+c.length-1, sourceSeed)) { //not in map
		return sourceSeed
	} else {
		diff := sourceSeed - c.sourceStart
		return c.destStart + diff
	}

}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)
	seeds := getSeeds(input[0])
	input = input[3:]
	almanacs := createAlmanacs(input)
	lowestLocation := math.MaxInt
	for _, s := range seeds {
		updatedValue := s
		for _, a := range almanacs {
			updatedValue = a.seedToLocation(updatedValue)
		}
		lowestLocation = min(lowestLocation, updatedValue)
	}
	return lowestLocation
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

func createAlmanacs(input []string) []*almanacMap {
	var almanac []*almanacMap
	a := new(almanacMap)
	for i := 0; i < len(input); i++ {
		line := input[i]
		if line == "" { //end of a almanac map, so skip the next line which is a map name
			almanac = append(almanac, a)
			i++
			a = new(almanacMap)
		} else {
			c := createConversion(line)
			a.addConversion(c)
		}
	}
	almanac = append(almanac, a)
	return almanac
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
