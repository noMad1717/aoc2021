package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noMad1717/aoc2021/utils"
)

type Coord struct {
	row int
	col int
}

func Part1(numbers []int, boards [][][]int) int {
	var marked [][]Coord
	for i, num := range numbers {
		for j, board := range boards {
			if i == 0 {
				marked = append(marked, []Coord{})
			}

			found, row, col := utils.JaggedContains(board, num)
			if found {
				marked[j] = append(marked[j], Coord{row, col})
				isBingo := checkBingo(marked[j])

				if isBingo {
					sum := countSum(board, marked[j])
					return sum * num
				}
			}
		}
	}
	return 0
}

func checkBingo(marked []Coord) bool {
	if len(marked) < 5 {
		return false
	}

	for i := 0; i < 5; i++ {
		r, c := 0, 0
		for _, m := range marked {
			if m.col == i {
				c++
			}
			if m.row == i {
				r++
			}
		}

		if r == 5 || c == 5 {
			return true
		}
	}

	return false
}

func countSum(board [][]int, marked []Coord) int {
	sum := 0
	for i, row := range board {
		for j, col := range row {
			isMarked := false
			for _, m := range marked {
				if m.row == i && m.col == j {
					isMarked = true
					break
				}
			}
			if !isMarked {
				sum += col
			}
		}
	}
	return sum
}

func Part2(numbers []int, boards [][][]int) int {
	var marked [][]Coord
	var bingoBoards []int
	lastBingoNum := 0

	for i, num := range numbers {
		for j, board := range boards {
			alreadyBingo, _ := utils.Contains(bingoBoards, j)
			if alreadyBingo {
				continue
			}

			if i == 0 {
				marked = append(marked, []Coord{})
			}

			found, row, col := utils.JaggedContains(board, num)
			if found {
				marked[j] = append(marked[j], Coord{row, col})

				isBingo := checkBingo(marked[j])

				if isBingo {
					bingoBoards = append(bingoBoards, j)
					lastBingoNum = num
				}
			}
		}
	}

	lastBingoBoard := bingoBoards[len(bingoBoards)-1]
	sum := countSum(boards[lastBingoBoard], marked[lastBingoBoard])
	return sum * lastBingoNum
}

func ParseInput(input []string) ([]int, [][][]int) {
	input = append(input, "")
	var numbers []int
	for _, val := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(val)
		numbers = append(numbers, num)
	}

	var boards [][][]int
	var board [][]int
	var row []int
	for _, line := range input[2:] {
		if line == "" {
			boards = append(boards, board)
			board = [][]int{}
			continue
		}

		for _, val := range strings.Split(line, " ") {
			if val == "" {
				continue
			}
			num, _ := strconv.Atoi(val)
			row = append(row, num)
		}
		board = append(board, row)
		row = []int{}
	}

	return numbers, boards
}

func main() {
	input := utils.FileToStringList("data")

	part1 := Part1(ParseInput(input))
	part2 := Part2(ParseInput(input))

	fmt.Println("Answer for part 1: ", part1)
	fmt.Println("Answer for part 2: ", part2)
}
