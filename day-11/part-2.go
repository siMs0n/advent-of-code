package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	items            [][]int // Each item is a list of remainders for each monkey
	operation        string
	operand          int // If 0, operand is itself
	remainder        int
	testArg          int
	testPassedMonkey int
	testFailedMonkey int
	itemsInspected   int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	// Parse monkeys
	monkeys := make([]Monkey, 0)
	currentMonkey := 0
	startingItems := make([][]int, 0)

	rDigits, _ := regexp.Compile("\\d+")
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Monkey") {
			monkey := new(Monkey)
			monkeys = append(monkeys, *monkey)
			currentMonkey = len(monkeys) - 1
		} else if strings.HasPrefix(line, "Starting items") {
			itemsAsStr := rDigits.FindAllString(line, -1)
			items := make([]int, 0)
			startingItems = append(startingItems, items)
			for _, item := range itemsAsStr {
				startingItems[len(startingItems)-1] = append(startingItems[len(startingItems)-1], getNumber(item))
			}
		} else if strings.HasPrefix(line, "Operation") {
			// Get operation and argument
			indexOfPlus := strings.Index(line, "+")
			if indexOfPlus > 0 {
				monkeys[currentMonkey].operation = "+"
			} else {
				monkeys[currentMonkey].operation = "*"
			}
			operand, err := strconv.Atoi(rDigits.FindString(line))
			if err == nil {
				monkeys[currentMonkey].operand = operand
			}
		} else if strings.HasPrefix(line, "Test") {
			monkeys[currentMonkey].testArg = getNumber(rDigits.FindString(line))
		} else if strings.HasPrefix(line, "If true") {
			monkeys[currentMonkey].testPassedMonkey = getNumber(rDigits.FindString(line))
		} else if strings.HasPrefix(line, "If false") {
			monkeys[currentMonkey].testFailedMonkey = getNumber(rDigits.FindString(line))
		}
	}

	// Getting remainders for each startingItem for each monkey
	for i, items := range startingItems {
		for _, item := range items {
			remainders := make([]int, 0)
			for _, m := range monkeys {
				remainders = append(remainders, getRemainder(item, m.testArg))
			}
			monkeys[i].items = append(monkeys[i].items, remainders)
		}
	}

	// Perform 10000 rounds
	for round := 1; round <= 10000; round++ {
		for i := range monkeys {
			for _, item := range monkeys[i].items {
				monkeys[i].itemsInspected++
				item = getNewItemWorryLevel(item, &monkeys[i], monkeys)
				if item[i]%monkeys[i].testArg == 0 {
					passItem(item, monkeys[i].testPassedMonkey, monkeys)
				} else {
					passItem(item, monkeys[i].testFailedMonkey, monkeys)
				}
			}
			monkeys[i].items = make([][]int, 0)
		}
	}

	mostInspected := 0
	secondMostInspected := 0
	for _, monkey := range monkeys {
		if monkey.itemsInspected > mostInspected {
			secondMostInspected = mostInspected
			mostInspected = monkey.itemsInspected
		} else if monkey.itemsInspected > secondMostInspected {
			secondMostInspected = monkey.itemsInspected
		}
	}

	monkeyBusiness := mostInspected * secondMostInspected
	fmt.Println("Money business", monkeyBusiness)
}

func getRemainder(num int, modNum int) int {
	res := num % modNum
	if res == 0 {
		return num
	} else {
		return res
	}
}

func passItem(item []int, recipientIndex int, monkeys []Monkey) {
	monkeys[recipientIndex].items = append(monkeys[recipientIndex].items, item)
}

func getNewItemWorryLevel(itemWorryLevels []int, monkey *Monkey, monkeys []Monkey) []int {
	for i, itemWorryLevel := range itemWorryLevels {

		if monkey.operation == "+" {
			itemWorryLevels[i] = getRemainder(itemWorryLevel+monkey.operand, monkeys[i].testArg)
		} else {
			// Operation is *
			if monkey.operand == 0 {
				itemWorryLevels[i] = getRemainder(itemWorryLevel*itemWorryLevel, monkeys[i].testArg)
			} else {
				itemWorryLevels[i] = getRemainder(itemWorryLevel*monkey.operand, monkeys[i].testArg)
			}
		}
	}
	return itemWorryLevels
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
