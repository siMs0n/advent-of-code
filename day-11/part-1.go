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
	items            []int
	operation        string
	operand          int // If 0, operand is itself
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
			for _, item := range itemsAsStr {
				monkeys[currentMonkey].items = append(monkeys[currentMonkey].items, getNumber(item))
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

	fmt.Println("monkeys", monkeys)

	// Perform 20 rounds
	for round := 1; round <= 20; round++ {
		for i := range monkeys {
			for _, item := range monkeys[i].items {
				monkeys[i].itemsInspected++
				item = getNewItemWorryLevel(item, &monkeys[i])
				// Monkey gets bored and worry level drops
				item = item / 3
				if item%monkeys[i].testArg == 0 {
					passItem(item, monkeys[i].testPassedMonkey, monkeys)
				} else {
					passItem(item, monkeys[i].testFailedMonkey, monkeys)
				}
			}
			monkeys[i].items = make([]int, 0)
		}
	}

	mostInspected := 0
	secondMostInspected := 0
	for _, monkey := range monkeys {
		if monkey.itemsInspected > mostInspected {
			secondMostInspected = mostInspected
			mostInspected = monkey.itemsInspected
		} else if monkey.itemsInspected > mostInspected {
			secondMostInspected = monkey.itemsInspected
		}
	}

	monkeyBusiness := mostInspected * secondMostInspected
	fmt.Println("Money business", monkeyBusiness)
}

func passItem(item int, recipientIndex int, monkeys []Monkey) {
	monkeys[recipientIndex].items = append(monkeys[recipientIndex].items, item)
}

func getNewItemWorryLevel(itemWorryLevel int, monkey *Monkey) int {
	if monkey.operation == "+" {
		return itemWorryLevel + monkey.operand
	}
	// Operation is *
	if monkey.operand == 0 {
		// If 0, operand is itself
		return itemWorryLevel * itemWorryLevel
	}
	return itemWorryLevel * monkey.operand
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
