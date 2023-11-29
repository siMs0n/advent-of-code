package main

import (
	"fmt"
	"os"
	"strconv"
)

type Range struct {
	start int
	end   int
}

type void struct{}

var member void

func main() {
	bytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	line := string(bytes)

	charactersProcessedAtFirstStart := -1
	letterMap := make(map[string]int)
	for i := 0; i < len(line); i++ {
		letter := string(line[i])
		letterMap[letter] += 1
		if len(letterMap) == 14 {
			charactersProcessedAtFirstStart = i + 1
			fmt.Println("letter", letter)
			break
		}

		if i >= 13 {
			letterToDecrement := string(line[i-13])
			//printHelp(i, letterToDecrement, letterMap)
			if letterMap[letterToDecrement] > 1 {
				letterMap[letterToDecrement] = letterMap[letterToDecrement] - 1
			} else {
				delete(letterMap, letterToDecrement)
			}
		}
	}

	fmt.Println("letterMap", letterMap)
	fmt.Println("charactersProcessedAtFirstStart:", charactersProcessedAtFirstStart)
}

func printHelp(i int, letter string, letterMap map[string]int) {
	if i >= 1978 {
		fmt.Println("i", letter)
		fmt.Println("letterToDecrement", letter)
		fmt.Println("letterMap", letterMap)
	}
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
