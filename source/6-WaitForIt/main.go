package main

import (
	"fmt"
	"math"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	part1()
}

func part1() {
	var races []Race
	//races = append(races, Race{Time: 7, Distance: 9})
	//races = append(races, Race{Time: 15, Distance: 40})
	//races = append(races, Race{Time: 30, Distance: 200})

	races = append(races, Race{Time: 54, Distance: 239})
	races = append(races, Race{Time: 70, Distance: 1142})
	races = append(races, Race{Time: 82, Distance: 1295})
	races = append(races, Race{Time: 75, Distance: 1253})

	result := 1

	for _, race := range races {
		result = result * evaluateRace(race)
	}

	fmt.Println(result)
}

func evaluateRace(race Race) int {
	a := float64(-1)
	b := float64(race.Time)
	c := float64(-race.Distance)
	discriminant := b*b - 4*a*c
	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	successes := int(math.Ceil(x2)) - int(math.Floor(x1)) - 1
	return successes
}
