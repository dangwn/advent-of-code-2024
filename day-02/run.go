package day02

import (
	"log"

	s "github.com/dangwn/advent-of-code-2024/shared"
)

func isLevelSafe(level1, level2 int) bool {
	return level1 != level2 && s.Abs(level1-level2) <= 3
}

func isSafe(report []int) bool {
	for ind, level := range report[1:] {
		if !isLevelSafe(report[ind], level) {
			return false
		}
	}
	return true
}

func isDecreasing(report []int) bool {
	for ind, val := range report[1:] {
		if report[ind] <= val {
			return false
		}
	}
	return true
}

func isIncreasing(report []int) bool {
	for ind, val := range report[1:] {
		if report[ind] >= val {
			return false
		}
	}
	return true
}

func isReportSafe(report []int) bool {
	return isSafe(report) && (isDecreasing(report) || isIncreasing(report))
}

func isReportMostlySafe(report []int) bool {
	n := len(report)
	for ind := range report {
		newRepo := make([]int, n)
		copy(newRepo, report)
		if isReportSafe(append(newRepo[:ind], newRepo[ind+1:]...)) {
			return true
		}
	}
	return false
}

func Part1(inputFile string) int {
	reports, err := s.GetIntArrays(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var safe int = 0
	for _, report := range reports {
		if isReportSafe(report) {
			safe++
		}
	}

	return safe
}

func Part2(inputFile string) int {
	reports, err := s.GetIntArrays(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var safe int = 0
	for _, report := range reports {
		if isReportSafe(report) || isReportMostlySafe(report) {
			safe++
		}
	}

	return safe
}
