package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxDistance = 99999

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	grid := make([][]rune, 0)
	distances := make([][]int, 0)
	//unvisitedNodes := make([]string, 0)
	nodes := make(map[string]bool)
	nodesAtElevationA := make([]string, 0)

	lineNr := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, make([]rune, 0))
		distances = append(distances, make([]int, 0))
		for i, ch := range line {
			grid[lineNr] = append(grid[lineNr], ch)
			if ch == []rune("E")[0] { // Start at E
				distances[lineNr] = append(distances[lineNr], 0)
			} else {
				distances[lineNr] = append(distances[lineNr], MaxDistance)
			}
			nodeKey := getNodeKey(lineNr, i)
			nodes[nodeKey] = false

			if ch == []rune("a")[0] {
				nodesAtElevationA = append(nodesAtElevationA, nodeKey)
			}
		}
		lineNr++
	}

	unvisitedNodes := len(nodes)
	for unvisitedNodes > 0 {
		currentNodeKey := minDistance(distances, nodes)

		currentCoords := getNodeCoordinates(currentNodeKey)
		currentDistance := distances[currentCoords[0]][currentCoords[1]]

		neighbours := getUnvisitedNeighbours(grid, nodes, currentNodeKey)

		// If currdistance + 1 is less than their distance, update their distance
		for _, n := range neighbours {
			coords := getNodeCoordinates(n)
			distance := distances[coords[0]][coords[1]]
			if currentDistance+1 < distance {
				distances[coords[0]][coords[1]] = currentDistance + 1
			}
		}

		nodes[currentNodeKey] = true
		unvisitedNodes--
	}

	shortestDistanceToE := MaxDistance
	for _, node := range nodesAtElevationA {
		coords := getNodeCoordinates(node)
		distance := distances[coords[0]][coords[1]]
		if distance <= shortestDistanceToE {
			shortestDistanceToE = distance
		}
	}
	fmt.Println("Shortest distance to E from any a:", shortestDistanceToE)
}

func getUnvisitedNeighbours(grid [][]rune, nodes map[string]bool, currentNode string) []string {
	coords := getNodeCoordinates(currentNode)
	currentHeight := getHeight(grid[coords[0]][coords[1]])
	neighbours := make([]string, 0)
	// Upper
	upper := getNodeKey(coords[0]-1, coords[1])
	if coords[0]-1 >= 0 && !nodes[upper] && getHeightDiff(grid, currentHeight, upper) >= -1 {
		neighbours = append(neighbours, upper)
	}
	// Lower
	lower := getNodeKey(coords[0]+1, coords[1])
	if coords[0]+1 < len(grid) && !nodes[lower] && getHeightDiff(grid, currentHeight, lower) >= -1 {
		neighbours = append(neighbours, lower)
	}
	// Left
	left := getNodeKey(coords[0], coords[1]-1)
	if coords[1]-1 >= 0 && !nodes[left] && getHeightDiff(grid, currentHeight, left) >= -1 {
		neighbours = append(neighbours, left)
	}
	// Right
	right := getNodeKey(coords[0], coords[1]+1)
	if coords[1]+1 < len(grid[coords[0]]) && !nodes[right] && getHeightDiff(grid, currentHeight, right) >= -1 {
		neighbours = append(neighbours, right)
	}
	return neighbours
}

func minDistance(distances [][]int, nodes map[string]bool) string {
	minDistance := MaxDistance
	minNode := ""
	for node, visited := range nodes {
		if !visited {
			coords := getNodeCoordinates(node)
			distance := distances[coords[0]][coords[1]]
			if distance <= minDistance {
				minDistance = distance
				minNode = node
			}
		}
	}
	return minNode
}

func getNodeKey(lineNr int, x int) string {
	return strconv.Itoa(lineNr) + "," + strconv.Itoa(x)
}

func getNodeCoordinates(node string) []int {
	coordinates := strings.Split(node, ",")
	return []int{getNumber(coordinates[0]), getNumber(coordinates[1])}
}

func getHeightDiff(grid [][]rune, height int, node string) int {
	coords := getNodeCoordinates(node)
	return getHeight(grid[coords[0]][coords[1]]) - height
}

func getHeight(ch rune) int {
	charCodeLowerCaseA := []rune("a")[0]
	lowerCaseStartingRangePrio := 1

	if ch == []rune("S")[0] {
		return 1
	}
	if ch == []rune("E")[0] {
		return 26
	}

	return int(ch-charCodeLowerCaseA) + lowerCaseStartingRangePrio
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
