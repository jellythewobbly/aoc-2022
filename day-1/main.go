package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v", p1)

	p2 := part2(&input)
	fmt.Printf("Part 2: %v", p2)
}

func part1(input *[]string) int64 {
	var max, currentMax int64
	for _, v := range *input {
		if v == "" {
			if currentMax > max {
				max = currentMax
			}
			currentMax = 0
		} else {
			i, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				currentMax += i
			} else {
				fmt.Println("ParseInt err")
			}
		}
	}

	if currentMax > max {
		return currentMax
	}
	return max
}

func part2(input *[]string) int64 {
	var h1, h2, h3, currentMax int64
	for _, v := range *input {
		if v == "" {
			compare(currentMax, &h1, &h2, &h3)
			currentMax = 0
		} else {
			i, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				currentMax += i
			} else {
				fmt.Println("ParseInt err")
			}
		}
	}

	compare(currentMax, &h1, &h2, &h3)
	return h1 + h2 + h3
}

func compare(num int64, h1 *int64, h2 *int64, h3 *int64) {
	if num > *h1 {
		*h3 = *h2
		*h2 = *h1
		*h1 = num
	} else if num > *h2 {
		*h3 = *h2
		*h2 = num
	} else if num > *h3 {
		*h3 = num
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
