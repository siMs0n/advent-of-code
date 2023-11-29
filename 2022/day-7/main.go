package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DirectoryOld struct {
	name     string
	children []interface{}
}

type File struct {
	name string
	size int
}

type FileSystemNode struct {
	name     string
	size     int
	children []*FileSystemNode
}

type Directory struct {
	name            string
	files           []*File
	subdirectories  []*Directory
	parentDirectory *Directory
	sizeOfContents  int // ?
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	root := Directory{name: "/"}
	currentDir := &root
	fmt.Println("current dir", currentDir)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if string(line[0:4]) == "$ cd" {
			// Change directory
			param := line[5:]
			currentDir = changeDirectory(&root, currentDir, param)
			fmt.Println("current dir", currentDir)
		} else if string(line[0:4]) == "$ ls" {
			// Do nothing
		} else if string(line[0:3]) == "dir" {
			// Directory, create a new one, maybe need to check for duplicates
			dir := Directory{name: line[4:], parentDirectory: currentDir}
			currentDir.subdirectories = append(currentDir.subdirectories, &dir)
		} else {
			// Is file
			f := strings.Split(line, " ")
			fileSize := getNumber(f[0])
			fileName := f[1]
			file := File{name: fileName, size: fileSize}
			currentDir.files = append(currentDir.files, &file)
		}
	}

	// After parseing the tree, calculate dir sizes
	totalSize := calculateSizeOfDirectories(&root)
	fmt.Println("Total size of directories:", totalSize)

	totalSum := sumOfDirectoriesAtSize(&root, 0)
	fmt.Println("Total sum of directories > 10000", totalSum)

	totalSpace := 70000000
	spaceNeeded := 30000000
	unusedSpace := totalSpace - root.sizeOfContents
	spaceToFree := spaceNeeded - unusedSpace
	fmt.Println("Space to free:", spaceToFree)
	directoryToDelete := findDirectoryToDelete(&root, &root, spaceToFree)
	fmt.Println("Directory to delete:", directoryToDelete.sizeOfContents)
}

func changeDirectory(root *Directory, currentDir *Directory, param string) *Directory {
	// TODO: Throw error if dir doesn't exist
	if param == "/" {
		return currentDir
	} else if param == ".." {
		return currentDir.parentDirectory
	} else {
		for _, subdirectory := range currentDir.subdirectories {
			if subdirectory.name == param {
				return subdirectory
			}
		}
	}
	return currentDir
}

func calculateSizeOfDirectories(directory *Directory) int {
	totalSize := 0

	for _, file := range directory.files {
		totalSize += file.size
	}
	for _, subdirectory := range directory.subdirectories {
		totalSize += calculateSizeOfDirectories(subdirectory)
	}
	directory.sizeOfContents = totalSize

	return totalSize
}

func sumOfDirectoriesAtSize(directory *Directory, sum int) int {
	_sum := sum
	for _, subdirectory := range directory.subdirectories {
		_sum += sumOfDirectoriesAtSize(subdirectory, sum)
	}
	if directory.sizeOfContents <= 100000 {
		return _sum + directory.sizeOfContents
	} else {
		return _sum
	}
}

func findDirectoryToDelete(directory *Directory, currentBest *Directory, spaceNeeded int) *Directory {
	for _, subdirectory := range directory.subdirectories {
		currentBest = findDirectoryToDelete(subdirectory, currentBest, spaceNeeded)
	}
	if directory.sizeOfContents >= spaceNeeded && directory.sizeOfContents < currentBest.sizeOfContents {
		return directory
	} else {
		return currentBest
	}
}

func getNumber(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return num
}
