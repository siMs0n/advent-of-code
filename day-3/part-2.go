package main

import (
	"bufio"
	"errors"
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
	lineCount := 0
	groupRugsacks := make([]string, 3)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineCount += 1
		groupMemberNr := lineCount % 3
		groupRugsacks[groupMemberNr] = line
		if groupMemberNr == 0 {
			commonItem, err := findCommonItem(groupRugsacks)
			if err != nil {
				panic(err)
			}
			totalScore += findPrio(commonItem)
		}
	}

	fmt.Println("Total score:", totalScore)
}

func findCommonItem(rugsacks []string) (rune, error) {
	for _, ch := range rugsacks[0] {
		if indexOf(ch, rugsacks[1]) >= 0 && indexOf(ch, rugsacks[2]) >= 0 {
			return ch, nil
		}
	}
	return '0', errors.New("Found no common item")
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

func indexOf(element rune, data string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
