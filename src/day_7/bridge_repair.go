package bridge_repair

import (
	"advent_of_code_2024/src/utils"
	"strconv"
	"strings"
)

func listAdd(listA, listB []int) []int {
	output := make([]int, len(listA)+len(listB))
	copy(output, listA)
	copy(output[len(listA):], listB)
	return output
}

func solveLine(target int, values []int) bool {
	if len(values) == 1 {
		return target == values[0]
	}
	newMulti := []int{values[0] * values[1]}
	newAdd := []int{values[0] + values[1]}
	rem := values[2:]
	multiList := listAdd(newMulti, rem)
	addList := listAdd(newAdd, rem)
	if solveLine(target, multiList) || solveLine(target, addList) {
		return true
	}
	return false
}

func concat(a int, b int) int {
	outputStr := strconv.Itoa(a) + strconv.Itoa(b)
	outputInt, _ := strconv.Atoi(outputStr)
	return outputInt
}

func solveLineAdvanced(target int, values []int) bool {
	if len(values) == 1 {
		return target == values[0]
	}
	newMulti := []int{values[0] * values[1]}
	newAdd := []int{values[0] + values[1]}
	newConcat := []int{concat(values[0], values[1])}
	rem := values[2:]
	multiList := listAdd(newMulti, rem)
	addList := listAdd(newAdd, rem)
	concatList := listAdd(newConcat, rem)
	if solveLineAdvanced(target, multiList) || solveLineAdvanced(target, addList) || solveLineAdvanced(target, concatList) {
		return true
	}
	return false
}

func parseLine(fileLine string) (int, []int) {
	items := strings.Split(fileLine, ":")
	target, _ := strconv.Atoi(items[0])
	var output []int
	for _, item := range strings.Split(items[1], " ") {
		number, _ := strconv.Atoi(item)
		if number != 0 {
			output = append(output, number)
		}
	}
	return target, output
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	var output int
	for i := 0; i < len(data); i++ {
		target, values := parseLine(data[i])
		if solveLine(target, values) {
			output += target
		}
	}

	return output
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToRows(fileName)
	var output int
	for i := 0; i < len(data); i++ {
		target, values := parseLine(data[i])
		if solveLineAdvanced(target, values) {
			output += target
		}
	}

	return output
}
