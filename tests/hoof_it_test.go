package tests

import (
	"advent_of_code_2024/src/day_10"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestHoofItPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_10/problem_input_1.txt"
	actualOutput := hoof_it.SolvePartOne(filename)
	expectedOutput := 36
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestHoofItPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_10/problem_input_2.txt"
	actualOutput := hoof_it.SolvePartOne(filename)
	expectedOutput := 659
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestHoofItPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_10/problem_input_1.txt"
	actualOutput := hoof_it.SolvePartTwo(filename)
	expectedOutput := 81
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestHoofItPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_10/problem_input_2.txt"
	actualOutput := hoof_it.SolvePartTwo(filename)
	expectedOutput := 659
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
