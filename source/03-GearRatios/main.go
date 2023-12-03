package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Point struct {
	x int
	y int
}

func main() {
	part1()
}

func part1() {
	grid := helper.ReadGrid("input.txt")
	yMax := len(grid) - 1
	xMax := len(grid[0]) - 1

	symbolZoneMap := make(map[Point]bool)
	for y, row := range grid {
		for x, tile := range row {
			if isSymbol(tile) {
				symbolZoneMap[Point{x, y}] = true
				if x > 0 && y > 0 {
					symbolZoneMap[Point{x - 1, y - 1}] = true
				}
				if y > 0 {
					symbolZoneMap[Point{x, y - 1}] = true
				}
				if x < xMax && y > 0 {
					symbolZoneMap[Point{x + 1, y - 1}] = true
				}
				if x > 0 {
					symbolZoneMap[Point{x - 1, y}] = true
				}
				if x < xMax {
					symbolZoneMap[Point{x + 1, y}] = true
				}
				if x > 0 && y < yMax {
					symbolZoneMap[Point{x - 1, y + 1}] = true
				}
				if y < yMax {
					symbolZoneMap[Point{x, y + 1}] = true
				}
				if x < xMax && y < yMax {
					symbolZoneMap[Point{x + 1, y + 1}] = true
				}
			}
		}
	}

	sum := 0

	for y, row := range grid {
		currentNumberString := ""
		isPart := false
		for x, tile := range row {
			if isNumber(tile) {
				currentNumberString = currentNumberString + tile
				if symbolZoneMap[Point{x, y}] {
					isPart = true
				}
			} else {
				//End of number
				if currentNumberString != "" {
					number, err := strconv.Atoi(currentNumberString)
					if err != nil {
						log.Fatal(err)
					}
					if isPart {
						sum = sum + number
					}
					currentNumberString = ""
					isPart = false
				}
			}
			if x == xMax {
				//End of row
				if currentNumberString != "" {
					number, err := strconv.Atoi(currentNumberString)
					if err != nil {
						log.Fatal(err)
					}
					if isPart {
						sum = sum + number
					}
					currentNumberString = ""
					isPart = false
				}
			}
		}
	}

	fmt.Println(sum)
}

func isSymbol(input string) bool {
	if input == "." {
		return false
	}
	runes := []rune(input)
	return !unicode.IsDigit(runes[0])
}

func isNumber(input string) bool {
	runes := []rune(input)
	return unicode.IsDigit(runes[0])
}
