package warehouse_woes

import (
	"advent_of_code_2024/src/utils"
	"fmt"
	"strings"
)

func initWarehouse(data []string) ([][]string, utils.Coord, string) {
	i := 0
	var robotPos utils.Coord
	var output [][]string
	for data[i] != "" {
		var slice []string
		for j := 0; j < len(data[i]); j++ {
			slice = append(slice, string(data[i][j]))
		}
		output = append(output, slice)
		if strings.Contains(data[i], "@") {
			robotPos = utils.Coord{X: i, Y: strings.Index(data[i], "@")}
		}
		i++
	}
	instructions := ""
	for _, line := range data[i+1:] {
		instructions += line
	}
	return output, robotPos, instructions

}

func canWalk(slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "." {
			return true
		} else if slice[i] == "#" {
			return false
		}
	}
	return false
}

func walkRobot(path []string) ([]string, bool) {
	robotMoved := false

	output := make([]string, len(path))
	copy(output, path)

	if canWalk(output) {
		robotMoved = true
		previous := "."
		for i := 1; i < len(output); i++ {
			next := output[i-1]
			output[i-1] = previous
			previous = next
			if previous == "." {
				return output, robotMoved
			}
		}
	}
	return output, robotMoved
}

func traverseMap(mapping [][]string, robot utils.Coord, instructions string) [][]string {
	for i := 0; i < len(instructions); i++ {
		var robotLine []string
		var walkedRobot []string
		var moved bool
		//fmt.Println(string(instructions[i]))
		if string(instructions[i]) == "^" {
			for j := robot.X; j >= 0; j-- {
				robotLine = append(robotLine, mapping[j][robot.Y])
			}
			walkedRobot, moved = walkRobot(robotLine)
			for j := robot.X; j >= 0; j-- {
				mapping[j][robot.Y] = walkedRobot[robot.X-j]
			}
			if moved {
				robot = utils.Coord{X: robot.X - 1, Y: robot.Y}
			}
		} else if string(instructions[i]) == "v" {
			for j := robot.X; j < len(mapping); j++ {
				robotLine = append(robotLine, mapping[j][robot.Y])
			}
			walkedRobot, moved = walkRobot(robotLine)
			for j := robot.X; j < len(mapping); j++ {
				mapping[j][robot.Y] = walkedRobot[j-robot.X]
			}
			if moved {
				robot = utils.Coord{X: robot.X + 1, Y: robot.Y}
			}
		} else if string(instructions[i]) == "<" {
			for j := robot.Y; j >= 0; j-- {
				robotLine = append(robotLine, mapping[robot.X][j])
			}
			walkedRobot, moved = walkRobot(robotLine)
			for j := robot.Y; j >= 0; j-- {
				mapping[robot.X][j] = walkedRobot[robot.Y-j]
			}
			if moved {
				robot = utils.Coord{X: robot.X, Y: robot.Y - 1}
			}
		} else {
			for j := robot.Y; j < len(mapping[0]); j++ {
				robotLine = append(robotLine, mapping[robot.X][j])
			}
			walkedRobot, moved = walkRobot(robotLine)
			for j := robot.Y; j < len(mapping[0]); j++ {
				mapping[robot.X][j] = walkedRobot[j-robot.Y]
			}
			if moved {
				robot = utils.Coord{X: robot.X, Y: robot.Y + 1}
			}
		}
		//fmt.Println(robotLine)
		//fmt.Println(walkedRobot)
	}
	return mapping
}

func calculateMap(mapping [][]string) int {
	output := 0
	for i := 0; i < len(mapping); i++ {
		for j := 0; j < len(mapping[i]); j++ {
			fmt.Print(mapping[i][j])
			if mapping[i][j] == "O" {
				output += i*100 + j
			}
		}
		fmt.Print("\n")
	}
	return output
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	mapping, robot, instructions := initWarehouse(data)
	fmt.Println(mapping, robot, instructions)
	finalMapping := traverseMap(mapping, robot, instructions)
	return calculateMap(finalMapping)
}
