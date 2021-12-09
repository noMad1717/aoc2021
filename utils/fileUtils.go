package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func FileToStringList(path string) []string {
	data, _ := ioutil.ReadFile(path)
	stringified := string(data)
	return strings.Split(strings.TrimSpace(stringified), "\n")
}

func FileToIntList(path string) []int {
	stringList := FileToStringList(path)
	var intList []int

	for _, x := range stringList {
		intVal, _ := strconv.Atoi(x)
		intList = append(intList, intVal)
	}

	return intList
}

func FileToIntMatrix(path string) [][]int {
	stringList := FileToStringList(path)
	var matrix [][]int

	for _, x := range stringList {
		var row []int
		for _, char := range strings.Split(x, "") {
			num, _ := strconv.Atoi(char)
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	return matrix
}
