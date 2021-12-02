package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7

	actual := Part1(data)
	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 5

	actual := Part2(data)
	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}
