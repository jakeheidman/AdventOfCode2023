package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type race struct {
	time   int
	record int
}

func makeRace(time, record int) *race {
	r := new(race)
	r.time = time
	r.record = record
	return r
}

func (r *race) numRecordPermutations() int {
	permutations := 0
	for t := 1; t < r.time; t++ {
		remainingTime := r.time - t
		distance := remainingTime * t
		if distance > r.record {
			permutations++
		}
	}
	return permutations
}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)
	times := input[0]
	distances := input[1]
	races := getRaces(times, distances)
	answer := 1
	for _, r := range races {
		p := r.numRecordPermutations()
		answer *= p
	}
	return answer
}

func getRaces(times, distances string) []*race {
	var races []*race
	var allTimes []int
	var allDistances []int
	times = strings.TrimLeft(times, "Time: ")
	distances = strings.TrimLeft(distances, "Distance: ")
	for _, time := range strings.Split(times, " ") {
		if time != "" {
			t, _ := strconv.Atoi(time)
			allTimes = append(allTimes, t)
		}
	}
	for _, dist := range strings.Split(distances, " ") {
		if dist != "" {
			d, _ := strconv.Atoi(dist)
			allDistances = append(allDistances, d)
		}
	}
	for i := 0; i < len(allTimes); i++ {
		newRace := makeRace(allTimes[i], allDistances[i])
		races = append(races, newRace)
	}

	return races
}

func Part2(filename string) int {
	return 0
}

func main() {
	//getRaces("Time:      7  15   30", "")
	fmt.Println(Part1("input.txt"))
}
