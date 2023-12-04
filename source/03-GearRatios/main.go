package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Point struct {
	X int
	Y int
}

func main() {
	part1()
	part2()
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
				// End of number
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
				// End of row
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

type PartNumber struct {
	Value int
	Start Point
	End   Point
}

func (pn PartNumber) isNeighbour(point Point) bool {
	if (point.X >= pn.Start.X-1) && (point.X <= pn.End.X+1) {
		if (point.Y >= pn.Start.Y-1) && (point.Y <= pn.Start.Y+1) { // Start.Y == End.Y
			return true
		}
	}
	return false
}

func part2() {
	grid := helper.ReadGrid("input.txt")
	xMax := len(grid[0]) - 1

	gearMap := make(map[Point]bool)
	for y, row := range grid {
		for x, tile := range row {
			if isGear(tile) {
				gearMap[Point{x, y}] = true
			}
		}
	}

	var partNumbers []PartNumber
	for y, row := range grid {
		start := Point{}
		end := Point{}
		currentNumberString := ""
		for x, tile := range row {
			if isNumber(tile) {
				// Start of new number
				if currentNumberString == "" {
					start.X = x
					start.Y = y
				}
				currentNumberString = currentNumberString + tile
			} else {
				// End of number
				if currentNumberString != "" {
					number, err := strconv.Atoi(currentNumberString)
					if err != nil {
						log.Fatal(err)
					}
					end.X = x - 1
					end.Y = y
					pn := PartNumber{
						Value: number,
						Start: start,
						End:   end,
					}
					partNumbers = append(partNumbers, pn)
					currentNumberString = ""
				}
			}
			if x == xMax {
				// End of row
				if currentNumberString != "" {
					number, err := strconv.Atoi(currentNumberString)
					if err != nil {
						log.Fatal(err)
					}
					end.X = x
					end.Y = y
					pn := PartNumber{
						Value: number,
						Start: start,
						End:   end,
					}
					partNumbers = append(partNumbers, pn)
					currentNumberString = ""
				}
			}
		}
	}

	sum := 0

	for gear := range gearMap {
		nbOfNeighbours := 0
		var neighbours []PartNumber
		for _, pn := range partNumbers {
			if pn.isNeighbour(gear) {
				nbOfNeighbours++
				neighbours = append(neighbours, pn)
			}
		}
		if nbOfNeighbours == 2 {
			gearRatio := neighbours[0].Value * neighbours[1].Value
			sum = sum + gearRatio
		}
	}

	fmt.Println(sum)
}

func isGear(input string) bool {
	return input == "*"
}
