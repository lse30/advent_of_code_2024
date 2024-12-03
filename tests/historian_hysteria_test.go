package tests

import (
	"advent_of_code_2024/src/day_1"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestSolvePartOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_1/problem_input_1.txt"
	actualOutput := historian_hysteria.SolvePartOne(filename)
	expectedOutput := 11
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestSolvePartTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_1/problem_input_1.txt"
	actualOutput := historian_hysteria.SolvePartTwo(filename)
	expectedOutput := 31
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestSolvePartOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_1/problem_input_2.txt"
	actualOutput := historian_hysteria.SolvePartOne(filename)
	expectedOutput := 1873376
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestSolvePartTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_1/problem_input_2.txt"
	actualOutput := historian_hysteria.SolvePartTwo(filename)
	expectedOutput := 18997088
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
