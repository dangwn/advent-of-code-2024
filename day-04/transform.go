package day04

func stringArrToRuneArr(arr []string) [][]rune {
	var result [][]rune
	for _, str := range arr {
		runes := []rune(str)
		result = append(result, runes)
	}
	return result
}

func rotateArray(arr [][]rune) [][]rune {
	if len(arr) == 0 {
		return nil
	}

	rows := len(arr)
	cols := len(arr[0])

	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-i-1] = arr[i][j]
		}
	}

	return rotated
}

func getRotatedIndex(i, j, width int) (int, int) {
	iNew := width - 1 - j
	jNew := i
	return iNew, jNew
}
