package main

import "testing"

func TestPart1(t *testing.T) {
	expected := 1588
	actual := Part1(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 2188189693529
	actual := Part2(data)

	if expected != actual {
		t.Errorf("Test failed! Expected %d but got %d", expected, actual)
	}
}

var data = []string{
	"NNCB",
	"",
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C",
}
