package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	input := getInput()

	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)

}

func part1(input *string) int {
	hm := make(map[byte]int)
	for slow, fast := 0, 0; fast < len(*input); fast++ {
		charRight := (*input)[fast]
		if p, ok := hm[charRight]; ok {
			slow = int(math.Max(float64(slow), float64(p+1)))
		}
		hm[charRight] = fast
		length := fast - slow + 1
		if length == 4 {
			return fast + 1
		}
	}
	return 0
}

func part2(input *string) int {
	hm := make(map[byte]int)
	for slow, fast := 0, 0; fast < len(*input); fast++ {
		charRight := (*input)[fast]
		if p, ok := hm[charRight]; ok {
			slow = int(math.Max(float64(slow), float64(p+1)))
		}
		hm[charRight] = fast
		length := fast - slow + 1
		if length == 14 {
			return fast + 1
		}
	}
	return 0
}

func getInput() string {
	fileReader, err := os.Open("./input")
	if err != nil {
		fmt.Println("error in reading file")
		return ""
	}

	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)

	scanner.Scan()
	res := scanner.Text()
	return res
}
