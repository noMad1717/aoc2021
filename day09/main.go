package main

import (
	"fmt"
	"sort"

	"github.com/noMad1717/aoc2021/utils"
)

type Point struct {
	x   int
	y   int
	val int
}

func Part1(input [][]int) int {
	lowPoints := findLowPoints(input)
	sum := 0
	for _, val := range lowPoints {
		sum += val.val + 1
	}
	return sum
}

func Part2(input [][]int) int {
	lowPoints := findLowPoints(input)
	var lengths []int
	for _, point := range lowPoints {
		basin := findBasin(input, point.x, point.y)
		l := 0
		for _, k := range basin {
			l += len(k)
		}
		lengths = append(lengths, l)
	}
	sort.Ints(lengths)
	lengths = lengths[len(lengths)-3:]
	sum := 1
	for _, k := range lengths {
		sum *= k
	}
	return sum
}

func findLowPoints(input [][]int) []Point {
	var lowPoints []Point
	for i, row := range input {
		if i == 0 {
			lowPoints = append(lowPoints, findLowInEdgeRow(row, input[i+1], i)...)
		} else if i == len(input)-1 {
			lowPoints = append(lowPoints, findLowInEdgeRow(row, input[i-1], i)...)
		} else {
			lowPoints = append(lowPoints, findLowInRow(row, input[i-1], input[i+1], i)...)
		}
	}

	return lowPoints
}

func findLowInRow(row []int, top []int, bottom []int, rowIndex int) []Point {
	var low []Point
	for i, val := range row {
		if i == 0 {
			if val < top[i] && val < bottom[i] && val < row[i+1] {
				low = append(low, Point{i, rowIndex, val})
			}
		} else if i == len(row)-1 {
			if val < top[i] && val < bottom[i] && val < row[i-1] {
				low = append(low, Point{i, rowIndex, val})
			}
		} else {
			if val < top[i] && val < bottom[i] && val < row[i-1] && val < row[i+1] {
				low = append(low, Point{i, rowIndex, val})
			}
		}
	}
	return low
}

func findLowInEdgeRow(row []int, compare []int, rowIndex int) []Point {
	var low []Point
	for i, val := range row {
		if i == 0 {
			if val < compare[i] && val < row[i+1] {
				low = append(low, Point{i, rowIndex, val})
			}
		} else if i == len(row)-1 {
			if val < compare[i] && val < row[i-1] {
				low = append(low, Point{i, rowIndex, val})
			}
		} else {
			if val < compare[i] && val < row[i-1] && val < row[i+1] {
				low = append(low, Point{i, rowIndex, val})
			}
		}
	}
	return low
}

func findBasin(input [][]int, startX int, startY int) map[int]map[int]bool {
	var processed = make(map[int]map[int]bool)
	scanX(input, startY, startX, processed)
	return processed
}

func scanX(input [][]int, y int, startX int, processed map[int]map[int]bool) {
	for i := startX; i >= 0; i-- {
		val := input[y][i]
		if val == 9 {
			break
		}
		processX(input, y, i, processed)
	}

	for i := startX; i <= len(input[0])-1; i++ {
		val := input[y][i]
		if val == 9 {
			break
		}
		processX(input, y, i, processed)
	}
}

func processX(input [][]int, y int, x int, processed map[int]map[int]bool) {
	if !isProcessed(processed, y, x) {
		addSetIfNotExists(processed, y)
		processed[y][x] = true
		scanY(input, y, x, processed)
	}
}

func scanY(input [][]int, startY int, x int, processed map[int]map[int]bool) {
	for i := startY; i >= 0; i-- {
		val := input[i][x]
		if val == 9 {
			break
		}
		processY(input, i, x, processed)
	}

	for i := startY; i <= len(input)-1; i++ {
		val := input[i][x]
		if val == 9 {
			break
		}
		processY(input, i, x, processed)
	}
}

func processY(input [][]int, y int, x int, processed map[int]map[int]bool) {
	if !isProcessed(processed, y, x) {
		addSetIfNotExists(processed, y)
		processed[y][x] = true
		scanX(input, y, x, processed)
	}
}

func addSetIfNotExists(processed map[int]map[int]bool, y int) {
	_, exists := processed[y]
	if !exists {
		processed[y] = make(map[int]bool)
	}
}

func isProcessed(processed map[int]map[int]bool, y int, x int) bool {
	xIndices, existsY := processed[y]
	if !existsY {
		return false
	}
	_, existsX := xIndices[x]
	if !existsX {
		return false
	}
	return true
}

func main() {
	input := utils.FileToIntMatrix("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
