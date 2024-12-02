package shared

import (
	"strconv"
)

func StringArrToInt(arr []string) ([]int, error) {
	var out []int
	for _, val := range arr {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return out, err
		}
		out = append(out, intVal)
	}
	return out, nil
}

func GetIntArrays(inputFile string) ([][]int, error) {
	var out [][]int
	rawInput, err := ReadArrays(inputFile)
	if err != nil {
		return out, err
	}

	for _, input := range rawInput {
		intArr, err := StringArrToInt(input)
		if err != nil {
			return out, err
		}
		out = append(out, intArr)
	}

	return out, nil
}
