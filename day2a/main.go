package main

import (
	"fmt"
	"strings"
)

var (
	opponentRock     string
	opponentPaper    string
	opponentScissors string
	answerRock       string
	answerPaper      string
	answerScissors   string
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
	answerRock = "X"
	answerPaper = "Y"
	answerScissors = "Z"
	valueRock, valuePaper, valueScissors = 1, 2, 3
	lose, draw, win = 0, 3, 6
	myScore = 0
}

// Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.
func RockPaperScissors(opponent, answer string) int {
	score := 0
	switch opponent {
	case opponentRock:
		switch answer {
		case answerRock:
			score += draw
			score += valueRock
		case answerPaper:
			score += win
			score += valuePaper
		case answerScissors:
			score += lose
			score += valueScissors
		}
	case opponentPaper:
		switch answer {
		case answerRock:
			score += lose
			score += valueRock
		case answerPaper:
			score += draw
			score += valuePaper
		case answerScissors:
			score += win
			score += valueScissors
		}
	case opponentScissors:
		switch answer {
		case answerRock:
			score += win
			score += valueRock
		case answerPaper:
			score += lose
			score += valuePaper
		case answerScissors:
			score += draw
			score += valueScissors
		}
	}
	return score
}
