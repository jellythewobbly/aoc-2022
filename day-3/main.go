package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input *[]string) int {
	sum := 0
	for _, line := range *input {
		halfLength := len(line) / 2
		first := line[:halfLength]
		second := line[halfLength:]
		charMap := map[byte]bool{}
		for i := range first {
			b := first[i]
			charMap[b] = true
		}
		for i := range second {
			searchByte := second[i]
			if _, found := charMap[searchByte]; found {
				if isUppercase(searchByte) {
					sum += int(searchByte - 38)
				} else {
					sum += int(searchByte - 96)
				}
				break
			}
		}
	}
	return sum
}

func part2(input *[]string) int {
	sum := 0
	for i := 0; i < len(*input)/3; i++ {
		first := (*input)[i*3+0]
		second := (*input)[i*3+1]
		third := (*input)[i*3+2]

		shortest, other1, other2 := findShortest(&first, &second, &third)
		shortestMap := parseToMap(shortest)
		other1Map := parseToMap(other1)
		other2Map := parseToMap(other2)

		common1 := findAllCommon(shortestMap, other1Map)
		common2 := findAllCommon(shortestMap, other2Map)
		for k := range *common1 {
			if _, found := (*common2)[k]; found {
				if isUppercase(k) {
					sum += int(k - 38)
				} else {
					sum += int(k - 96)
				}
				break
			}
		}
	}
	return sum
}

func isUppercase(char byte) bool {
	return char >= 65 && char <= 90
}

func findAllCommon(shorter *map[byte]bool, longer *map[byte]bool) *map[byte]bool {
	result := map[byte]bool{}
	for k := range *shorter {
		if _, found := (*longer)[k]; found {
			result[k] = true
		}
	}
	return &result
}

func findShortest(ptr1 *string, ptr2 *string, ptr3 *string) (*string, *string, *string) {
	l1 := len(*ptr1)
	l2 := len(*ptr2)
	l3 := len(*ptr3)

	if l1 < l2 && l1 < l3 {
		return ptr1, ptr2, ptr3
	}
	if l2 <= l1 && l2 < l3 {
		return ptr2, ptr1, ptr3
	}
	return ptr3, ptr1, ptr2
}

func parseToMap(input *string) *map[byte]bool {
	result := map[byte]bool{}
	for i := range *input {
		b := (*input)[i]
		result[b] = true
	}
	return &result
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
