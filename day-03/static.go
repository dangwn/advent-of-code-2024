package day03

type MatchType int

type Match struct {
	Index int
	Text  string
	Type  MatchType
}

const (
	MULT MatchType = iota
	DO
	DONT

	instructionPattern string = `mul\([0-9]+,[0-9]+\)`
	doPattern                 = `do\(\)`
	dontPattern               = `don't\(\)`
)
