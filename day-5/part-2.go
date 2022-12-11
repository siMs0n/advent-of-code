package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Range struct {
	start int
	end   int
}

type Stack []byte

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str byte) {
	*s = append(*s, str)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) MultiPop(nr int) ([]byte, bool) {
	result := make([]byte, 0)
	if s.IsEmpty() {
		return result, false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		indexToStart := index - nr + 1
		for i := indexToStart; i <= index; i++ {
			element := (*s)[i] // Index into the slice and obtain the element.
			result = append(result, element)
		}
		*s = (*s)[:indexToStart] // Remove it from the stack by slicing it off.
		return result, true
	}
}

func (s *Stack) Peek() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element, true
	}
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	// Find where to start parsing, not at end
	heightOfCrates := 8

	stacks := make([]Stack, 0)
	for i := 1; i < len(lines[heightOfCrates])-1; i = i + 4 {
		var stack Stack
		stacks = append(stacks, stack)
	}
	for i := heightOfCrates - 1; i >= 0; i-- {
		stackNr := 0
		for j := 1; j < len(lines[i])-1; j = j + 4 {
			if !unicode.IsSpace(rune(lines[i][j])) {
				stacks[stackNr].Push(lines[i][j])
			}
			stackNr++
		}
	}

	fmt.Println("Stacks at the start")
	for i := 0; i < len(stacks); i++ {
		for j := 0; j < len(stacks[i]); j++ {
			fmt.Print(string(stacks[i][j]))
		}
		fmt.Println("")
	}

	// Loop
	startOfMoves := heightOfCrates + 2
	r, _ := regexp.Compile("\\d+")
	for i := startOfMoves; i < len(lines); i++ {
		//for i := startOfMoves; i < startOfMoves+1; i++ {
		// Parse moves with regexp
		matches := r.FindAllString(lines[i], -1)
		nrToMove := getNumber(matches[0])
		fromStack := getNumber(matches[1])
		toStack := getNumber(matches[2])
		// Perform move by finding correct stack
		crates, _ := stacks[fromStack-1].MultiPop(nrToMove)
		for j := 0; j < nrToMove; j++ {
			stacks[toStack-1].Push(crates[j])
		}
	}

	fmt.Println("Stacks in the end")
	for i := 0; i < len(stacks); i++ {
		for j := 0; j < len(stacks[i]); j++ {
			fmt.Print(string(stacks[i][j]))
		}
		fmt.Println("")
	}

	// Get grates by peeking at first place of each stack and appending to result string
	crates := make([]byte, 0)
	for i := 0; i < len(stacks); i++ {
		topCrate, _ := stacks[i].Peek()
		crates = append(crates, topCrate)
	}

	fmt.Println("Crates:", string(crates))
}

func contains(range1 Range, range2 Range) bool {
	return range1.start <= range2.start && range1.end >= range2.end
}

func getRange(rangeAsString string) Range {
	rangeList := strings.Split(rangeAsString, "-")
	rangeItem := Range{start: getNumber(rangeList[0]), end: getNumber(rangeList[1])}
	return rangeItem
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
