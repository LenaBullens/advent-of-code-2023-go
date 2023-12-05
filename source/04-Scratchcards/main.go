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
	Amount         int
}

func main() {
	part2()
}

func part1() {
	lines := helper.ReadLines("input.txt")

	cards := parseInput(lines)

	sum := 0

	for _, card := range cards {
		nbOfMatches := nbOfMatches(card)
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

func part2() {
	lines := helper.ReadLines("input.txt")

	cards := parseInput(lines)

	sum := 0

	for i := 0; i < len(cards); i++ {
		nbOfMatches := nbOfMatches(cards[i])
		for j := 1; j <= nbOfMatches; j++ {
			cards[i+j].Amount = cards[i+j].Amount + cards[i].Amount
		}
		sum = sum + cards[i].Amount
	}

	fmt.Println(sum)
}

func parseInput(lines []string) []Card {
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
		card.Amount = 1

		cards = append(cards, card)
	}

	return cards
}

func nbOfMatches(card Card) int {
	nbOfMatches := 0
	for _, number := range card.Numbers {
		if card.WinningNumbers[number] {
			nbOfMatches++
		}
	}
	return nbOfMatches
}
