package tests

import (
	"advent_of_code_2024/src/day_5"
	"advent_of_code_2024/tests/test_utils"
	"testing"
)

func TestPrintQueuePtOneExample(t *testing.T) {
	fileName := test_utils.FilePrefix + "day_5/problem_input_1.txt"
	actualOutput := print_queue.PrintQueue(fileName, false)
	expectedOutput := 143
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestPrintQueuePtOne(t *testing.T) {
	fileName := test_utils.FilePrefix + "day_5/problem_input_2.txt"
	actualOutput := print_queue.PrintQueue(fileName, false)
	expectedOutput := 7074
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestPrintQueuePtTwoExample(t *testing.T) {
	fileName := test_utils.FilePrefix + "day_5/problem_input_1.txt"
	actualOutput := print_queue.PrintQueue(fileName, true)
	expectedOutput := 123
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}

func TestPrintQueuePtTwo(t *testing.T) {
	fileName := test_utils.FilePrefix + "day_5/problem_input_2.txt"
	actualOutput := print_queue.PrintQueue(fileName, true)
	expectedOutput := 4828
	test_utils.AssertEqual(t, expectedOutput, actualOutput)
}
