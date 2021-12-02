package main

import "testing"

func TestPart1(t *testing.T) {
	data := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expected := 150

	actual := Part1(data)
	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expected := 900

	actual := Part2(data)
	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}
