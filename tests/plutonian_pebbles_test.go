package tests

import (
	"advent_of_code_2024/src/day_11"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestPlutonianPebblesPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_11/problem_input_1.txt"
	actualOutput := plutonian_pebbles.Solve(filename, 25)
	expectedOutput := 55312
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestPlutonianPebblesPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_11/problem_input_2.txt"
	actualOutput := plutonian_pebbles.Solve(filename, 25)
	expectedOutput := 203953
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestPlutonianPebblesPtTwo(t *testing.T) {
	filename := test_utils.FilePrefix + "day_11/problem_input_2.txt"
	actualOutput := plutonian_pebbles.Solve(filename, 75)
	expectedOutput := 242090118578155
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
