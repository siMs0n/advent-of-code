package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	startOfBuffer := -1
	set := make(map[string]void)
	set[string(line[0])] = member
	set[string(line[1])] = member
	set[string(line[2])] = member
	for i := 3; i < len(line); i++ {
		set[string(line[i])] = member
		if len(set) == 4 {
			startOfBuffer = i
			break
		}

		delete(set, string(line[i-3]))
	}

	fmt.Println("startOfBuffer:", startOfBuffer)
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
