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

	registerX := 1
	cycle := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		instruction := string(line[0:4])
		cycle++
		/*fmt.Println("Cycle", cycle)
		fmt.Println("Begin executing", line)*/
		drawPixel(cycle, registerX)
		if instruction == "addx" {
			value := getNumber(string(line[5:]))
			// addx takes two cycles
			cycle++
			drawPixel(cycle, registerX)
			registerX += value
		} else {
			// noop
		}
	}
	fmt.Println("")
	//fmt.Println("sumOfSignalStrenghts", sumOfSignalStrenghts)
}

func drawPixel(cycle int, spritePosition int) {
	pixelPosition := (cycle - 1) % 40
	isNewRow := pixelPosition == 0 && cycle > 40
	pixel := "."
	/*fmt.Println("cycle", cycle)
	fmt.Println("spritePosition", spritePosition)
	fmt.Println("pixelPosition", pixelPosition)*/

	if absInt(spritePosition-pixelPosition) <= 1 {
		pixel = "#"
	}
	if isNewRow {
		fmt.Println("")
	}
	fmt.Print(pixel)
}

// For now, consider the signal strength (the cycle number multiplied by the value of the X register)
// during the 20th cycle and every 40 cycles after that
// (that is, during the 20th, 60th, 100th, 140th, 180th, and 220th cycles).
func checkSignalStrength(cycle int, registerX int) int {
	if (cycle-20)%40 == 0 {
		signalStrength := cycle * registerX
		fmt.Println("Signal strenght", signalStrength)
		fmt.Println("At cycle", cycle)
		return signalStrength
	}
	return 0
}

func absInt(x int) int {
	return absDiffInt(x, 0)
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
