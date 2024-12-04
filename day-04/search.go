package day04

import (
	s "github.com/dangwn/advent-of-code-2024/shared"
)

const PATTERN string = "XMAS"
const REVERSE_PATTERN string = "SAMX"

func searchHorizontal(grid [][]rune, width, i, j int) int {
	var count int = 0
	if j > 2 && string(grid[i][j-3:j+1]) == REVERSE_PATTERN {
		count++
	}
	if j < width-3 && string(grid[i][j:j+4]) == PATTERN {
		count++
	}
	return count
}

func searchDiagonal(grid [][]rune, width, height, i, j int) int {
	var count = 0
	if s.AllTrue([]bool{
		i > 2,
		j > 2,
		grid[i-1]
	})
}
