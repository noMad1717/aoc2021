package main

import "testing"

func TestPart1(t *testing.T) {
	data := []string{"3,4,3,1,2"}

	expected := 5934
	actual := Part1(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := []string{"3,4,3,1,2"}

	expected := 26984457539
	actual := Part2(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}
