package dayone

import (
	"log"
	"slices"

	s "github.com/dangwn/advent-of-code-2024/shared"
)

func getArrays(inputFile string) ([]int, []int) {
	rawInput, err := s.Read(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var arr1, arr2 []int
	for _, input := range rawInput {
		intArr, err := convArr(input)
		if err != nil {
			log.Fatal(err)
		}
		arr1 = append(arr1, intArr[0])
		arr2 = append(arr2, intArr[1])
	}

	return arr1, arr2
}

func Part1(inputFile string) int {
	arr1, arr2 := getArrays(inputFile)

	slices.Sort(arr1)
	slices.Sort(arr2)

	var sum int = 0
	for i := range arr1 {
		sum += absDiff(arr1[i], arr2[i])
	}

	return sum
}

func Part2(inputFile string) int {
	arr1, arr2 := getArrays(inputFile)

	counter := makeCounter(arr2)
	var sum int = 0

	for _, num := range arr1 {
		if count, ok := counter[num]; ok {
			sum += num * count
		}
	}

	return sum
}
