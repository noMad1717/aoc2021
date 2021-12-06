package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

func Part1(input []string) int {
	gens := getGenerations(input[0])
	return countNewGens(gens, 80)
}

func Part2(input []string) int {
	gens := getGenerations(input[0])
	return countNewGens(gens, 256)
}

func getGenerations(input string) [9]int {
	var gens = new([9]int)
	for _, val := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(val)
		gens[num]++
	}
	return *gens
}

func countNewGens(gens [9]int, iters int) int {
	for i := 0; i < iters; i++ {
		gens = getNextGen(gens)
	}

	count := 0
	for _, c := range gens {
		count += c
	}
	return count
}

func getNextGen(gens [9]int) [9]int {
	var nextGen = new([9]int)
	for i, count := range gens {
		if i == 0 {
			nextGen[6] = count + gens[7]
			nextGen[8] = count
		} else if i != 7 {
			nextGen[i-1] = gens[i]
		}
	}
	return *nextGen
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
