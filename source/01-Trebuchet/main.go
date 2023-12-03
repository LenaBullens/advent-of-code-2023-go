package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := helper.ReadLines("input.txt")

	sum := 0

	for _, line := range lines {
		firstDigit, _ := findFirstDigit(line)
		lastDigit, _ := findLastDigit(line)
		numberString := firstDigit + lastDigit
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + number
	}

	fmt.Println(sum)
}

func findFirstDigit(input string) (string, int) {
	for i, rune := range input {
		if unicode.IsDigit(rune) {
			return string(rune), i
		}
	}
	return "", -1
}

func findLastDigit(input string) (string, int) {
	runes := []rune(input)
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			return string(runes[i]), i
		}
	}
	return "", -1
}

func part2() {
	lines := helper.ReadLines("input.txt")

	words := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0

	for _, line := range lines {
		firstDigit := ""
		lastDigit := ""

		firstNumericDigit, firstNumericPos := findFirstDigit(line)
		firstWordDigit, firstWordPos := findFirstWord(line, words)
		if firstNumericPos < firstWordPos {
			firstDigit = firstNumericDigit
		} else {
			firstDigit = firstWordDigit
		}

		lastNumericDigit, lastNumericPos := findLastDigit(line)
		lastWordDigit, lastWordPos := findLastWord(line, words)
		if lastNumericPos > lastWordPos {
			lastDigit = lastNumericDigit
		} else {
			lastDigit = lastWordDigit
		}

		numberString := firstDigit + lastDigit
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + number
	}

	fmt.Println(sum)
}

func findFirstWord(input string, words [9]string) (string, int) {
	digit := ""
	pos := math.MaxInt
	for i := 0; i < len(words); i++ {
		idx := strings.Index(input, words[i])
		if idx != -1 {
			if idx < pos {
				pos = idx
				digit = strconv.Itoa(i + 1)
			}
		}
	}
	return digit, pos
}

func findLastWord(input string, words [9]string) (string, int) {
	digit := ""
	pos := -1
	for i := 0; i < len(words); i++ {
		idx := strings.LastIndex(input, words[i])
		if idx != -1 {
			if idx > pos {
				pos = idx
				digit = strconv.Itoa(i + 1)
			}
		}
	}
	return digit, pos
}
