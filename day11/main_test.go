package main

import "testing"

func TestPart1(t *testing.T) {
	data := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}

	expected := 1656
	actual := Part1(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}

	expected := 195
	actual := Part2(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}
