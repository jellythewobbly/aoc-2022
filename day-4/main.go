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
	count := 0
	for _, line := range *input {
		split := strings.Split(line, ",")

		a := split[0]
		aSplit := strings.Split(a, "-")
		aStart, _ := strconv.Atoi(aSplit[0])
		aEnd, _ := strconv.Atoi(aSplit[1])

		b := split[1]
		bSplit := strings.Split(b, "-")
		bStart, _ := strconv.Atoi(bSplit[0])
		bEnd, _ := strconv.Atoi(bSplit[1])

		if isFullyContain(aStart, aEnd, bStart, bEnd) {
			count++
		}
	}
	return count
}

func part2(input *[]string) int {
	count := 0
	for _, line := range *input {
		split := strings.Split(line, ",")

		a := split[0]
		aSplit := strings.Split(a, "-")
		aStart, _ := strconv.Atoi(aSplit[0])
		aEnd, _ := strconv.Atoi(aSplit[1])

		b := split[1]
		bSplit := strings.Split(b, "-")
		bStart, _ := strconv.Atoi(bSplit[0])
		bEnd, _ := strconv.Atoi(bSplit[1])

		if hasOverlap(aStart, aEnd, bStart, bEnd) {
			count++
		}
	}
	return count
}

func isFullyContain(aStart int, aEnd int, bStart int, bEnd int) bool {
	diffA := bStart - aStart
	diffB := bEnd - aEnd

	var outStart, outEnd, inStart, inEnd int

	if diffA < diffB {
		outStart, outEnd, inStart, inEnd = bStart, bEnd, aStart, aEnd
	} else {
		outStart, outEnd, inStart, inEnd = aStart, aEnd, bStart, bEnd
	}

	return outStart <= inStart && outEnd >= inEnd
}

func hasOverlap(aStart int, aEnd int, bStart int, bEnd int) bool {
	return aStart <= bEnd && bStart <= aEnd
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
