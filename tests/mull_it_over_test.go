package tests

import (
	"advent_of_code_2024/src/day_3"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestMullItOverPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_3/problem_input_1.txt"
	actualOutput := mull_it_over.SolvePartOne(filename)
	expectedOutput := 161
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestMullItOverPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_3/problem_input_2.txt"
	actualOutput := mull_it_over.SolvePartOne(filename)
	expectedOutput := 175015740
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestMullItOverPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_3/problem_input_3.txt"
	actualOutput := mull_it_over.SolvePartTwo(filename)
	expectedOutput := 48
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestMullItOverPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_3/problem_input_2.txt"
	actualOutput := mull_it_over.SolvePartTwo(filename)
	expectedOutput := 112272912
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
