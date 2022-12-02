package main

import (
	"fmt"
	"strings"
)

var (
	opponentRock     string
	opponentPaper    string
	opponentScissors string
	outcomeLose      string
	outcomeDraw      string
	outcomeWin       string
	valueRock        int
	valuePaper       int
	valueScissors    int
	lose             int
	draw             int
	win              int
	myScore          int
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	_ = lines
	initVars()
	for _, line := range lines {
		inputs := strings.Split(line, " ")
		myScore += RockPaperScissors(inputs[0], inputs[1])
	}
	fmt.Println(myScore)
}

func initVars() {
	opponentRock = "A"
	opponentPaper = "B"
	opponentScissors = "C"
	outcomeLose = "X"
	outcomeDraw = "Y"
	outcomeWin = "Z"
	valueRock, valuePaper, valueScissors = 1, 2, 3
	lose, draw, win = 0, 3, 6
	myScore = 0
}

func RockPaperScissors(opponent, desiredOutcome string) int {
	score := 0
	switch opponent {
	case opponentRock:
		switch desiredOutcome {
		case outcomeLose:
			score += lose
			score += valueScissors
		case outcomeDraw:
			score += draw
			score += valueRock
		case outcomeWin:
			score += win
			score += valuePaper
		}
	case opponentPaper:
		switch desiredOutcome {
		case outcomeLose:
			score += lose
			score += valueRock
		case outcomeDraw:
			score += draw
			score += valuePaper
		case outcomeWin:
			score += win
			score += valueScissors
		}
	case opponentScissors:
		switch desiredOutcome {
		case outcomeLose:
			score += lose
			score += valuePaper
		case outcomeDraw:
			score += draw
			score += valueScissors
		case outcomeWin:
			score += win
			score += valueRock
		}
	}
	return score
}
