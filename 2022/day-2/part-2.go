package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	movePoints := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	opponentMoves := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}
	desiredOutcomes := map[string]string{
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}

	totalScore := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Fields(line)
		opponentMove, desiredOutcome := opponentMoves[parts[0]], desiredOutcomes[parts[1]]
		myMove := getMyMove(opponentMove, desiredOutcome)
		totalScore += movePoints[myMove] + roundResult(opponentMove, myMove)
	}

	fmt.Println("Total score:", totalScore)
}

func getMyMove(opponentMove string, desiredOutcome string) string {

	moves := []string{"Rock", "Paper", "Scissors"}
	if desiredOutcome == "Win" {
		winIndex := (indexOf(opponentMove, moves) + 1) % 3
		return moves[winIndex]
	} else if desiredOutcome == "Lose" {
		loseIndex := (indexOf(opponentMove, moves) + 2) % 3
		return moves[loseIndex]
	} else {
		return opponentMove
	}
}

func roundResult(opponentMove string, myMove string) int {
	winningMoves := map[string]string{
		"Rock":     "Paper",
		"Paper":    "Scissors",
		"Scissors": "Rock",
	}
	if winningMoves[opponentMove] == myMove {
		return 6
	} else if opponentMove == myMove {
		return 3
	} else {
		return 0
	}
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
