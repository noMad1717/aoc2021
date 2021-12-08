package utils

import "strings"

func ToStringSet(input string) map[string]bool {
	var set = make(map[string]bool)
	for _, char := range strings.Split(input, "") {
		set[char] = true
	}
	return set
}

func Difference(a map[string]bool, b map[string]bool) map[string]bool {
	var diff = make(map[string]bool)
	for key, _ := range a {
		if !b[key] {
			diff[key] = true
		}
	}
	return diff
}

func Union(a map[string]bool, b map[string]bool) map[string]bool {
	var union = make(map[string]bool)
	for key, _ := range a {
		union[key] = true
	}
	for key, _ := range b {
		union[key] = true
	}
	return union
}

func Intersection(a map[string]bool, b map[string]bool) map[string]bool {
	var intersection = make(map[string]bool)
	if len(b) > len(a) {
		a, b = b, a
	}
	for key, _ := range a {
		if b[key] {
			intersection[key] = true
		}
	}
	return intersection
}
