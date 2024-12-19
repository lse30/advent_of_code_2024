package garden_groups

import (
	"advent_of_code_2024/src/utils"
	"fmt"
)

type region struct {
	identifier string
	coords     []utils.Coord
}

type node struct {
	coords  utils.Coord
	nodeMap [][]int
}

func (c *node) initMapping() {
	c.nodeMap = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
}

func (c *node) printMapping() {
	for _, line := range c.nodeMap {
		fmt.Println(line)
	}
}

func buildRegion(data [][]string, coord utils.Coord, currentRegion region) region {
	if currentRegion.identifier != data[coord.X][coord.Y] {
		return currentRegion
	}
	currentRegion.coords = append(currentRegion.coords, coord)
	data[coord.X][coord.Y] = "."
	var directions = []utils.Coord{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}
	for _, direction := range directions {
		newCoord := direction.AddCoord(coord)
		if utils.InBounds(newCoord.X, newCoord.Y, len(data), len(data[0])) {
			currentRegion = buildRegion(data, newCoord, currentRegion)
		}
	}
	return currentRegion
}

func createRegions(data [][]string) []region {
	var output []region
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if string(data[i][j]) != "." {
				coord := utils.Coord{X: i, Y: j}
				initRegion := region{identifier: data[i][j], coords: []utils.Coord{}}
				newRegion := buildRegion(data, coord, initRegion)
				output = append(output, newRegion)
			}
		}
	}
	return output
}

func checkAdjacent(coord1 utils.Coord, coord2 utils.Coord) bool {
	xDelta := utils.AbsValue(coord1.X - coord2.X)
	yDelta := utils.AbsValue(coord1.Y - coord2.Y)
	if (xDelta == 1 && yDelta == 0) || (xDelta == 0 && yDelta == 1) {
		return true
	}
	return false
}

func evaluateRegion(regionMap region) int {
	area := len(regionMap.coords)
	var parsedData []utils.Coord
	perimeter := 0
	for i := 0; i < area; i++ {
		perimeter += 4
		for j := 0; j < len(parsedData); j++ {
			if checkAdjacent(regionMap.coords[i], parsedData[j]) {
				perimeter -= 2
			}
		}
		parsedData = append(parsedData, regionMap.coords[i])
	}
	//fmt.Printf("Region %s has a value of %d * %d = %d\n", regionMap.identifier, area, perimeter, area*perimeter)
	return area * perimeter
}

func evaluateMap(targetNode node, newCoord utils.Coord) {
	deltaX := newCoord.X - targetNode.coords.X
	deltaY := newCoord.Y - targetNode.coords.Y
	if utils.AbsValue(deltaX) <= 1 && utils.AbsValue(deltaY) <= 1 {
		targetNode.nodeMap[deltaX+1][deltaY+1] = 1
	}
}

func scoreblock(a int, b int, c int) int {
	if a == 0 {
		if b == 0 {
			if c == 0 {
				return 3
			} else {
				return 0
			}
		} else {
			if c == 0 {
				return 0
			} else {
				return 1
			}
		}
	} else {
		if b == 0 {
			if c == 0 {
				return 3
			} else {
				return 1
			}
		} else {
			if c == 0 {
				return 1
			} else {
				return 0
			}
		}
	}
}

func countCorners(a node) int {
	outputA := scoreblock(a.nodeMap[0][0], a.nodeMap[0][1], a.nodeMap[1][0])
	outputB := scoreblock(a.nodeMap[0][2], a.nodeMap[1][2], a.nodeMap[0][1])
	outputC := scoreblock(a.nodeMap[2][2], a.nodeMap[2][1], a.nodeMap[1][2])
	outputD := scoreblock(a.nodeMap[2][0], a.nodeMap[1][0], a.nodeMap[2][1])
	output := outputA + outputB + outputC + outputD
	//fmt.Printf("Map Score %d\n", output)
	//a.printMapping()
	return output
}

func evaluateRegionV2(regionMap region) int {
	area := len(regionMap.coords)
	sides := 0
	for i := 0; i < len(regionMap.coords); i++ {
		newNode := node{coords: regionMap.coords[i]}
		newNode.initMapping()
		for j := 0; j < len(regionMap.coords); j++ {
			evaluateMap(newNode, regionMap.coords[j])
		}
		corners := countCorners(newNode)
		sides += corners
	}
	sides /= 3
	fmt.Printf("Region %s has a value of %d * %d = %d\n", regionMap.identifier, area, sides, area*sides)
	return area * sides
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToUnits(fileName)
	regionsList := createRegions(data)
	output := 0
	for _, item := range regionsList {
		regionValue := evaluateRegion(item)
		output += regionValue
	}
	return output
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToUnits(fileName)
	regionsList := createRegions(data)
	output := 0
	for _, item := range regionsList {
		regionValue := evaluateRegionV2(item)
		output += regionValue
	}
	return output
}
