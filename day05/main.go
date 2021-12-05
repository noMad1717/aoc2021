package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

type Coord struct {
	x int
	y int
}

type Range struct {
	from Coord
	to   Coord
}

func Part1(input []string) int {
	lines := extrapolate(parseCoords(input), false)
	overlaps := countIntersections(lines, 2)
	return overlaps
}

func Part2(input []string) int {
	lines := extrapolate(parseCoords(input), true)
	overlaps := countIntersections(lines, 2)
	return overlaps
}

func parseCoords(input []string) []Range {
	var ranges []Range
	for _, line := range input {
		rangeParts := strings.Split(line, " -> ")
		ranges = append(ranges, Range{getCoord(rangeParts[0]), getCoord(rangeParts[1])})
	}
	return ranges
}

func getCoord(pair string) Coord {
	parts := strings.Split(pair, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	coord := Coord{x, y}
	return coord
}

func countIntersections(lineMap map[string]int, min int) int {
	count := 0
	for _, val := range lineMap {
		if val >= min {
			count++
		}
	}
	return count
}

func extrapolate(ranges []Range, includeDiagonal bool) map[string]int {
	var lineMap = make(map[string]int)
	for _, r := range ranges {
		line := drawLine(r, includeDiagonal)
		for _, seg := range line {
			val, exists := lineMap[seg]
			if exists {
				lineMap[seg] = val + 1
			} else {
				lineMap[seg] = 1
			}
		}
	}
	return lineMap
}

func drawLine(r Range, includeDiagonal bool) []string {
	var lineSegments []string
	from := r.from
	to := r.to

	if from.x == to.x {
		if from.y < to.y {
			lineSegments = appendSegs(lineSegments, from.x, from.y, to.y, false)
		} else {
			lineSegments = appendSegs(lineSegments, from.x, to.y, from.y, false)
		}
	} else if from.y == to.y {
		if from.x < to.x {
			lineSegments = appendSegs(lineSegments, from.y, from.x, to.x, true)
		} else {
			lineSegments = appendSegs(lineSegments, from.y, to.x, from.x, true)
		}
	} else if includeDiagonal {
		lineSegments = appendSegsDiagonal(lineSegments, from, to)
	}
	return lineSegments
}

func appendSegs(list []string, c int, from int, to int, invert bool) []string {
	for i := from; i <= to; i++ {
		if !invert {
			list = append(list, fmt.Sprintf("%d,%d", c, i))
		} else {
			list = append(list, fmt.Sprintf("%d,%d", i, c))
		}
	}
	return list
}

func appendSegsDiagonal(list []string, from Coord, to Coord) []string {
	y := from.y
	if from.x < to.x {
		for x := from.x; x <= to.x; x++ {
			list = append(list, fmt.Sprintf("%d,%d", x, y))
			if from.y < to.y {
				y++
			} else {
				y--
			}
		}
	} else {
		for x := from.x; x >= to.x; x-- {
			list = append(list, fmt.Sprintf("%d,%d", x, y))
			if from.y < to.y {
				y++
			} else {
				y--
			}
		}
	}
	return list
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
