package main

import "testing"

func TestPart1(t *testing.T) {
	data := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expected := 198

	actual := Part1(data)
	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expected := 230

	actual := Part2(data)
	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}
