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
		gameRounds[i] = GameRound{translateArg(roundArgs[0]), translateArg(roundArgs[1])}
	}

	totalScore := 0

	for i := 0; i < len(gameRounds); i++ {
		totalScore += solveGameRound(gameRounds[i])
	}

	fmt.Println(totalScore)

}

func solveGameRound(gameRound GameRound) int {
	score := 0

	switch gameRound.b {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}

	score += calcScore(gameRound.a, gameRound.b)

	return score
}

func translateArg(arg string) string {
	if arg == "A" || arg == "X" {
		return Rock
	}
	if arg == "B" || arg == "Y" {
		return Paper
	}
	if arg == "C" || arg == "Z" {
		return Scissors
	}

	return ""
}

func calcScore(a string, b string) int {

	if a == Rock && b == Paper {
		return 6
	}
	if a == Rock && b == Scissors {
		return 0
	}
	if a == Paper && b == Rock {
		return 0
	}
	if a == Paper && b == Scissors {
		return 6
	}
	if a == Scissors && b == Rock {
		return 6
	}
	if a == Scissors && b == Paper {
		return 0
	}

	// if a == b
	return 3
}
