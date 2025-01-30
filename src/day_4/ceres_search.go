package ceres_search

import (
	"advent_of_code_2024/src/utils"
)

type coord struct {
	x int
	y int
}

// checkForWord Given a certain direction and starting coordinate, checks the map for a given word
func checkForWord(startCoord coord, direction coord, letterMap [][]string, target string) bool {
	for index := 0; index < len(target); index++ {
		newX := startCoord.x + (direction.x * index)
		newY := startCoord.y + (direction.y * index)
		if newX >= 0 && newX < len(letterMap[0]) && newY >= 0 && newY < len(letterMap) {
			if letterMap[newX][newY] != string(target[index]) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// checkDirections Extends out the search in all directions from the starting coordinate
func checkDirections(startCoord coord, target string, letterMap [][]string) int {
	output := 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if checkForWord(startCoord, coord{x, y}, letterMap, target) {
				output++
			}
		}
	}
	return output
}

// wordSearch Counts the number of occurrences of a word within the word search map
func wordSearch(letterMap [][]string, target string) int {
	output := 0
	for i := 0; i < len(letterMap); i++ {
		line := letterMap[i]
		for j := 0; j < len(line); j++ {
			if line[j] == string(target[0]) {
				// finds the first letter
				output += checkDirections(coord{i, j}, target, letterMap)
			}
		}
	}
	return output
}

// checkExpansion Checks if the word search contains an 'X' of correct letters in the map
func checkExpansion(startCoord coord, letterMap [][]string) int {
	forwardLeg := letterMap[startCoord.x-1][startCoord.y-1] + letterMap[startCoord.x+1][startCoord.y+1]
	backLeg := letterMap[startCoord.x+1][startCoord.y-1] + letterMap[startCoord.x-1][startCoord.y+1]
	if (forwardLeg == "MS" || forwardLeg == "SM") && (backLeg == "MS" || backLeg == "SM") {
		return 1
	}
	return 0
}

// crossSearch Counts the number of words that form an 'X' of correct letters in the map
func crossSearch(letterMap [][]string) int {
	output := 0
	for i := 1; i < len(letterMap)-1; i++ {
		line := letterMap[i]
		for j := 1; j < len(line)-1; j++ {
			if line[j] == "A" {
				// finds the first letter
				output += checkExpansion(coord{i, j}, letterMap)
			}
		}
	}
	return output
}

func SolvePartOne(fileName string) int {
	//fileName := "src/day_4/problem_input_1.txt"
	data := utils.ReadFileToUnits(fileName)
	count := wordSearch(data, "XMAS")
	return count
}

func SolvePartTwo(fileName string) int {
	//fileName := "src/day_4/problem_input_1.txt"
	data := utils.ReadFileToUnits(fileName)
	count := crossSearch(data)
	return count
}
