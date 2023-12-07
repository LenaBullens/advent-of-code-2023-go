package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

func main() {
	lines := helper.ReadLines("input.txt")

	var hands []Hand
	for _, line := range lines {
		hand := NewHand(line)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a Hand, b Hand) int {
		return compareHands(a, b)
	})

	winnings := 0

	for i := 0; i < len(hands); i++ {
		winnings = winnings + hands[i].Bid*(i+1)
	}

	fmt.Println(winnings)
}

type Card int

const (
	InvalidCard = Card(0)
	Two         = Card(2)
	Three       = Card(3)
	Four        = Card(4)
	Five        = Card(5)
	Six         = Card(6)
	Seven       = Card(7)
	Eight       = Card(8)
	Nine        = Card(9)
	Ten         = Card(10)
	Jack        = Card(11)
	Queen       = Card(12)
	King        = Card(13)
	Ace         = Card(14)
)

func parseCard(input string) Card {
	switch input {
	case "2":
		return Two
	case "3":
		return Three
	case "4":
		return Four
	case "5":
		return Five
	case "6":
		return Six
	case "7":
		return Seven
	case "8":
		return Eight
	case "9":
		return Nine
	case "T":
		return Ten
	case "J":
		return Jack
	case "Q":
		return Queen
	case "K":
		return King
	case "A":
		return Ace
	}
	return InvalidCard
}

type Type int

const (
	FiveOfAKind  = Type(7)
	FourOfAKind  = Type(6)
	FullHouse    = Type(5)
	ThreeOfAKind = Type(4)
	TwoPair      = Type(3)
	OnePair      = Type(2)
	HighCard     = Type(1)
	InvalidType  = Type(0)
)

type Hand struct {
	Cards []Card
	Type  Type
	Bid   int
}

func NewHand(input string) Hand {
	result := Hand{}
	splitString := strings.Split(input, " ")

	var cards []Card
	cardStrings := strings.Split(splitString[0], "")
	for _, cardString := range cardStrings {
		card := parseCard(cardString)
		cards = append(cards, card)
	}
	result.Cards = cards

	bid, err := strconv.Atoi(splitString[1])
	if err != nil {
		log.Fatal(err)
	}
	result.Bid = bid

	result.Type = evaluateType(result.Cards)

	return result
}

func evaluateType(cards []Card) Type {
	// Count the number of uniques
	set := make(map[Card]bool)
	for _, card := range cards {
		set[card] = true
	}
	nbOfUniques := len(set)
	if nbOfUniques == 1 {
		// All the same card => five of a kind
		return FiveOfAKind
	}
	if nbOfUniques == 2 {
		// Two different cards => AAABB (full house) or AAAAB (four of a kind)
		// 1-4 is the same as 4-1 | 2-3 is the same as 3-2

		// Count occurence of one card, if it's 4 or 1 => four of a kind, otherwise full house
		firstCard := InvalidCard
		firstAmount := 0
		for _, card := range cards {
			if firstCard == InvalidCard {
				firstCard = card
			}
			if firstCard == card {
				firstAmount++
			}
		}
		if firstAmount == 1 || firstAmount == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	}
	if nbOfUniques == 3 {
		// Three different cards => AAABC (three of a kind) or AABBC (two pair)

		// Count occurence of two cards, if it's 3 & 1, 1 & 3 or 1 & 1 => three of a kind, otherwise two pair.
		firstCard := InvalidCard
		firstAmount := 0
		secondCard := InvalidCard
		secondAmount := 0
		for _, card := range cards {
			if firstCard == InvalidCard {
				firstCard = card
			}
			if firstCard == card {
				firstAmount++
			} else if secondCard == InvalidCard {
				secondCard = card
			}
			if secondCard == card {
				secondAmount++
			}
		}
		if firstAmount == 3 || secondAmount == 3 || (firstAmount == 1 && secondAmount == 1) {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	}
	if nbOfUniques == 4 {
		// One duplicate => one pair
		return OnePair
	}
	if nbOfUniques == 5 {
		// All different cards => high card
		return HighCard
	}

	return InvalidType
}

func compareHands(a Hand, b Hand) int {
	if a.Type < b.Type {
		return -1
	} else if a.Type > b.Type {
		return 1
	} else {
		for i := 0; i < 5; i++ {
			if a.Cards[i] < b.Cards[i] {
				return -1
			} else if a.Cards[i] > b.Cards[i] {
				return 1
			}
		}
	}
	return 0
}
