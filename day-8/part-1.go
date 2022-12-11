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

	grid := make([][]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numberRow := make([]int, 0)
		for i := 0; i < len(line); i++ {
			numberRow = append(numberRow, getNumber(string(line[i])))
		}
		grid = append(grid, numberRow)
	}

	// Print grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println("")
	}

	totalAmountVisible := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			left := isVisibleLeft(grid, x, y)
			right := isVisibleRight(grid, x, y)
			top := isVisibleTop(grid, x, y)
			bottom := isVisibleBottom(grid, x, y)
			if left || right || top || bottom {
				totalAmountVisible++
			}
		}
	}

	fmt.Println("Amount of trees visible:", totalAmountVisible)
}

func isVisibleLeft(grid [][]int, x int, y int) bool {
	height := grid[y][x]
	for i := x - 1; i >= 0; i-- {
		if grid[y][i] >= height {
			return false
		}
	}
	return true
}

func isVisibleRight(grid [][]int, x int, y int) bool {
	height := grid[y][x]
	for i := x + 1; i < len(grid[y]); i++ {
		if grid[y][i] >= height {
			return false
		}
	}
	return true
}

func isVisibleTop(grid [][]int, x int, y int) bool {
	height := grid[y][x]
	for i := y - 1; i >= 0; i-- {
		if grid[i][x] >= height {
			return false
		}
	}
	return true
}

func isVisibleBottom(grid [][]int, x int, y int) bool {
	height := grid[y][x]
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] >= height {
			return false
		}
	}
	return true
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
