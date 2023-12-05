package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/LenaBullens/advent-of-code-2022-go/source/helper"
)

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(number int) bool {
	return number >= r.Start && number <= r.End
}

type Mapping struct {
	Source      Range
	Destination Range
}

func main() {
	//part1()
	part2()
}

func part1() {
	lines := helper.ReadLines("input.txt")

	// Seeds
	seedsLine := lines[0]
	seedNumbersString := seedsLine[7:]
	splitString1 := strings.Split(seedNumbersString, " ")
	var seedNumbers []int
	for _, seedNumberString := range splitString1 {
		seedNumber := parseInt(seedNumberString)
		seedNumbers = append(seedNumbers, seedNumber)
	}

	// Seeds to soil
	i := 3
	var seedToSoilMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		seedToSoilMappings = append(seedToSoilMappings, mapping)
		i++
	}

	// Soil to fertilizer
	i = i + 2
	var soilToFertilizerMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		soilToFertilizerMappings = append(soilToFertilizerMappings, mapping)
		i++
	}

	// Fertilizer to water
	i = i + 2
	var fertilizerToWaterMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		fertilizerToWaterMappings = append(fertilizerToWaterMappings, mapping)
		i++
	}

	// Water to light
	i = i + 2
	var waterToLightMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		waterToLightMappings = append(waterToLightMappings, mapping)
		i++
	}

	// Light to temperature
	i = i + 2
	var lightToTemperatureMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		lightToTemperatureMappings = append(lightToTemperatureMappings, mapping)
		i++
	}

	// Temperature to humidity
	i = i + 2
	var temperatureToHumidityMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		temperatureToHumidityMappings = append(temperatureToHumidityMappings, mapping)
		i++
	}

	// Humidity to location
	i = i + 2
	var humidityToLocationMappings []Mapping
	for i < len(lines) {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		humidityToLocationMappings = append(humidityToLocationMappings, mapping)
		i++
	}

	lowest := math.MaxInt

	for _, seedNumber := range seedNumbers {
		soilNumber := doMap(seedNumber, seedToSoilMappings)
		fertilizerNumber := doMap(soilNumber, soilToFertilizerMappings)
		waterNumber := doMap(fertilizerNumber, fertilizerToWaterMappings)
		lightNumber := doMap(waterNumber, waterToLightMappings)
		temperatureNumber := doMap(lightNumber, lightToTemperatureMappings)
		humidityNumber := doMap(temperatureNumber, temperatureToHumidityMappings)
		locationNumber := doMap(humidityNumber, humidityToLocationMappings)

		if locationNumber < lowest {
			lowest = locationNumber
		}
	}

	fmt.Println(lowest)
}

func part2() {
	lines := helper.ReadLines("input.txt")

	// Seeds
	seedsLine := lines[0]
	seedNumbersString := seedsLine[7:]
	splitString1 := strings.Split(seedNumbersString, " ")
	var seedRanges []Range
	for i := 0; i < len(splitString1); i = i + 2 {
		seedStart := parseInt(splitString1[i])
		offset := parseInt(splitString1[i+1])
		r := Range{
			Start: seedStart,
			End:   seedStart + offset - 1,
		}
		seedRanges = append(seedRanges, r)
	}

	// Seeds to soil
	i := 3
	var seedToSoilMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		seedToSoilMappings = append(seedToSoilMappings, mapping)
		i++
	}

	// Soil to fertilizer
	i = i + 2
	var soilToFertilizerMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		soilToFertilizerMappings = append(soilToFertilizerMappings, mapping)
		i++
	}

	// Fertilizer to water
	i = i + 2
	var fertilizerToWaterMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		fertilizerToWaterMappings = append(fertilizerToWaterMappings, mapping)
		i++
	}

	// Water to light
	i = i + 2
	var waterToLightMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		waterToLightMappings = append(waterToLightMappings, mapping)
		i++
	}

	// Light to temperature
	i = i + 2
	var lightToTemperatureMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		lightToTemperatureMappings = append(lightToTemperatureMappings, mapping)
		i++
	}

	// Temperature to humidity
	i = i + 2
	var temperatureToHumidityMappings []Mapping
	for lines[i] != "" {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		temperatureToHumidityMappings = append(temperatureToHumidityMappings, mapping)
		i++
	}

	// Humidity to location
	i = i + 2
	var humidityToLocationMappings []Mapping
	for i < len(lines) {
		splitString2 := strings.Split(lines[i], " ")
		destinationStart := parseInt(splitString2[0])
		sourceStart := parseInt(splitString2[1])
		width := parseInt(splitString2[2])

		source := Range{
			Start: sourceStart,
			End:   sourceStart + width - 1,
		}
		destination := Range{
			Start: destinationStart,
			End:   destinationStart + width - 1,
		}
		mapping := Mapping{
			Source:      source,
			Destination: destination,
		}
		humidityToLocationMappings = append(humidityToLocationMappings, mapping)
		i++
	}

	done := false
	location := 0
	for !done {
		humidityNumber := doReverseMap(location, humidityToLocationMappings)
		temperatureNumber := doReverseMap(humidityNumber, temperatureToHumidityMappings)
		lightNumber := doReverseMap(temperatureNumber, lightToTemperatureMappings)
		waterNumber := doReverseMap(lightNumber, waterToLightMappings)
		fertilizerNumber := doReverseMap(waterNumber, fertilizerToWaterMappings)
		soilNumber := doReverseMap(fertilizerNumber, soilToFertilizerMappings)
		seedNumber := doReverseMap(soilNumber, seedToSoilMappings)

		for _, seedRange := range seedRanges {
			if seedRange.Contains(seedNumber) {
				done = true
			}
		}

		if !done {
			location++
		}
	}

	fmt.Println(location)
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func doMap(input int, mappings []Mapping) int {
	done := false
	var result int
	for _, mapping := range mappings {
		if mapping.Source.Contains(input) {
			offset := input - mapping.Source.Start
			result = mapping.Destination.Start + offset
			done = true
			break
		}
	}
	if !done {
		result = input
	}

	return result
}

func doReverseMap(input int, mappings []Mapping) int {
	done := false
	var result int
	for _, mapping := range mappings {
		if mapping.Destination.Contains(input) {
			offset := input - mapping.Destination.Start
			result = mapping.Source.Start + offset
			done = true
			break
		}
	}
	if !done {
		result = input
	}

	return result
}
