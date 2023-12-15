package main

import (
	"fmt"
	"strings"

	"github.com/jakeheidman/AdventOfCode2023/helpers"
)

type node struct {
	left  string
	right string
	value string
}

type desertMap map[string]*node

func createNode(left, right, value string) *node {
	n := new(node)
	n.left = left
	n.right = right
	n.value = value
	return n
}

func buildNodeMap(input []string) desertMap {
	m := make(desertMap)
	for _, line := range input {
		split := strings.Split(line, " = ")
		cur := split[0]
		elements := strings.Trim(split[1], "()")
		edgedNodes := strings.Split(elements, ", ")
		n := createNode(edgedNodes[0], edgedNodes[1], cur)
		m[cur] = n
	}
	return m
}

func getFirstNode(line string) string {
	return strings.Split(line, " = ")[0]
}

func traverseDesert(directions, firstNode string, desert desertMap) int {
	fmt.Println("Starting desert traversal")
	numberOfSteps := 0
	curNode := firstNode
	for {
		for _, step := range directions {
			curNode = takeStep(step, curNode, desert)
			numberOfSteps++
			fmt.Printf("number of steps: %d\n", numberOfSteps)
			if curNode == "ZZZ" {
				return numberOfSteps
			}
		}
	}

}

func takeStep(step rune, curNode string, desert desertMap) string {
	k := string(step)
	if k == "R" {
		return desert[curNode].right
	} else {
		return desert[curNode].left
	}
}

func Part1(filename string) int {
	input := helpers.ParseInput(filename)
	directions := input[0]
	mapInput := input[2:]
	firstNode := "AAA"
	nodeMap := buildNodeMap(mapInput)
	numberOfSteps := traverseDesert(directions, firstNode, nodeMap)
	return numberOfSteps
}

func Part2(filename string) int {
	return 0
}

func main() {
	fmt.Println(Part1("input.txt"))
}
