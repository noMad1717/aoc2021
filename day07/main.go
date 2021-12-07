package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

func Part1(input []string) int {
	positions := parsePositions(input[0])
	result, _ := getShortestTotal(positions, false)
	return result
}

func Part2(input []string) int {
	positions := parsePositions(input[0])
	result, _ := getShortestTotal(positions, true)
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

func getShortestTotal(positions []int, useGauss bool) (int, int) {
	shortestDistance, shortestPos := -1, -1
	for target := positions[0]; target <= positions[len(positions)-1]; target++ {
		dist := 0
		for _, pos := range positions {
			if !useGauss {
				dist += getDistance(target, pos)
			} else {
				n := getDistance(target, pos)
				dist += (n * (n + 1)) / 2
			}

			if shortestDistance > -1 && dist > shortestDistance {
				break
			}
		}
		if dist < shortestDistance || shortestDistance == -1 {
			shortestDistance = dist
			shortestPos = target
		}
	}

	return shortestDistance, shortestPos
}

func getDistance(target int, from int) int {
	if target == from {
		return 0
	}
	if target > from {
		return target - from
	}
	return from - target
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
