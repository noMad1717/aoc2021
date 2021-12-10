package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

func Part1(input []string) int {
	total := 0
	for _, line := range input {
		s, _ := getScore(line)
		total += s
	}
	return total
}

func Part2(input []string) int {
	var scores []int
	for _, line := range input {
		s, stack := getScore(line)
		if s == 0 {
			scores = append(scores, getCompletionScore(stack))
		}
	}
	sort.Ints(scores)
	l := len(scores)
	if l%2 == 0 {
		middle := l/2 - 1
		return scores[middle]
	} else {
		middle := l / 2
		return scores[middle]
	}
}

func getScore(input string) (int, []string) {
	closing := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
		">": "<",
	}

	var stack []string
	for _, char := range strings.Split(input, "") {
		if exists(closing, char) && closing[char] != stack[0] {
			switch char {
			case ")":
				return 3, stack
			case "]":
				return 57, stack
			case "}":
				return 1197, stack
			case ">":
				return 25137, stack
			default:
				return 0, stack
			}
		} else if !exists(closing, char) {
			stack = push(stack, char)
		} else {
			stack = pop(stack)
		}
	}
	return 0, stack
}

func getCompletionScore(remainders []string) int {
	closing := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
	}

	points := 0
	for _, open := range remainders {
		points *= 5
		c := closing[open]
		switch c {
		case ")":
			points += 1
			break
		case "]":
			points += 2
			break
		case "}":
			points += 3
			break
		case ">":
			points += 4
			break
		}
	}
	return points
}

func exists(closing map[string]string, elem string) bool {
	_, exists := closing[elem]
	return exists
}

func push(stack []string, elem string) []string {
	pushed := []string{elem}
	return append(pushed, stack...)
}

func pop(stack []string) []string {
	return stack[1:]
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
