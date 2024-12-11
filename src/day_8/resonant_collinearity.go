package resonant_collinearity

import (
	"advent_of_code_2024/src/utils"
	"strconv"
)

type nodeData struct {
	frequency rune
	locations []utils.Coord
}

// inNodeList checks if a node (antenna) with that frequency already exists
func inNodeList(nodeList []nodeData, freq rune) int {
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i].frequency == freq {
			return i
		}
	}
	return -1
}

// buildAntennaMapping Builds a map of all the antennas on the grid grouped by frequency.
func buildAntennaMapping(data []string) []nodeData {
	var output []nodeData
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] != '.' {
				index := inNodeList(output, rune(data[i][j]))
				if index != -1 {
					output[index].locations = append(output[index].locations, utils.Coord{X: i, Y: j})
				} else {
					location := utils.Coord{X: i, Y: j}
					output = append(output, nodeData{frequency: rune(data[i][j]), locations: []utils.Coord{location}})
				}
			}
		}
	}
	return output
}

// inBounds checks if a given coord is within a grid
func inBounds(coord utils.Coord, xLimit int, yLimit int) bool {
	return coord.X >= 0 && coord.X < xLimit && coord.Y >= 0 && coord.Y < yLimit
}

// inBoundsAdv checks if a given set of points is within a grid
func inBoundAdv(x int, y int, xLimit int, yLimit int) bool {
	return x >= 0 && x < xLimit && y >= 0 && y < yLimit
}

// coordToStr Converts a coord object to a string type.
func coordToStr(coord utils.Coord) string {
	x := strconv.Itoa(coord.X)
	y := strconv.Itoa(coord.Y)
	return x + ":" + y
}

// contains checks in a string is in a list of strings
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// analyseNodeMapping iterates through all combinations within a given frequency to generate their resonant frequencies
func analyseNodeMapping(nodeList []nodeData, xLimit int, yLimit int, useAdvanced bool) []string {
	var output []string
	for x := 0; x < len(nodeList); x++ {
		locationList := nodeList[x].locations
		for i := 0; i < len(locationList)-1; i++ {
			for j := i + 1; j < len(locationList); j++ {
				var resonateFrequencies []utils.Coord
				if useAdvanced {
					resonateFrequencies = GenerateFrequenciesAdv(locationList[i], locationList[j], xLimit, yLimit)
				} else {
					resonateFrequencies = GenerateFrequencies(locationList[i], locationList[j])
				}
				for _, freq := range resonateFrequencies {
					if inBounds(freq, xLimit, yLimit) {
						stringFormat := coordToStr(freq)
						if !contains(output, stringFormat) {
							output = append(output, stringFormat)
						}
					}
				}
			}
		}
	}
	return output
}

// GenerateFrequencies Generates two frequencies in a line between two points at equal distances away
func GenerateFrequencies(antennaA utils.Coord, antennaB utils.Coord) []utils.Coord {
	deltaX := antennaA.X - antennaB.X
	deltaY := antennaA.Y - antennaB.Y
	positionRelativeA := utils.Coord{X: antennaA.X + deltaX, Y: antennaA.Y + deltaY}
	positionRelativeB := utils.Coord{X: antennaB.X - deltaX, Y: antennaB.Y - deltaY}
	return []utils.Coord{positionRelativeA, positionRelativeB}

}

// GenerateFrequenciesAdv For pt2, Generates every frequency in a line between two points including the antenna themselves
func GenerateFrequenciesAdv(antennaA utils.Coord, antennaB utils.Coord, xLimit int, yLimit int) []utils.Coord {
	deltaX := antennaA.X - antennaB.X
	deltaY := antennaA.Y - antennaB.Y

	var output []utils.Coord

	x := antennaA.X
	y := antennaA.Y
	for inBoundAdv(x, y, xLimit, yLimit) {
		output = append(output, utils.Coord{X: x, Y: y})
		x += deltaX
		y += deltaY
	}
	x = antennaB.X
	y = antennaB.Y
	for inBoundAdv(x, y, xLimit, yLimit) {
		output = append(output, utils.Coord{X: x, Y: y})
		x -= deltaX
		y -= deltaY
	}
	return output
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	nodeMapping := buildAntennaMapping(data)
	resultList := analyseNodeMapping(nodeMapping, len(data), len(data[0]), false)
	return len(resultList)
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	nodeMapping := buildAntennaMapping(data)
	resultList := analyseNodeMapping(nodeMapping, len(data), len(data[0]), true)
	return len(resultList)
}
