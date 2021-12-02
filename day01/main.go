package main

import (
	"fmt"

	"github.com/noMad1717/aoc2021/utils"
)

func Part1(input []int) int {
	count := 0
	for i, val := range input {
		if i > 0 && val > input[i-1] {
			count++
		}
	}
	return count
}

func Part2(input []int) int {
	count, indexA, indexB := 0, 0, 3
	for indexB < len(input) {
		if input[indexB] > input[indexA] {
			count++
		}
		indexA++
		indexB++
	}
	return count
}

func main() {
	input := utils.FileToIntList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
