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

func Part1(filename string) int {
	return 0
}
func main() {
	fmt.Println(Part1("input.txt"))
}
