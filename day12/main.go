package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

type cave struct {
	name        string
	isSmall     bool
	connections []string
}

func Part1(input []string) int {
	caves := parseCaves(input)
	paths := findPaths(caves, 1)
	return len(paths)
}

func Part2(input []string) int {
	caves := parseCaves(input)
	paths := findPaths(caves, 2)
	return len(paths)
}

func parseCaves(input []string) map[string]cave {
	var caves = make(map[string]cave)
	for _, line := range input {
		pairs := strings.Split(line, "-")
		parseCave(caves, pairs[0], pairs[1])
		parseCave(caves, pairs[1], pairs[0])
	}

	return caves
}

func parseCave(caves map[string]cave, name string, pairName string) {
	c, exists := caves[name]
	if !exists {
		c = cave{name, isSmallCave(name), []string{pairName}}
	} else {
		c.connections = append(c.connections, pairName)
	}
	caves[name] = c
}

func isSmallCave(name string) bool {
	matched, _ := regexp.MatchString(`^[a-z]*$`, name)
	return matched
}

func findPaths(caves map[string]cave, maxVisits int) [][]string {
	start := caves["start"]
	var paths [][]string
	for _, cName := range start.connections {
		c := caves[cName]
		visited := map[string]int{start.name: maxVisits}
		if c.isSmall && c.name != "end" {
			visited[c.name] = 1
		}
		paths = append(paths, exploreCave(caves, c, visited, maxVisits)...)
	}

	var fullPaths [][]string
	for _, path := range paths {
		if len(path) > 0 && path[len(path)-1] == "end" {
			path = append([]string{start.name}, path...)
			fullPaths = append(fullPaths, path)
		}
	}

	return fullPaths
}

func exploreCave(caves map[string]cave, current cave, visitedSmalls map[string]int, maxVisits int) [][]string {
	if current.name == "end" {
		return [][]string{{current.name}}
	}

	var paths [][]string
	for _, cName := range current.connections {
		c := caves[cName]
		var foundPaths [][]string
		if visits, exists := visitedSmalls[c.name]; !exists || (!isExhausted(visitedSmalls) && exists && visits < maxVisits) {
			var visited = make(map[string]int)
			for key, v := range visitedSmalls {
				visited[key] = v
			}
			if c.isSmall && c.name != "end" {
				if _, e := visited[c.name]; e {
					visited[c.name] = 2
				} else {
					visited[c.name] = 1
				}
			}
			foundPaths = exploreCave(caves, c, visited, maxVisits)
			for _, fp := range foundPaths {
				fp = append([]string{current.name}, fp...)
				paths = append(paths, fp)
			}
		}
	}

	return paths
}

func isExhausted(visits map[string]int) bool {
	for k, v := range visits {
		if k != "start" && k != "end" && v == 2 {
			return true
		}
	}
	return false
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
