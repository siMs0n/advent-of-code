package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalScore := 0
	for fileScanner.Scan() {
		letters := make(map[rune]bool)
		line := fileScanner.Text()
		sizeOfRugSack := len(line)
		for i, ch := range line {
			if i < sizeOfRugSack/2 {
				letters[ch] = true
			} else if letters[ch] {
				totalScore += findPrio(ch)
				break
			}
		}
	}

	fmt.Println("Total score:", totalScore)
}

func findPrio(ch rune) int {
	charCodeLowerCaseA := []rune("a")[0]
	charCodeUpperCaseA := []rune("A")[0]
	lowerCaseStartingRangePrio := 1
	upperCaseStartingRangePrio := 27

	if ch < charCodeLowerCaseA {
		return int(ch-charCodeUpperCaseA) + upperCaseStartingRangePrio
	} else {
		return int(ch-charCodeLowerCaseA) + lowerCaseStartingRangePrio
	}
}
