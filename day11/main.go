package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

type octopus struct {
	id, level  int
	neighbours []int
}

func Part1(input []string) int {
	count := 0
	octopuses := parseOctopuses(input)
	for i := 0; i < 100; i++ {
		count += execRound(octopuses)
	}
	return count
}

func Part2(input []string) int {
	count := 0
	octopuses := parseOctopuses(input)
	i := 0
	for count != 100 {
		count = execRound(octopuses)
		i++
	}

	return i
}

func parseOctopuses(input []string) map[int]octopus {
	var octopuses = make(map[int]octopus)
	id := 0
	for y, line := range input {
		for _, level := range strings.Split(line, "") {
			levelNum, _ := strconv.Atoi(level)
			neighbours := getNeighbours(id, y, len(line), len(input)*len(line))
			octopuses[id] = octopus{id, levelNum, neighbours}
			id++
		}
	}
	return octopuses
}

func getNeighbours(id int, y int, maxLen int, max int) []int {
	var neighbours []int
	for i := id - 1; i <= id+1; i++ {
		minRowVal := y * maxLen
		maxRowVal := (y * maxLen) + (maxLen - 1)
		if i < minRowVal || i > maxRowVal {
			continue
		}

		if i != id {
			neighbours = addIfInGrid(neighbours, i, max)
		}
		neighbours = addIfInGrid(neighbours, i-maxLen, max)
		neighbours = addIfInGrid(neighbours, i+maxLen, max)
	}
	return neighbours
}

func addIfInGrid(list []int, val int, max int) []int {
	if val >= 0 && val < max {
		return append(list, val)
	}
	return list
}

func execRound(octopuses map[int]octopus) int {
	var newFlashers = make(map[int]bool)
	for _, oct := range octopuses {
		oct.level++
		octopuses[oct.id] = oct
		if oct.level > 9 {
			newFlashers[oct.id] = true
		}
	}

	var flashers = make(map[int]bool)
	doFlash(octopuses, newFlashers, flashers)
	for len(newFlashers) > 0 {
		newFlashers = make(map[int]bool)
		for _, oct := range octopuses {
			if oct.level > 9 {
				newFlashers[oct.id] = true
			}
		}
		doFlash(octopuses, newFlashers, flashers)
	}

	return len(flashers)
}

func doFlash(octopuses map[int]octopus, flashers map[int]bool, hasFlashed map[int]bool) map[int]bool {
	var newFlashers = make(map[int]bool)

	for id := range flashers {
		oct := octopuses[id]
		for _, n := range oct.neighbours {
			ne := octopuses[n]
			if _, isFlashed := hasFlashed[n]; !isFlashed {
				ne.level++
				if ne.level > 9 {
					newFlashers[n] = true
				}
				octopuses[n] = ne
			}
		}
		oct.level = 0
		octopuses[id] = oct
		hasFlashed[id] = true
	}

	return newFlashers
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
