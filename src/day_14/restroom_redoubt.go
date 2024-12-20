package restroom_redoubt

import (
	"advent_of_code_2024/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type robotData struct {
	pos         utils.Coord
	velocityVec utils.Coord
}

func buildPlan(height int, width int) [][]int {
	plan := make([][]int, height)
	for i := range plan {
		plan[i] = make([]int, width)
	}
	return plan
}

func readRobotLine(robot string) (utils.Coord, utils.Coord) {
	data := strings.Split(robot, "=")
	posX, _ := strconv.Atoi(strings.Split(data[1], ",")[0])
	posY, _ := strconv.Atoi(strings.Split(strings.Split(data[1], ",")[1], " ")[0])

	velX, _ := strconv.Atoi(strings.Split(data[2], ",")[0])
	velY, _ := strconv.Atoi(strings.Split(data[2], ",")[1])

	return utils.Coord{X: posX, Y: posY}, utils.Coord{X: velX, Y: velY}
}

func parseRobotData(data []string) []robotData {
	var output []robotData
	for _, robot := range data {
		pos, vel := readRobotLine(robot)
		output = append(output, robotData{pos: pos, velocityVec: vel})
	}
	return output
}

func movePosition(start, speed, width, moves int) int {
	if speed < 0 {
		speed += width
	}
	return (start + moves*speed) % width
}

func fastForward(plan [][]int, data []robotData, moves int) [][]int {
	for _, robot := range data {
		endX := movePosition(robot.pos.X, robot.velocityVec.X, len(plan[0]), moves)
		endY := movePosition(robot.pos.Y, robot.velocityVec.Y, len(plan), moves)
		plan[endY][endX]++
	}
	return plan
}

func printMap(plan [][]int) {
	for _, line := range plan {
		printLine := ""
		for i := 0; i < len(line); i++ {
			if line[i] == 0 {
				printLine += "."
			} else {
				printLine += strconv.Itoa(line[i])
			}
		}
		fmt.Println(printLine)
	}
}

func solveQuadrant(plan [][]int, iLower, iUpper, jLower, jUpper int) int {
	output := 0
	//fmt.Printf("Finding robots in space %d-%d and %d-%d\n", iLower, iUpper, jLower, jUpper)
	for i := iLower; i < iUpper; i++ {
		for j := jLower; j < jUpper; j++ {
			if plan[i][j] != 0 {
				output += plan[i][j]
			}
		}
	}
	return output
}

func solveSafetyFactor(plan [][]int) int {
	height := len(plan)
	width := len(plan[0])
	quadA := solveQuadrant(plan, 0, height/2, 0, width/2)
	quadB := solveQuadrant(plan, 0, height/2, (width/2)+1, width)
	quadC := solveQuadrant(plan, (height/2)+1, height, 0, width/2)
	quadD := solveQuadrant(plan, (height/2)+1, height, (width/2)+1, width)
	return quadA * quadB * quadC * quadD
}

func SolvePartOne(fileName string, h int, w int, moves int) int {
	data := utils.ReadFileToRows(fileName)
	robotData := parseRobotData(data)
	floorPlan := buildPlan(h, w)
	endMap := fastForward(floorPlan, robotData, moves)
	printMap(endMap)
	output := solveSafetyFactor(endMap)
	return output
}

func SolvePartTwo(fileName string, h int, w int) int {
	data := utils.ReadFileToRows(fileName)
	robotData := parseRobotData(data)
	minSafety := -1
	outputMoves := 0
	for move := 0; move <= 10001; move++ {
		floorPlan := buildPlan(h, w)
		endMap := fastForward(floorPlan, robotData, move)
		output := solveSafetyFactor(endMap)
		if minSafety == -1 || output < minSafety {
			minSafety = output
			outputMoves = move
			printMap(endMap)
		}
	}
	return outputMoves
}
