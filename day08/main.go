package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

type Entry struct {
	patterns []string
	output   []string
}

func Part1(input []string) int {
	count := 0
	entries := parseInput(input)

	counts := []int{2, 3, 4, 7}
	for _, entry := range entries {
		count += countDigits(entry.output, counts)
	}
	return count
}

func Part2(input []string) int {
	entries := parseInput(input)
	sum := 0
	for _, entry := range entries {
		key := getKey(entry.patterns)
		str := ""
		for _, output := range entry.output {
			oLen := len(output)
			if oLen == 2 {
				str += "1"
			} else if oLen == 3 {
				str += "7"
			} else if oLen == 4 {
				str += "4"
			} else if oLen == 7 {
				str += "8"
			} else {
				oSet := utils.ToStringSet(output)
				for k, set := range key {
					if oLen == len(set) && len(utils.Difference(set, oSet)) == 0 {
						str += k
						break
					}
				}
			}
		}
		num, _ := strconv.Atoi(str)
		sum += num
	}
	return sum
}

func getKey(patterns []string) map[string]map[string]bool {
	var sets []map[string]bool
	var key = make(map[string]map[string]bool)
	for _, pattern := range patterns {
		set := utils.ToStringSet(pattern)
		switch len(set) {
		case 2:
			key["1"] = set
			break
		case 3:
			key["7"] = set
			break
		case 4:
			key["4"] = set
			break
		case 7:
			key["8"] = set
			break
		default:
			sets = append(sets, set)
			break
		}
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < len(sets); j++ {
			s := sets[j]
			_, eZero := key["0"]
			_, eNine := key["9"]
			_, eThree := key["3"]
			_, eFive := key["5"]
			_, eTwo := key["2"]
			_, eSix := key["6"]
			if len(s) == 6 {
				if !eZero && len(utils.Difference(key["4"], s)) == 1 && len(utils.Difference(key["7"], s)) == 0 {
					key["0"] = s
					continue
				}
				if !eNine && len(utils.Difference(s, utils.Union(key["4"], key["7"]))) == 1 {
					key["9"] = s
					continue
				}
				if !eSix && len(utils.Difference(key["4"], s)) == 1 && len(utils.Difference(key["7"], s)) == 1 {
					key["6"] = s
					continue
				}
			} else {
				if !eThree && len(utils.Difference(key["7"], s)) == 0 {
					key["3"] = s
					continue
				}
				if !eFive && len(utils.Difference(key["4"], s)) == 1 && len(utils.Difference(key["7"], s)) == 1 {
					key["5"] = s
					continue
				}
				if !eTwo && len(utils.Difference(key["4"], s)) == 2 {
					key["2"] = s
					continue
				}
			}
			if j == len(sets)-1 && len(key) < 7 {
				j = 0
			}
		}
		if len(key) == 7 {
			break
		}
	}
	return key
}

func parseInput(input []string) []Entry {
	var entries []Entry
	for _, line := range input {
		parts := strings.Split(line, " | ")
		entries = append(entries, Entry{strings.Split(parts[0], " "), strings.Split(parts[1], " ")})
	}
	return entries
}

func countDigits(vals []string, counts []int) int {
	count := 0
	for _, output := range vals {
		exists, _ := utils.Contains(counts, len(output))
		if exists {
			count++
		}
	}
	return count
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
