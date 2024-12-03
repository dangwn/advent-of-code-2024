package day03

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func findMatches(text string) []string {
	r, _ := regexp.Compile(instructionPattern)
	return r.FindAllString(text, -1)
}

func findMatchIndices(text, pattern string, matchType MatchType) []Match {
	r, _ := regexp.Compile(pattern)

	var output []Match
	matches := r.FindAllStringIndex(text, -1)
	for _, m := range matches {
		start, end := m[0], m[1]
		output = append(output, Match{
			Index: start,
			Text:  text[start:end],
			Type:  matchType,
		})
	}
	return output
}

func getDigitsFromMatch(text string) (int, int) {
	stripped := text[4 : len(text)-1]
	digits := strings.Split(stripped, ",")

	if len(digits) != 2 {
		log.Fatal("Could not get digits from ", text)
	}

	x, err := strconv.Atoi(digits[0])
	if err != nil {
		log.Fatal("Could not get first digit from", err)
	}
	y, err := strconv.Atoi(digits[1])
	if err != nil {
		log.Fatal("Could not get second digit from", err)
	}

	return x, y
}

func doMult(m string) int {
	x, y := getDigitsFromMatch(m)
	return x * y
}

func gatherCommands(program string) []Match {
	var commands []Match
	commands = append(commands, findMatchIndices(program, instructionPattern, MULT)...)
	commands = append(commands, findMatchIndices(program, doPattern, DO)...)
	commands = append(commands, findMatchIndices(program, dontPattern, DONT)...)

	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Index < commands[j].Index
	})

	return commands
}
