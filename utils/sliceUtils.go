package utils

func Contains(slice []int, search int) (bool, int) {
	for i, val := range slice {
		if val == search {
			return true, i
		}
	}
	return false, -1
}

func JaggedContains(slice [][]int, search int) (bool, int, int) {
	for i, row := range slice {
		found, col := Contains(row, search)
		if found {
			return true, i, col
		}
	}
	return false, -1, -1
}
