package tests

import (
	"advent_of_code_2024/src/day_1"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	filename := "src/day_1/problem_input_2.txt"
	actualOutput := historian_hysteria.solvePartOne(filename)
	expectedOutput := 2
	test_utils.AssertEqual(t, actualOutput, expectedOutput)
}
