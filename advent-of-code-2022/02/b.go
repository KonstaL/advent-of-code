package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
)

type GameRound struct {
	a string
	b string
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}

	strFile := string(data)
	lines := strings.Split(strFile, "\n")

	gameRounds := make([]GameRound, len(lines))

	for i := 0; i < len(lines); i++ {
		roundArgs := strings.Split(lines[i], " ")
		gameRounds[i] = translateArgs(roundArgs[0], roundArgs[1])
	}

	totalScore := 0

	for i := 0; i < len(gameRounds); i++ {
		totalScore += solveGameRound(gameRounds[i])
	}

	fmt.Println(totalScore)

}

func solveGameRound(gameRound GameRound) int {
	score := 0

	if gameRound.a == Rock && gameRound.b == Paper {
		score += 6
	}

	if gameRound.a == Paper && gameRound.b == Scissors {
		score += 6
	}

	if gameRound.a == Scissors && gameRound.b == Rock {
		score += 6
	}

	if gameRound.a == gameRound.b {
		score += 3
	}

	switch gameRound.b {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}

	return score
}

func translateArgs(aStr string, bStr string) GameRound {
	a := Rock
	if aStr == "B" {
		a = Paper
	}
	if aStr == "C" {
		a = Scissors
	}

	if bStr == "X" && a == Rock {
		return GameRound{a, Scissors}
	}
	if bStr == "X" && a == Paper {
		return GameRound{a, Rock}
	}
	if bStr == "X" && a == Scissors {
		return GameRound{a, Paper}
	}

	if bStr == "Y" {
		return GameRound{a, a}
	}

	if bStr == "Z" && a == Rock {
		return GameRound{a, Paper}
	}
	if bStr == "Z" && a == Paper {
		return GameRound{a, Scissors}
	}
	if bStr == "Z" && a == Scissors {
		return GameRound{a, Rock}
	}

	// This should never happen
	return GameRound{a, Rock}
}
