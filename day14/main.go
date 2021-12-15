package main

import (
	"fmt"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

func Part1(input []string) int {
	return execute(input, 10)
}

func Part2(input []string) int {
	template := input[0]
	rules := parseRules(input[2:])
	return doStuff(template, rules)
}

func execute(input []string, iterations int) int {
	var occurrences = make(map[string]int)
	var pairs []string
	for i := 1; i < len(input[0]); i++ {
		pairs = append(pairs, fmt.Sprintf("%c%c", input[0][i-1], input[0][i]))
		incCount(string(input[0][i-1]), 1, occurrences)
		incCount(string(input[0][i]), 1, occurrences)
	}
	rules := parseRules(input[2:])
	execStep(pairs, rules, occurrences, iterations)

	min, max := findMinMax(occurrences)
	return max - min
}

func parseRules(input []string) map[string]string {
	var rules = make(map[string]string)
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}
	return rules
}

func incCount(char string, count int, occurrences map[string]int) {
	if _, exists := occurrences[char]; exists {
		occurrences[char] += count
	} else {
		occurrences[char] = count
	}
}

func execStep(pairs []string, rules map[string]string, occurrences map[string]int, i int) {
	for j := 0; j < i; j++ {
		var next []string
		for _, pair := range pairs {
			next = append(next, getPairs(pair, rules, occurrences)...)
		}
		pairs = next
	}
}

func getPairs(pair string, rules map[string]string, occurrences map[string]int) []string {
	ins := rules[pair]
	first := string(pair[0]) + ins
	second := ins + string(pair[1])
	incCount(ins, 1, occurrences)
	return []string{first, second}
}

func findMinMax(occurrences map[string]int) (int, int) {
	min, max := 0, 0
	for _, val := range occurrences {
		if min == 0 || val < min {
			min = val
		}
		if max == 0 || val > max {
			max = val
		}
	}
	return min, max
}

func doStuff(template string, rules map[string]string) int {
	var occurrences = make(map[string]int)
	var pairs = make(map[string]int)

	incCount(string(template[0]), 1, occurrences)
	for i := 1; i < len(template); i++ {
		pair := fmt.Sprintf("%c%c", template[i-1], template[i])
		incPairCount(pair, 1, pairs)
		incCount(string(template[i]), 1, occurrences)
	}

	for i := 0; i < 40; i++ {
		pairs = countPairs(pairs, rules, occurrences)
	}

	min, max := findMinMax(occurrences)
	return max - min
}

func countPairs(pairs map[string]int, rules map[string]string, occurrences map[string]int) map[string]int {
	var newPairs = make(map[string]int)
	for k, v := range pairs {
		newPairs[k] = v
	}

	for pair, count := range pairs {
		ins := rules[pair]
		incCount(ins, count, occurrences)
		incPairCount(fmt.Sprintf("%c%s", pair[0], ins), count, newPairs)
		incPairCount(fmt.Sprintf("%s%c", ins, pair[1]), count, newPairs)
		newPairs[pair] -= count
	}

	return newPairs
}

func incPairCount(pair string, count int, pairs map[string]int) {
	if _, exists := pairs[pair]; exists {
		pairs[pair] += count
	} else {
		pairs[pair] = count
	}
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
