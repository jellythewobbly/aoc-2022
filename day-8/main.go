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
	trees := 0
	for rowI, row := range *input {
		for colI, char := range row {
			if rowI == 0 || rowI == len(row)-1 || colI == 0 || colI == len(*input)-1 {
				trees++
				continue
			}
			treeHeight := int(char - '0')
			rowVisible, columnVisible := true, true

			for cI := 0; cI < len(row); cI++ {
				if cI == colI {
					if rowVisible {
						break
					}
					continue
				}
				currentTreeHeight := int(row[cI] - '0')
				if currentTreeHeight >= treeHeight {
					if cI > colI {
						rowVisible = false
						break
					}
					cI = colI
				}
			}

			if !rowVisible {
				for rI := 0; rI < len(*input); rI++ {
					if rI == rowI {
						if columnVisible {
							break
						}
						continue
					}
					currentTreeHeight := int((*input)[rI][colI] - '0')
					if currentTreeHeight >= treeHeight {
						if rI > rowI {
							columnVisible = false
							break
						}
						rI = rowI
					}
				}
			}
			if rowVisible || columnVisible {
				trees++
			}
		}
	}
	return trees
}

func part2(input *[]string) int {
	maxScore := 0
	for rowI, row := range *input {
		for colI, char := range row {
			if rowI == 0 || rowI == len(row)-1 || colI == 0 || colI == len(*input)-1 {
				continue
			}
			treeHeight := int(char - '0')
			up, down, left, right := 0, 0, 0, 0
			for rI := rowI - 1; rI >= 0; rI-- {
				currentTreeHeight := int((*input)[rI][colI] - '0')
				if currentTreeHeight >= treeHeight {
					if currentTreeHeight == treeHeight {
						up++
					}
					break
				}
				up++
			}

			for rI := rowI + 1; rI < len(*input); rI++ {
				currentTreeHeight := int((*input)[rI][colI] - '0')
				if currentTreeHeight >= treeHeight {
					if currentTreeHeight == treeHeight {
						down++
					}
					break
				}
				down++
			}

			for cI := colI - 1; cI >= 0; cI-- {
				currentTreeHeight := int(row[cI] - '0')
				if currentTreeHeight >= treeHeight {
					if currentTreeHeight == treeHeight {
						left++
					}
					break
				}
				left++
			}

			for cI := colI + 1; cI < len(row); cI++ {
				currentTreeHeight := int(row[cI] - '0')
				if currentTreeHeight >= treeHeight {
					if currentTreeHeight == treeHeight {
						right++
					}
					break
				}
				right++
			}

			currentScore := up * down * left * right
			if currentScore > maxScore {
				maxScore = currentScore
			}
		}
	}
	return maxScore
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
