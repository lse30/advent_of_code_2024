package tests

import (
	"advent_of_code_2024/src/day_12"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestGardenGroupsPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_12/problem_input_1.txt"
	actualOutput := garden_groups.SolvePartOne(filename)
	expectedOutput := 1930
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestGardenGroupsPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_12/problem_input_2.txt"
	actualOutput := garden_groups.SolvePartOne(filename)
	expectedOutput := 1396562
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestGardenGroupsPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_12/problem_input_1.txt"
	actualOutput := garden_groups.SolvePartTwo(filename)
	expectedOutput := 1206
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestGardenGroupsPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_12/problem_input_2.txt"
	actualOutput := garden_groups.SolvePartTwo(filename)
	expectedOutput := 844132
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
