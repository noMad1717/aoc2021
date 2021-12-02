package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

func ParseCommand(command string) (string, int) {
	commandParts := strings.Split(command, " ")
	action := commandParts[0]
	val, _ := strconv.Atoi(commandParts[1])
	return action, val
}

func AdjustPosition(command string, horizontal int, depth int) (int, int) {
	action, val := ParseCommand(command)

	switch action {
	case "forward":
		horizontal += val
		break
	case "down":
		depth += val
		break
	case "up":
		depth -= val
		break
	}

	return horizontal, depth
}

func AdjustPosWithAim(command string, horizontal int, depth int, aim int) (int, int, int) {
	action, val := ParseCommand(command)

	switch action {
	case "forward":
		horizontal += val
		depth += (val * aim)
		break
	case "down":
		aim += val
		break
	case "up":
		aim -= val
		break
	}

	return horizontal, depth, aim
}

func Part1(input []string) int {
	horizontal, depth := 0, 0
	for _, command := range input {
		horizontal, depth = AdjustPosition(command, horizontal, depth)
	}
	return horizontal * depth
}

func Part2(input []string) int {
	horizontal, depth, aim := 0, 0, 0
	for _, command := range input {
		horizontal, depth, aim = AdjustPosWithAim(command, horizontal, depth, aim)
	}
	return horizontal * depth
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
