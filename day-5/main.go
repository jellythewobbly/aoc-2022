package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	starting1, steps1 := getInput()
	p1 := part1(starting1, steps1)
	fmt.Printf("Part 1: %v\n", p1)

	starting2, steps2 := getInput()
	p2 := part2(starting2, steps2)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(starting [][]rune, steps []string) string {
	for _, step := range steps {
		split := strings.Split(step, " ")
		quantity, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])
		height := len(starting[from-1])
		itemsToMove := starting[from-1][height-quantity:]
		starting[from-1] = starting[from-1][:height-quantity]

		for i := len(itemsToMove) - 1; i >= 0; i-- {
			item := itemsToMove[i]
			starting[to-1] = append(starting[to-1], item)
		}
	}

	var result strings.Builder
	for _, slice := range starting {
		top := slice[len(slice)-1]
		result.WriteString(string(top))
	}

	return result.String()
}

func part2(starting [][]rune, steps []string) string {
	for _, step := range steps {
		split := strings.Split(step, " ")
		quantity, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])
		height := len(starting[from-1])
		itemsToMove := starting[from-1][height-quantity:]
		starting[from-1] = starting[from-1][:height-quantity]

		for _, item := range itemsToMove {
			starting[to-1] = append(starting[to-1], item)
		}
	}

	var result strings.Builder
	for _, slice := range starting {
		top := slice[len(slice)-1]
		result.WriteString(string(top))
	}

	return result.String()
}

func getInput() ([][]rune, []string) {
	fileReader, err := os.Open("./input")
	if err != nil {
		fmt.Println("error in reading file")
		return [][]rune{}, []string{}
	}

	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)

	var starting []string
	var steps []string
	areSteps := false

	for scanner.Scan() {
		if !areSteps {
			if scanner.Text() == "" {
				areSteps = true
				continue
			}
			starting = append(starting, scanner.Text())
		} else {
			steps = append(steps, scanner.Text())
		}
	}

	var start [][]rune

	columns := starting[len(starting)-1]
	for column, char := range columns {
		if char != 32 {
			var val []rune
			for rowIndex := len(starting) - 2; rowIndex >= 0; rowIndex-- {
				a := starting[rowIndex][column]
				if a != 32 {
					val = append(val, rune(starting[rowIndex][column]))
				}
			}
			start = append(start, val)
		}
	}

	return start, steps
}
