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
	myMoves := map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	totalScore := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Fields(line)
		opponentMove, myMove := opponentMoves[parts[0]], myMoves[parts[1]]
		totalScore += movePoints[myMove] + roundResult(opponentMove, myMove)
	}

	fmt.Println("Total score:", totalScore)
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
