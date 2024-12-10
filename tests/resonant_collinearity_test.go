package tests

import (
	"advent_of_code_2024/src/day_8"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestResonantCollinearityPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_8/problem_input_1.txt"
	actualOutput := resonant_collinearity.SolvePartOne(filename)
	expectedOutput := 14
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestResonantCollinearityPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_8/problem_input_2.txt"
	actualOutput := resonant_collinearity.SolvePartOne(filename)
	expectedOutput := 398
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestResonantCollinearityPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_8/problem_input_1.txt"
	actualOutput := resonant_collinearity.SolvePartTwo(filename)
	expectedOutput := 34
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestResonantCollinearityPtTwoExampleTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_8/problem_input_3.txt"
	actualOutput := resonant_collinearity.SolvePartTwo(filename)
	expectedOutput := 9
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestResonantCollinearityPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_8/problem_input_2.txt"
	actualOutput := resonant_collinearity.SolvePartTwo(filename)
	expectedOutput := 1333
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
