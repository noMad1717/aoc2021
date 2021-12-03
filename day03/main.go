package main

import (
	"fmt"
	"strconv"

	"github.com/noMad1717/aoc2021/utils"
)

func mostCommonBit(input []string) string {
	result := ""
	max := len(input[0])
	i := 0
	for i < max {
		zeroCount, oneCount := 0, 0
		for _, binary := range input {
			if string(binary[i]) == "0" {
				zeroCount++
			} else {
				oneCount++
			}
		}
		if zeroCount > oneCount {
			result += "0"
		} else {
			result += "1"
		}
		i++
	}
	return result
}

func invert(input string) string {
	result := ""
	for _, char := range input {
		if string(char) == "0" {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func iterate(input []string, mostCommon bool) string {
	result := input
	max := len(input[0])
	i := 0
	for i < max {
		result = filter(result, i, mostCommon)
		if len(result) == 1 {
			break
		}
		i++
	}
	return result[0]
}

func filter(input []string, pos int, mostCommon bool) []string {
	var zeroes []string
	var ones []string

	for _, val := range input {
		char := string(val[pos])
		if char == "0" {
			zeroes = append(zeroes, val)
		} else {
			ones = append(ones, val)
		}
	}

	if len(zeroes) == len(ones) {
		if mostCommon {
			return ones
		} else {
			return zeroes
		}
	}

	if mostCommon {
		if len(zeroes) > len(ones) {
			return zeroes
		} else {
			return ones
		}
	} else {
		if len(zeroes) > len(ones) {
			return ones
		} else {
			return zeroes
		}
	}
}

func toDecimal(input string) int {
	decimal, _ := strconv.ParseInt(input, 2, 64)
	return int(decimal)
}

func Part1(input []string) int {
	mostCommon := mostCommonBit(input)
	inverted := invert(mostCommon)
	return toDecimal(mostCommon) * toDecimal(inverted)
}

func Part2(input []string) int {
	mostCommon := iterate(input, true)
	leastCommon := iterate(input, false)
	return toDecimal(mostCommon) * toDecimal(leastCommon)
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
