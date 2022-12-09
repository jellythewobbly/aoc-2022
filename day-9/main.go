package main

import (
	"bufio"
	"fmt"
	"math"
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
	headX, headY := 0, 0
	tailX, tailY := 0, 0
	positionMap := map[string]bool{"0,0": true}
	for _, row := range *input {
		split := strings.Split(row, " ")
		direction := split[0]
		steps, _ := strconv.Atoi(split[1])
		for i := 0; i < steps; i++ {
			move(direction, &headX, &headY)
			tailMove := computeStep(&headX, &headY, &tailX, &tailY)
			if len(tailMove) > 0 {
				move(tailMove, &tailX, &tailY)
				positionMap[fmt.Sprintf("%v,%v", tailX, tailY)] = true
			}
		}
	}
	return len(positionMap)
}

func part2(input *[]string) int {
	var knots [][]int
	for i := 0; i < 10; i++ {
		knots = append(knots, []int{0, 0})
	}
	positionMap := map[string]bool{"0,0": true}
	for _, row := range *input {
		split := strings.Split(row, " ")
		direction := split[0]
		steps, _ := strconv.Atoi(split[1])
		for i := 0; i < steps; i++ {
			move(direction, &(knots[0][0]), &(knots[0][1]))
			for knotI := 1; knotI < len(knots); knotI++ {
				prevX, prevY := &(knots[knotI-1][0]), &(knots[knotI-1][1])
				knotX, knotY := &(knots[knotI][0]), &(knots[knotI][1])
				nextMove := computeStep(prevX, prevY, knotX, knotY)
				if len(nextMove) > 0 {
					move(nextMove, knotX, knotY)
					if knotI == len(knots)-1 {
						positionMap[fmt.Sprintf("%v,%v", *knotX, *knotY)] = true
					}
				}
			}
		}
	}
	return len(positionMap)
}

func move(direction string, x *int, y *int) {
	if direction == "D" {
		*y--
	} else if direction == "U" {
		*y++
	} else if direction == "L" {
		*x--
	} else if direction == "R" {
		*x++
	} else if direction == "DR" {
		*y--
		*x++
	} else if direction == "DL" {
		*y--
		*x--
	} else if direction == "UR" {
		*y++
		*x++
	} else if direction == "UL" {
		*y++
		*x--
	}
}

func computeStep(headX *int, headY *int, tailX *int, tailY *int) string {
	if isWithin1Cell(headX, headY, tailX, tailY) {
		return ""
	}
	if *headX == *tailX {
		if *headY > *tailY {
			return "U"
		}
		return "D"
	}

	if *headY == *tailY {
		if *headX > *tailX {
			return "R"
		}
		return "L"
	}
	if *headX > *tailX {
		if *headY > *tailY {
			return "UR"
		}
		return "DR"
	}
	if *headY > *tailY {
		return "UL"
	}
	return "DL"
}

func isWithin1Cell(headX *int, headY *int, tailX *int, tailY *int) bool {
	return int(math.Abs(float64(*headX-*tailX))) <= 1 && int(math.Abs(float64(*headY-*tailY))) <= 1
}

// func part2(input *[]string) int {
// }

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
