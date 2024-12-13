package tests

import (
	"advent_of_code_2024/src/day_9"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestDiskFragmenterPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_9/problem_input_1.txt"
	actualOutput := disk_fragmenter.SolvePartOne(filename)
	expectedOutput := 1928
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestDiskFragmenterPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_9/problem_input_2.txt"
	actualOutput := disk_fragmenter.SolvePartOne(filename)
	expectedOutput := 6366665108136
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestDiskFragmenterPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_9/problem_input_1.txt"
	actualOutput := disk_fragmenter.SolvePartTwo(filename)
	expectedOutput := 34
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestDiskFragmenterPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_9/problem_input_2.txt"
	actualOutput := disk_fragmenter.SolvePartTwo(filename)
	expectedOutput := 1333
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
