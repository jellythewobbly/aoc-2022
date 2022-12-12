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
	fmt.Println("Part 2:")
	for _, row := range p2 {
		fmt.Println(row)
	}
}

func part1(input *[]string) int {
	sum := 0
	value := 1
	cycle := 0
	cycleSearch := 20
	for _, instruction := range *input {

		split := strings.Split(instruction, " ")
		command := split[0]
		if command == "noop" {
			processCycle(&cycle, &cycleSearch, &value, &sum)
			continue
		}
		addValue, _ := strconv.Atoi(split[1])
		processCycle(&cycle, &cycleSearch, &value, &sum)
		processCycle(&cycle, &cycleSearch, &value, &sum)
		value += addValue
	}
	return sum
}

func processCycle(cycle *int, cycleSearch *int, currentValue *int, sum *int) {
	*cycle++
	if *cycle == *cycleSearch {
		*sum += *cycle * *currentValue
		*cycleSearch += 40
	}
}

func part2(input *[]string) []string {
	spritePosition := 1
	cycle := 1
	crt := []string{""}
	currentIndex := 0
	for _, instruction := range *input {
		split := strings.Split(instruction, " ")
		command := split[0]
		if command == "noop" {
			drawPixel(&(crt[currentIndex]), cycle, spritePosition)
			processCycleCRT(&cycle, &crt, &currentIndex)
			continue
		}
		positionMove, _ := strconv.Atoi(split[1])
		drawPixel(&(crt[currentIndex]), cycle, spritePosition)
		processCycleCRT(&cycle, &crt, &currentIndex)
		drawPixel(&(crt[currentIndex]), cycle, spritePosition)
		processCycleCRT(&cycle, &crt, &currentIndex)
		spritePosition += positionMove
	}
	return crt
}

func drawPixel(current *string, cycle int, spritePosition int) {
	pixel := "."
	withinSprite := cycle-spritePosition <= 2 && cycle-spritePosition >= 0
	if withinSprite {
		pixel = "#"
	}
	*current = fmt.Sprintf("%v%v", *current, pixel)
}

func processCycleCRT(cycle *int, crt *([]string), currentIndex *int) {
	*cycle++
	if len((*crt)[*currentIndex]) == 40 {
		*crt = append(*crt, "")
		*currentIndex++
		*cycle = 1
	}
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
