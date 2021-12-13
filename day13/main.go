package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

type point struct {
	x, y int
}

type fold struct {
	axis  string
	value int
}

func Part1(input []string) int {
	points, instructions := parseInput(input)
	doFold(points, instructions[0])
	return len(points)
}

func Part2(input []string) {
	points, instructions := parseInput(input)
	for _, inst := range instructions {
		doFold(points, inst)
	}

	maxX, maxY := 0, 0
	for _, pt := range points {
		if pt.x > maxX {
			maxX = pt.x
		}
		if pt.y > maxY {
			maxY = pt.y
		}
	}

	for y := 0; y <= maxY; y++ {
		row := ""
		for x := 0; x <= maxX; x++ {
			if _, exists := points[createKey(x, y)]; exists {
				row += "#"
			} else {
				row += " "
			}
		}
		fmt.Println(row)
	}
}

func parseInput(input []string) (map[string]point, []fold) {
	var points = make(map[string]point)
	breakIndex := 0
	for i, line := range input {
		if line == "" {
			breakIndex = i + 1
			break
		}
		parsePoints(points, line)
	}
	folds := parseInstructions(input[breakIndex:])

	return points, folds
}

func parsePoints(points map[string]point, input string) {
	parts := strings.Split(input, ",")
	xVal, _ := strconv.Atoi(parts[0])
	yVal, _ := strconv.Atoi(parts[1])
	points[input] = point{xVal, yVal}
}

func parseInstructions(input []string) []fold {
	var instructions []fold
	for _, line := range input {
		parts := strings.Split(line, " ")
		inst := parts[len(parts)-1]
		instElems := strings.Split(inst, "=")
		pos, _ := strconv.Atoi(instElems[1])
		instructions = append(instructions, fold{instElems[0], pos})
	}

	return instructions
}

func doFold(points map[string]point, instruction fold) {
	if instruction.axis == "y" {
		for key, pt := range points {
			if pt.y > instruction.value {
				newY := changeVal(pt.y, instruction.value)
				points[createKey(pt.x, newY)] = point{pt.x, newY}
				delete(points, key)
			}
		}
	} else {
		for key, pt := range points {
			if pt.x > instruction.value {
				newX := changeVal(pt.x, instruction.value)
				points[createKey(newX, pt.y)] = point{newX, pt.y}
				delete(points, key)
			}
		}
	}
}

func changeVal(inputVal int, foldVal int) int {
	rest := inputVal % foldVal
	if rest != 0 {
		return foldVal - rest
	}
	return 0
}

func createKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(input)

	fmt.Println("Answer for part 1: ", part1)

	fmt.Println("Answer for part 2:")
	Part2(input)
}
