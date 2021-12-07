package main

import "testing"

func TestPart1(t *testing.T) {
	data := []string{"16,1,2,0,4,2,7,1,2,14"}

	expected := 37
	actual := Part1(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	data := []string{"16,1,2,0,4,2,7,1,2,14"}

	expected := 168
	actual := Part2(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}
