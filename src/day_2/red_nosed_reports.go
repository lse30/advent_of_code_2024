package red_nosed_reports

import (
	"advent_of_code_2024/src/utils"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// SolvePartOne Analyses the file containing the reports to determine which are true and false
func SolvePartOne(fileName string) int {
	reportsList := readFile(fileName)
	safeReports := 0
	for _, report := range reportsList {
		reportIsSafe := analyseReport(report)
		if reportIsSafe {
			safeReports++
		}
	}
	return safeReports
}

// SolvePartTwo Analyses the file containing the reports to determine which are true and false
func SolvePartTwo(fileName string) int {
	reportsList := readFile(fileName)
	safeReports := 0
	for _, report := range reportsList {
		reportIsSafe := ReportProblemDampener(report)
		if reportIsSafe {
			safeReports++
		}
	}
	return safeReports
}

// getGradient finds if the report is Flat, Descending or Ascending
func getGradient(a int, b int) string {
	if a == b {
		return "Flat"
	}
	if a > b {
		return "Descending"
	}
	return "Ascending"
}

// analyseReport Checks if given report is valid
func analyseReport(report []int) bool {
	// We can assume the report contains at least two numbers
	gradient := getGradient(report[0], report[1])
	if gradient == "Flat" {
		return false
	}
	current := report[0]
	for i := 1; i < len(report); i++ {
		deltaGradient := getGradient(current, report[i])
		if gradient != deltaGradient {
			return false
		}
		deltaValue := utils.AbsValue(current - report[i])
		if deltaValue > 3 || deltaValue < 1 {
			return false
		}
		current = report[i]
	}
	return true
}

// generateSubList creates a new list with a given index ignored
func generateSubList(original []int, ignoreIndex int) []int {
	var output []int
	for i := 0; i < len(original); i++ {
		if i != ignoreIndex {
			output = append(output, original[i])
		}
	}
	return output
}

// ReportProblemDampener allows for one index to be ignored while analysing a report
func ReportProblemDampener(report []int) bool {
	if len(report) < 2 {
		return true
	}
	for i := 0; i < len(report); i++ {
		newReport := generateSubList(report, i)
		if analyseReport(newReport) {
			return true
		}
	}
	return false
}

// readFile Opens and reads a file then converts rows to reports
func readFile(fileName string) [][]int {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	scanner := bufio.NewScanner(file)

	var numbers []int
	var output [][]int
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			numbers = append(numbers, num)
		}
		output = append(output, numbers)
		numbers = []int{}
	}
	return output
}
