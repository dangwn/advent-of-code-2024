package day04

import "fmt"

// import (
// 	"log"

// 	s "github.com/dangwn/advent-of-code-2024/shared"
// )

func Part1(inputFile string) int {
	// grid, err := s.ReadLines(inputFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rawGrid := []string{
		"S", "A", "M", "X", "X", "M", "A", "S",
	}

	grid := stringArrToRuneArr(rawGrid)
	invGrid := rotateArray(grid)

	height := len(grid)
	width := len(grid[0])
	var total int = 0

	for i, row := range grid {
		for j, char := range row {
			fmt.Println(i, j)
			if char != 'X' {
				continue
			}
			total += searchHorizontal(grid, width, i, j)

			invI, invJ := getRotatedIndex(i, j, width)
			total += searchHorizontal(invGrid, height, invI, invJ)
		}
	}

	return total
}
