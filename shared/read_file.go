package shared

import (
	"bufio"
	"os"
	"strings"
)

func readLines(filePath string) ([]string, error) {
	var out []string

	file, err := os.Open(filePath)
	if err != nil {
		return out, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, scanner.Err()
}

func Read(filePath string) ([][]string, error) {
	var output [][]string

	lines, err := readLines(filePath)
	if err != nil {
		return output, err
	}

	for _, line := range lines {
		output = append(output, strings.Fields(line))
	}
	return output, nil
}
