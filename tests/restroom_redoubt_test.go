package tests

import (
	"advent_of_code_2024/src/day_14"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestRestroomRedoubtPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_14/problem_input_1.txt"
	actualOutput := restroom_redoubt.SolvePartOne(filename, 7, 11, 100)
	expectedOutput := 12
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestRestroomRedoubtPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_14/problem_input_2.txt"
	actualOutput := restroom_redoubt.SolvePartOne(filename, 103, 101, 100)
	expectedOutput := 228410028
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestRestroomRedoubtPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_14/problem_input_2.txt"
	actualOutput := restroom_redoubt.SolvePartTwo(filename, 103, 101)
	expectedOutput := 8258
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestRestroomRedoubtPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_14/problem_input_3.txt"
	actualOutput := restroom_redoubt.SolvePartTwo(filename, 103, 101)
	expectedOutput := 7083
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
