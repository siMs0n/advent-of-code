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
	sumOfSignalStrenghts := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		instruction := string(line[0:4])
		cycle++
		sumOfSignalStrenghts += checkSignalStrength(registerX, cycle)
		if instruction == "addx" {
			value := getNumber(string(line[5:]))
			// addx takes two cycles
			cycle++
			sumOfSignalStrenghts += checkSignalStrength(registerX, cycle)
			registerX += value
		} else {
			// noop
		}
	}

	fmt.Println("sumOfSignalStrenghts", sumOfSignalStrenghts)
}

// For now, consider the signal strength (the cycle number multiplied by the value of the X register)
// during the 20th cycle and every 40 cycles after that
// (that is, during the 20th, 60th, 100th, 140th, 180th, and 220th cycles).
func checkSignalStrength(registerX int, cycle int) int {
	if (cycle-20)%40 == 0 {
		signalStrength := cycle * registerX
		fmt.Println("Signal strenght", signalStrength)
		fmt.Println("At cycle", cycle)
		return signalStrength
	}
	return 0
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
