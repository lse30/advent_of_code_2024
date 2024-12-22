package tests

import (
	"advent_of_code_2024/src/day_15"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestWarehouseWoesPtOneExample(t *testing.T) {
	filename := test_utils.FilePrefix + "day_15/problem_input_1.txt"
	actualOutput := warehouse_woes.SolvePartOne(filename)
	expectedOutput := 10092
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestWarehouseWoesPtOne(t *testing.T) {
	filename := test_utils.FilePrefix + "day_15/problem_input_2.txt"
	actualOutput := warehouse_woes.SolvePartOne(filename)
	expectedOutput := 1360570
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
