package tests

import (
	"advent_of_code_2024/src/day_7"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestBridgeRepairPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_7/problem_input_1.txt"
	actualOutput := bridge_repair.SolvePartOne(filename)
	expectedOutput := 3749
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestBridgeRepairPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_7/problem_input_2.txt"
	actualOutput := bridge_repair.SolvePartOne(filename)
	expectedOutput := 3598800864292
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestBridgeRepairPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_7/problem_input_1.txt"
	actualOutput := bridge_repair.SolvePartTwo(filename)
	expectedOutput := 11387
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestBridgeRepairPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_7/problem_input_2.txt"
	actualOutput := bridge_repair.SolvePartTwo(filename)
	expectedOutput := 340362529351427
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
