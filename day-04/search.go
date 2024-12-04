package day04

const PATTERN string = "XMAS"
const REVERSE_PATTERN string = "SAMX"

func checkMas(m, a, s rune) bool {
	return m == 'M' && a == 'A' && s == 'S'
}

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

func searchVertical(grid [][]rune, height, i, j int) int {
	var count int = 0
	if i > 2 && checkMas(grid[i-1][j], grid[i-2][j], grid[i-3][j]) {
		count++
	}
	if i < height-3 && checkMas(grid[i+1][j], grid[i+2][j], grid[i+3][j]) {
		count++
	}
	return count
}

func searchDiagonals(grid [][]rune, width, height, i, j int) int {
	var count int = 0
	if i > 2 && j > 2 && checkMas(grid[i-1][j-1], grid[i-2][j-2], grid[i-3][j-3]) {
		count++
	}
	if i < height-3 && j < width-3 && checkMas(grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3]) {
		count++
	}
	if i > 2 && j < width-3 && checkMas(grid[i-1][j+1], grid[i-2][j+2], grid[i-3][j+3]) {
		count++
	}
	if i < height-3 && j > 2 && checkMas(grid[i+1][j-1], grid[i+2][j-2], grid[i+3][j-3]) {
		count++
	}
	return count
}
