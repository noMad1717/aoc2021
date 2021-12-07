package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

func Part1(input []string) int {
	positions := parsePositions(input[0])
	result := getShortestTotal(positions, false)
	return result
}

func Part2(input []string) int {
	positions := parsePositions(input[0])
	result := getShortestTotal(positions, true)
	return result
}

func parsePositions(input string) []int {
	var positions []int
	for _, val := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(val)
		positions = append(positions, num)
	}
	sort.Ints(positions)
	return positions
}

func getShortestTotal(positions []int, useGauss bool) int {
	shortestDistance := -1
	for target := positions[0]; target <= positions[len(positions)-1]; target++ {
		dist := 0
		for _, pos := range positions {
			n := int(math.Abs(float64(target - pos)))
			if !useGauss {
				dist += n
			} else {
				dist += (n * (n + 1)) / 2
			}

			if shortestDistance > -1 && dist > shortestDistance {
				break
			}
		}
		if shortestDistance == -1 || dist < shortestDistance {
			shortestDistance = dist
		}
	}

	return shortestDistance
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
