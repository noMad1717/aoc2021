package main

import "testing"

func TestPart1(t *testing.T) {
	expectedSmall := 10
	expectedMedium := 19
	expectedLarge := 226

	actualSmall := Part1(dataSmall)
	actualMedium := Part1(dataMedium)
	actualLarge := Part1(dataLarge)

	if expectedSmall != actualSmall {
		t.Errorf("Test failed! Expected %d but got %d", expectedSmall, actualSmall)
	}

	if expectedMedium != actualMedium {
		t.Errorf("Test failed! Expected %d but got %d", expectedMedium, actualMedium)
	}

	if expectedLarge != actualLarge {
		t.Errorf("Test failed! Expected %d but got %d", expectedLarge, actualLarge)
	}
}

func TestPart2(t *testing.T) {
	expectedSmall := 36
	expectedMedium := 103
	expectedLarge := 3509

	actualSmall := Part2(dataSmall)
	actualMedium := Part2(dataMedium)
	actualLarge := Part2(dataLarge)

	if expectedSmall != actualSmall {
		t.Errorf("Test failed! Expected %d but got %d", expectedSmall, actualSmall)
	}

	if expectedMedium != actualMedium {
		t.Errorf("Test failed! Expected %d but got %d", expectedMedium, actualMedium)
	}

	if expectedLarge != actualLarge {
		t.Errorf("Test failed! Expected %d but got %d", expectedLarge, actualLarge)
	}
}

var dataSmall = []string{
	"start-A",
	"start-b",
	"A-c",
	"A-b",
	"b-d",
	"A-end",
	"b-end",
}

var dataMedium = []string{
	"dc-end",
	"HN-start",
	"start-kj",
	"dc-start",
	"dc-HN",
	"LN-dc",
	"HN-end",
	"kj-sa",
	"kj-HN",
	"kj-dc",
}

var dataLarge = []string{
	"fs-end",
	"he-DX",
	"fs-he",
	"start-DX",
	"pj-DX",
	"end-zg",
	"zg-sl",
	"zg-pj",
	"pj-he",
	"RW-he",
	"fs-DX",
	"pj-RW",
	"zg-RW",
	"start-pj",
	"he-WI",
	"zg-he",
	"pj-fs",
	"start-RW",
}
