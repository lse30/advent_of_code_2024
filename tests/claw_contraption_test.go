package tests

import (
	claw_contraption "advent_of_code_2024/src/day_13"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestClawContraptionPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_13/problem_input_1.txt"
	actualOutput := claw_contraption.SolvePartOne(filename)
	expectedOutput := 480
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestClawContraptionPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_13/problem_input_2.txt"
	actualOutput := claw_contraption.SolvePartOne(filename)
	expectedOutput := 32067
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestClawContraptionPtTwoExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_13/problem_input_1.txt"
	actualOutput := claw_contraption.SolvePartTwo(filename)
	expectedOutput := 875318608908
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestClawContraptionPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_13/problem_input_2.txt"
	actualOutput := claw_contraption.SolvePartTwo(filename)
	expectedOutput := 92871736253789
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestGcd(t *testing.T) {
	a := 1071
	b := 462
	actualOutput := claw_contraption.Gcd(a, b)
	expectedOutput := 21
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestFindLcm(t *testing.T) {
	a := 21
	b := 6
	actualOutput := claw_contraption.FindLcm(a, b)
	expectedOutput := 42
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
