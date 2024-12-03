package day03

import (
	"log"
	"strings"

	s "github.com/dangwn/advent-of-code-2024/shared"
)

func Part1(filePath string) int {
	programs, err := s.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}
	program := strings.Join(programs, "_")

	matches := findMatches(program)
	multiples := s.Map(matches, doMult)

	return s.Sum(multiples)
}

func Part2(filePath string) int {
	programs, err := s.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	program := strings.Join(programs, "_")
	commands := gatherCommands(program)

	var total int = 0
	var shouldMult bool = true
	for _, command := range commands {
		switch command.Type {
		case MULT:
			if shouldMult {
				total += doMult(command.Text)
			}
		case DO:
			shouldMult = true
		case DONT:
			shouldMult = false
		}
	}

	return total
}
