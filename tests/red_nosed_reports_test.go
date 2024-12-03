package tests

import (
	"advent_of_code_2024/src/day_2"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestRedNosedReportsPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_2/problem_input_1.txt"
	actualOutput := red_nosed_reports.SolvePartOne(filename)
	expectedOutput := 2
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestRedNosedReportsPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_2/problem_input_2.txt"
	actualOutput := red_nosed_reports.SolvePartOne(filename)
	expectedOutput := 390
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestRedNosedReportsPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_2/problem_input_1.txt"
	actualOutput := red_nosed_reports.SolvePartTwo(filename)
	expectedOutput := 4
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestRedNosedReportsPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_2/problem_input_2.txt"
	actualOutput := red_nosed_reports.SolvePartTwo(filename)
	expectedOutput := 439
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
