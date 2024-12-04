package tests

import (
	"advent_of_code_2024/src/day_4"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestCeresSearchPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_4/problem_input_1.txt"
	actualOutput := ceres_search.SolvePartOne(filename)
	expectedOutput := 18
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestCeresSearchPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_4/problem_input_2.txt"
	actualOutput := ceres_search.SolvePartOne(filename)
	expectedOutput := 2642
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestCeresSearchPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_4/problem_input_1.txt"
	actualOutput := ceres_search.SolvePartTwo(filename)
	expectedOutput := 9
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestCeresSearchPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_4/problem_input_2.txt"
	actualOutput := ceres_search.SolvePartTwo(filename)
	expectedOutput := 1974
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
