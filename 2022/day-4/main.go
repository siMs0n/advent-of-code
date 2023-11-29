package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	assignmentPairContains := 0
	assignmentPairOverlaps := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		pairs := strings.Split(line, ",")
		first, second := pairs[0], pairs[1]
		range1 := getRange(first)
		range2 := getRange(second)
		if contains(range1, range2) || contains(range2, range1) {
			assignmentPairContains += 1
		}
		if overlaps(range1, range2) {
			assignmentPairOverlaps += 1
		}
	}

	fmt.Println("Assignment pair contains:", assignmentPairContains)
	fmt.Println("Assignment pairs overlaps:", assignmentPairOverlaps)
}

func contains(range1 Range, range2 Range) bool {
	return range1.start <= range2.start && range1.end >= range2.end
}

func overlaps(range1 Range, range2 Range) bool {
	return range1.start >= range2.start && range1.start <= range2.end ||
		range2.start >= range1.start && range2.start <= range1.end
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
