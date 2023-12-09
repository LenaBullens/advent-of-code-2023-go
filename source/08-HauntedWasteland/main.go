package main

import (
	"fmt"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	part1()
	part2()
}

type Node struct {
	Left  string
	Right string
}

func part1() {
	lines := helper.ReadLines("input.txt")

	instructions := strings.Split(lines[0], "")

	nodeMap := make(map[string]Node)
	for i := 2; i < len(lines); i++ {
		address := lines[i][:3]
		left := lines[i][7:10]
		right := lines[i][12:15]
		node := Node{
			Left:  left,
			Right: right,
		}
		nodeMap[address] = node
	}

	currentNode := nodeMap["AAA"]
	done := false
	index := 0
	steps := 0

	for !done {
		steps++
		if index >= len(instructions) {
			index = 0
		}
		direction := instructions[index]
		address := ""
		if direction == "R" {
			address = currentNode.Right
		} else {
			address = currentNode.Left
		}
		if address == "ZZZ" {
			done = true
			break
		} else {
			currentNode = nodeMap[address]
			index++
		}
	}

	fmt.Println(steps)
}

func part2() {
	lines := helper.ReadLines("input.txt")

	instructions := strings.Split(lines[0], "")

	var startAddresses []string
	nodeMap := make(map[string]Node)
	for i := 2; i < len(lines); i++ {
		address := lines[i][:3]
		if address[2:3] == "A" {
			startAddresses = append(startAddresses, address)
		}
		left := lines[i][7:10]
		right := lines[i][12:15]
		node := Node{
			Left:  left,
			Right: right,
		}
		nodeMap[address] = node
	}

	var intervals []int
	for _, startAddress := range startAddresses {
		currentNode := nodeMap[startAddress]
		done := false
		index := 0
		steps := 0

		for !done {
			steps++
			if index >= len(instructions) {
				index = 0
			}
			direction := instructions[index]
			address := ""
			if direction == "R" {
				address = currentNode.Right
			} else {
				address = currentNode.Left
			}
			if address[2:3] == "Z" {
				done = true
				break
			} else {
				currentNode = nodeMap[address]
				index++
			}
		}

		intervals = append(intervals, steps)
	}

	fmt.Println(lcm(intervals))
}

func lcm(input []int) int {
	// Least common multiple of a,b,c is equal to least common multiple of lcm(a,b) and c. Expand for more than three numbers.
	if len(input) == 2 {
		return lcm2(input[0], input[1])
	}
	copy := input[:len(input)-1]
	if len(input) > 2 {
		return lcm2(lcm(copy), input[len(input)-1])
	}
	// Should be unreachable
	return -1
}

func lcm2(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
