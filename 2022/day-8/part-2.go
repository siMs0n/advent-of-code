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

	highestScenicScore := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			left := viewingDistanceLeft(grid, x, y)
			right := viewingDistanceRight(grid, x, y)
			top := viewingDistanceTop(grid, x, y)
			bottom := viewingDistanceBottom(grid, x, y)
			/*fmt.Print(grid[y][x])
			fmt.Print(" ")
			fmt.Print(y)
			fmt.Print(",")
			fmt.Print(x)
			fmt.Print("left ", left)
			fmt.Print(" right ", right)
			fmt.Print(" top ", top)
			fmt.Println(" bottom", bottom)*/
			scenicScore := left * right * top * bottom
			highestScenicScore = Max(highestScenicScore, scenicScore)
		}
	}

	fmt.Println("Highest scenic score:", highestScenicScore)
}

func viewingDistanceLeft(grid [][]int, x int, y int) int {
	height := grid[y][x]
	viewingDistance := 0
	for i := x - 1; i >= 0; i-- {
		viewingDistance++
		if grid[y][i] >= height {
			break
		}
	}
	return viewingDistance
}

func viewingDistanceRight(grid [][]int, x int, y int) int {
	height := grid[y][x]
	viewingDistance := 0
	for i := x + 1; i < len(grid[y]); i++ {
		viewingDistance++
		if grid[y][i] >= height {
			break
		}
	}
	return viewingDistance
}

func viewingDistanceTop(grid [][]int, x int, y int) int {
	height := grid[y][x]
	viewingDistance := 0
	for i := y - 1; i >= 0; i-- {
		viewingDistance++
		if grid[i][x] >= height {
			break
		}
	}
	return viewingDistance
}

func viewingDistanceBottom(grid [][]int, x int, y int) int {
	height := grid[y][x]
	viewingDistance := 0
	for i := y + 1; i < len(grid); i++ {
		viewingDistance++
		if grid[i][x] >= height {
			break
		}
	}
	return viewingDistance
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
