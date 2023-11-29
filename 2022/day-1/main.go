package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	caloriesForElf := 0
	maxCalories := 0
	threeMaxCalories := []int{0, 0, 0}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			caloriesForElf = caloriesForElf + getNumber(line)
		} else {
			maxCalories = Max(caloriesForElf, maxCalories)
			keepMax(threeMaxCalories, caloriesForElf)
			caloriesForElf = 0
		}
	}

	fmt.Println("Max calories:", maxCalories)
	fmt.Println("Three max calories")
	for _, n := range threeMaxCalories {
		fmt.Println(n)
	}
	fmt.Println("Summed", arraySum(threeMaxCalories))
}

// Start at the beginning from biggest number
// Find first number smaller than num at index i
// Shift that number and the following numbers one step and remove last
// Place the new number on index i
// 5 3 1, 6
// 5 3 3
// 5 5 3
// 6 5 3
// 5 3 1, 4
// 5 3 3
// 5 4 3
// List sorted descending [3, 2, 1]
func keepMax(list []int, num int) {
	for i, n := range list {
		if num > n {
			shiftRight(list, i)
			list[i] = num
			break
		}
	}
}

// Shift items one position to the right ending at from pos, removing the last element
func shiftRight(list []int, pos int) {
	for i := len(list) - 1; i > pos; i-- {
		list[i] = list[i-1]
	}
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func arraySum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
