package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	tailPositions := make(map[string]bool)
	headPosition := Point{0, 0}
	tailPosition := Point{0, 0}
	tailPositions[getTailPositionKey(&tailPosition)] = true
	for fileScanner.Scan() {
		line := fileScanner.Text()
		direction := string(line[0])
		distance := getNumber(string(line[2:]))
		for i := 0; i < distance; i++ {
			moveHead(&headPosition, direction)
			moveTail(&headPosition, &tailPosition)
			tailPositions[getTailPositionKey(&tailPosition)] = true
		}
	}

	fmt.Println("Positions tail has visited:", len(tailPositions))
}

func moveHead(head *Point, direction string) {
	switch direction {
	case "U":
		head.y--
	case "D":
		head.y++
	case "L":
		head.x--
	case "R":
		head.x++
	}
}

// Moves the tail to catch up to the head
func moveTail(head *Point, tail *Point) {
	distanceX := absDiffInt(head.x, tail.x)
	distanceY := absDiffInt(head.y, tail.y)

	if (distanceX > 1 && distanceY == 1) || (distanceY > 1 && distanceX == 1) {
		// Diagonally move
		moveTailX(head, tail)
		moveTailY(head, tail)
	} else if distanceX > 1 {
		// move 1 X
		moveTailX(head, tail)
	} else if distanceY > 1 {
		// move 1 Y
		moveTailY(head, tail)
	}
}

func moveTailX(head *Point, tail *Point) {
	if head.x > tail.x {
		tail.x++
	} else {
		tail.x--
	}
}

func moveTailY(head *Point, tail *Point) {
	if head.y > tail.y {
		tail.y++
	} else {
		tail.y--
	}
}

func getTailPositionKey(tail *Point) string {
	return strconv.Itoa(tail.x) + "," + strconv.Itoa(tail.y)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
