package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Game struct {
	Id    int
	Hands []Hand
}

type Hand struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	input := helper.ReadLines("input.txt")

	var games []Game
	for _, line := range input {
		game := parseLine(line)
		games = append(games, game)
	}

	sum := 0
	for _, game := range games {
		valid, id := evaluateGame(game, 12, 13, 14)
		if valid {
			sum = sum + id
		}
	}

	fmt.Println(sum)
}

func parseLine(line string) Game {
	game := Game{}
	split1 := strings.Split(line, ":")

	gameString := split1[0]
	idString := gameString[5:] //Input only contains utf-8 compatible characters so no need to work with runes.
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal(err)
	}
	game.Id = id

	var hands []Hand
	handsString := split1[1]
	split2 := strings.Split(handsString, ";")
	for _, handString := range split2 {
		hand := Hand{}
		split3 := strings.Split(handString, ",")
		for _, cubeString := range split3 {
			cubeString = strings.TrimLeft(cubeString, " ")
			split4 := strings.Split(cubeString, " ")
			if split4[1] == "red" {
				hand.Red, err = strconv.Atoi(split4[0])
				if err != nil {
					log.Fatal(err)
				}
			} else if split4[1] == "green" {
				hand.Green, err = strconv.Atoi(split4[0])
				if err != nil {
					log.Fatal(err)
				}
			} else if split4[1] == "blue" {
				hand.Blue, err = strconv.Atoi(split4[0])
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		hands = append(hands, hand)
	}
	game.Hands = hands

	return game
}

func evaluateGame(game Game, red int, green int, blue int) (bool, int) {
	for _, hand := range game.Hands {
		if (hand.Red > red) || (hand.Green > green) || (hand.Blue > blue) {
			return false, 0
		}
	}
	return true, game.Id
}
