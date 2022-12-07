package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input *[]string) int {
	hm := make(map[string]int)
	pathStack := []string{}
	index := 0
	for index < len(*input) {
		parseCommand(input, &index, &hm, &pathStack)
	}
	totalSize := 0

	for _, v := range hm {
		if v <= 100000 {
			totalSize += v
		}
	}
	return totalSize
}

func part2(input *[]string) int {
	hm := make(map[string]int)
	pathStack := []string{}
	index := 0
	for index < len(*input) {
		parseCommand(input, &index, &hm, &pathStack)
	}
	totalDiskSpace := 70000000
	rootSize := hm["/"]
	unusedSpace := totalDiskSpace - rootSize
	spaceNeededForUpdate := 30000000
	needToDelete := spaceNeededForUpdate - unusedSpace

	smallest := rootSize
	for _, size := range hm {
		if size >= needToDelete && size < smallest {
			smallest = size
		}
	}
	return smallest
}

func parseCommand(input *[]string, index *int, hm *map[string]int, pathStack *[]string) {
	currentLine := (*input)[*index]
	split := strings.Split(currentLine, " ")
	commandName := split[1]
	if commandName == "cd" {
		pathName := split[2]
		parseCd(pathName, pathStack)
		*index++
	} else if commandName == "ls" {
		*index++
		parseLs(input, index, hm, pathStack)
	}
}

func parseCd(path string, pathStack *[]string) {
	if path == ".." {
		*pathStack = (*pathStack)[:len(*pathStack)-1]
	} else if path == "/" {
		*pathStack = append((*pathStack), path)
	} else {
		currentPath := (*pathStack)[len(*pathStack)-1]
		*pathStack = append((*pathStack), fmt.Sprintf("%v/%v", currentPath, path))
	}
}

func parseLs(input *[]string, index *int, hm *map[string]int, pathStack *[]string) {
	totalSize := 0

	for isStillValidLs(input, index) {
		currentLine := (*input)[*index]
		split := strings.Split(currentLine, " ")
		if split[0] != "dir" {
			fileSize, _ := strconv.Atoi(split[0])
			totalSize += fileSize
		}
		*index++
	}

	for pathIndex := len(*pathStack) - 1; pathIndex >= 0; pathIndex-- {
		path := (*pathStack)[pathIndex]
		(*hm)[path] += totalSize
	}
}

func isStillValidLs(input *[]string, index *int) bool {
	if *index == len(*input) {
		return false
	}
	currentLine := (*input)[*index]
	return !isCommand(currentLine)
}

func isCommand(input string) bool {
	return string(input[0]) == "$"
}

func getInput() []string {
	fileReader, err := os.Open("./input")
	if err != nil {
		fmt.Println("error in reading file")
		return []string{}
	}

	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)

	var res []string

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}
