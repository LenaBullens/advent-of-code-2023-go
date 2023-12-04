package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Card struct {
	WinningNumbers map[int]bool
	Numbers        []int
}

func main() {
	lines := helper.ReadLines("input.txt")
	var cards []Card

	for _, line := range lines {
		card := Card{}

		splitString1 := strings.Split(line, "|")
		splitString2 := strings.Split(splitString1[0], ":")

		winningNumbers := make(map[int]bool)
		winningNumbersString := strings.TrimSpace(splitString2[1])
		splitString3 := strings.Split(winningNumbersString, " ")
		for _, winningNumberString := range splitString3 {
			if winningNumberString != "" {
				winningNumber, err := strconv.Atoi(strings.TrimSpace(winningNumberString))
				if err != nil {
					log.Fatal(err)
				}
				winningNumbers[winningNumber] = true
			}
		}
		card.WinningNumbers = winningNumbers

		var numbers []int
		numbersString := strings.TrimSpace(splitString1[1])
		splitString4 := strings.Split(numbersString, " ")
		for _, numberString := range splitString4 {
			if numberString != "" {
				number, err := strconv.Atoi(strings.TrimSpace(numberString))
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, number)
			}
		}
		card.Numbers = numbers

		cards = append(cards, card)
	}

	sum := 0

	for _, card := range cards {
		nbOfMatches := 0
		for _, number := range card.Numbers {
			if card.WinningNumbers[number] {
				nbOfMatches++
			}
		}
		score := 0
		if nbOfMatches >= 1 {
			score = 1
		}
		for i := 0; i < (nbOfMatches - 1); i++ {
			score = score * 2
		}

		sum = sum + score
	}

	fmt.Println(sum)
}
