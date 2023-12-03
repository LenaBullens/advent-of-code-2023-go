package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	part1()
}

func part1() {
	lines := helper.ReadLines("input.txt")

	sum := 0

	for _, line := range lines {
		numberString := findFirstDigit(line) + findLastDigit(line)
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + number
	}

	fmt.Println(sum)
}

func findFirstDigit(input string) string {
	for _, rune := range input {
		if unicode.IsDigit(rune) {
			return string(rune)
		}
	}
	return ""
}

func findLastDigit(input string) string {
	runes := []rune(input)
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			return string(runes[i])
		}
	}
	return ""
}
