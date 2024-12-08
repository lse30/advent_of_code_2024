package tests

import (
	guardgallivant "advent_of_code_2024/src/day_6"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestGuardGallivantPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_6/problem_input_1.txt"
	actualOutput := guardgallivant.SolvePartOne(filename)
	expectedOutput := 41
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestGuardGallivantPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_6/problem_input_2.txt"
	actualOutput := guardgallivant.SolvePartOne(filename)
	expectedOutput := 4758
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestGuardGallivantPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_6/problem_input_1.txt"
	actualOutput := guardgallivant.SolvePartTwo(filename)
	expectedOutput := 6
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestGuardGallivantPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_6/problem_input_2.txt"
	actualOutput := guardgallivant.SolvePartTwo(filename)
	expectedOutput := 1670
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
