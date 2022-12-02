package main

import (
	"bufio"
	"fmt"
	"os"
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
	handScoreMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	totalScore := 0
	for _, round := range *input {
		split := strings.Split(round, " ")
		opponent := split[0]
		player := split[1]
		roundScore := handScoreMap[player] + parseOutcomeScore(player, opponent)
		totalScore += roundScore
	}

	return totalScore
}

func part2(input *[]string) int {
	outcomeScoreMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	totalScore := 0
	for _, round := range *input {
		split := strings.Split(round, " ")
		opponent := split[0]
		outcome := split[1]
		roundScore := outcomeScoreMap[outcome] + parsePlayerScore(outcome, opponent)
		totalScore += roundScore
	}

	return totalScore

}

func parseOutcomeScore(player string, opponent string) int {
	winMap := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	loseMap := map[string]string{
		"X": "B",
		"Y": "C",
		"Z": "A",
	}

	if o, _ := winMap[player]; o == opponent {
		return 6
	} else if o, _ := loseMap[player]; o == opponent {
		return 0
	}
	return 3
}

func parsePlayerScore(outcome string, opponent string) int {
	rockMap := map[string]string{
		"X": "B",
		"Y": "A",
		"Z": "C",
	}

	paperMap := map[string]string{
		"X": "C",
		"Y": "B",
		"Z": "A",
	}

	if o, _ := rockMap[outcome]; o == opponent {
		return 1
	} else if o, _ := paperMap[outcome]; o == opponent {
		return 2
	}
	return 3
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
