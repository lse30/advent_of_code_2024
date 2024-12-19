package claw_contraption

import (
	"advent_of_code_2024/src/utils"
	"fmt"
	"strconv"
	"strings"
)

// equation Stores all the important info for a given simultaneous equations set
type equation struct {
	xTarget  int
	yTarget  int
	buttonAx int
	buttonAy int
	buttonBx int
	buttonBy int
}

func parseButtonLine(line string) (int, int) {
	data := strings.Split(line, "+")
	xPiece := strings.Split(data[1], ",")[0]
	xDelta, _ := strconv.Atoi(xPiece)
	yDelta, _ := strconv.Atoi(data[2])
	return xDelta, yDelta
}

func parseTargetLine(line string) (int, int) {
	data := strings.Split(line, "=")
	xPiece := strings.Split(data[1], ",")[0]
	xDelta, _ := strconv.Atoi(xPiece)
	yDelta, _ := strconv.Atoi(data[2])
	return xDelta, yDelta
}

// createEquationListV2 reads all equations from the data
func createEquationList(data []string) []equation {
	var output []equation
	for i := 0; i < len(data); i += 4 {
		xTarget, yTarget := parseTargetLine(data[i+2])
		buttonAx, buttonAy := parseButtonLine(data[i])
		buttonBx, buttonBY := parseButtonLine(data[i+1])
		newEquation := equation{
			xTarget:  xTarget,
			yTarget:  yTarget,
			buttonAx: buttonAx,
			buttonAy: buttonAy,
			buttonBx: buttonBx,
			buttonBy: buttonBY,
		}
		output = append(output, newEquation)
	}
	return output
}

// createEquationListV2 reads all equations from the data and changes them for pt2
func createEquationListV2(data []string) []equation {
	var output []equation
	for i := 0; i < len(data); i += 4 {
		xTarget, yTarget := parseTargetLine(data[i+2])
		buttonAx, buttonAy := parseButtonLine(data[i])
		buttonBx, buttonBY := parseButtonLine(data[i+1])
		newEquation := equation{
			xTarget:  xTarget + 10000000000000,
			yTarget:  yTarget + 10000000000000,
			buttonAx: buttonAx,
			buttonAy: buttonAy,
			buttonBx: buttonBx,
			buttonBy: buttonBY,
		}
		output = append(output, newEquation)
	}
	return output
}

// FindLcm finds the least common multiple of two numbers
func FindLcm(a, b int) int {
	var gcd int
	if a >= b {
		gcd = Gcd(a, b)
	} else {
		gcd = Gcd(b, a)
	}
	return (a * b) / gcd
}

// Gcd Greatest common divisor algorithm
func Gcd(a, b int) int {
	// where a >= b
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// solveEquation Simultaneous equation solver for 2 equations, 2 variables.
func solveEquation(eq equation) (float64, float64) {
	lcm := FindLcm(eq.buttonAx, eq.buttonAy)
	eq1Mul := lcm / eq.buttonAx
	eq2Mul := lcm / eq.buttonAy
	var a, b float64
	b = float64(eq.yTarget*eq2Mul-eq.xTarget*eq1Mul) / float64(eq.buttonBy*eq2Mul-eq.buttonBx*eq1Mul)
	a = (float64(eq.xTarget) - (float64(eq.buttonBx) * b)) / float64(eq.buttonAx)
	return a, b
}

// evaluateCost if there is a valid solution then calculate the cost of each button
func evaluateCost(a, b float64) int {
	if a == float64(int64(a)) && a == float64(int64(a)) {
		fmt.Println("Real solution found")
		fmt.Printf("a = %d, b = %d\n", int(a), int(b))
		return (int(a) * 3) + int(b)
	}
	return 0
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	equationSet := createEquationList(data)
	output := 0
	for _, eq := range equationSet {
		a, b := solveEquation(eq)
		output += evaluateCost(a, b)
	}
	return output
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	equationSet := createEquationListV2(data)
	output := 0
	for _, eq := range equationSet {
		a, b := solveEquation(eq)
		output += evaluateCost(a, b)
	}
	return output
}
