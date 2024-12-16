package hoof_it

import (
	"advent_of_code_2024/src/utils"
	"fmt"
	"strconv"
)

func findTrailHeads(expected int, coord utils.Coord, mapping []string) []utils.Coord {
	current, _ := strconv.Atoi(string(mapping[coord.X][coord.Y]))
	if current == expected && current == 9 {
		return []utils.Coord{coord}
	} else if current != expected {
		return []utils.Coord{}
	}
	var directions = []utils.Coord{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}
	var output []utils.Coord
	for _, direction := range directions {
		newCoord := direction.AddCoord(coord)
		if utils.InBounds(newCoord.X, newCoord.Y, len(mapping), len(mapping[0])) {
			output = append(output, findTrailHeads(expected+1, newCoord, mapping)...)
		}
	}
	return output
}

// contains if an int is contained inside []int
func contains(slice []utils.Coord, value utils.Coord) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func countDistinct(coords []utils.Coord) int {
	var output []utils.Coord
	for i := 0; i < len(coords); i++ {
		if !contains(output, coords[i]) {
			output = append(output, coords[i])
		}
	}
	return len(output)
}

func analyseMap(mapping []string, partTwo bool) int {
	output := 0
	for i := 0; i < len(mapping); i++ {
		for j := 0; j < len(mapping[i]); j++ {
			if string(mapping[i][j]) == "0" {
				trails := findTrailHeads(0, utils.Coord{X: i, Y: j}, mapping)
				if partTwo {
					output += len(trails)
				} else {
					output += countDistinct(trails)
				}
			}
		}
	}
	return output
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	output := analyseMap(data, false)
	fmt.Println(output)
	return output
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	output := analyseMap(data, true)
	fmt.Println(output)
	return output
}
