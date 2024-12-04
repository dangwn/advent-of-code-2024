// This whole day is an absolute shambles and I'm ashamed to put
// my name to it
package day04

import (
	"log"

	s "github.com/dangwn/advent-of-code-2024/shared"
)

func stringArrToRuneArr(arr []string) [][]rune {
	var result [][]rune
	for _, str := range arr {
		runes := []rune(str)
		result = append(result, runes)
	}
	return result
}

func Part1(inputFile string) int {
	rawGrid, err := s.ReadLines(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	grid := stringArrToRuneArr(rawGrid)
	var total int = 0
	width := len(grid[0])
	height := len(grid)

	for i, row := range grid {
		for j, char := range row {
			if char != 'X' {
				continue
			}
			total += searchHorizontal(grid, width, i, j)
			total += searchVertical(grid, height, i, j)
			total += searchDiagonals(grid, width, height, i, j)
		}
	}

	return total
}

func Part2(inputFile string) int {
	rawGrid, err := s.ReadLines(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	grid := stringArrToRuneArr(rawGrid)
	var total int = 0
	width := len(grid[0])
	height := len(grid)

	for i, row := range grid[1 : height-1] {
		for j, char := range row[1 : width-1] {
			if char != 'A' {
				continue
			}
			if ((grid[i][j] == 'M' && grid[i+2][j+2] == 'S') ||
				(grid[i][j] == 'S' && grid[i+2][j+2] == 'M')) &&
				((grid[i][j+2] == 'S' && grid[i+2][j] == 'M') ||
					(grid[i][j+2] == 'M' && grid[i+2][j] == 'S')) {
				total++
			}
		}
	}
	return total
}
