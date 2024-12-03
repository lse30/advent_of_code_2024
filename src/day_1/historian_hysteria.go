package historian_hysteria

import (
	"advent_of_code_2024/src/utils"
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// SolvePartOne Aim is to calculate the total absolute value of each column once they have been ordered.
func SolvePartOne(fileName string) int {
	listA, listB := readFile(fileName)
	sort.Ints(listA)
	sort.Ints(listB)
	output := 0
	for i := 0; i < len(listA); i++ {
		distance := utils.AbsValue(listA[i] - listB[i])
		output += distance
	}
	return output
}

// SolvePartTwo Aim is to count the occurrences of the items on the right list for each value on the left.
// then sum the totals for every item
func SolvePartTwo(fileName string) int {
	leftList, rightList := readFile(fileName)
	output := 0
	for i := 0; i < len(leftList); i++ {
		occurrences := countListOccurrences(leftList[i], rightList)
		output += leftList[i] * occurrences
	}
	return output
}

// countListOccurrences Counts the number of occurrences of a given number within a list
func countListOccurrences(target int, items []int) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if items[i] == target {
			total++
		}
	}
	return total
}

// readFile Opens and reads a file then converts each column to an integer list
func readFile(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			numbers = append(numbers, num)
		}
	}
	return splitList(numbers)
}

// splitList Splits a giant list into two sub lists (unmerges the columns)
func splitList(original []int) ([]int, []int) {
	var outputA []int
	var outputB []int
	for i := 0; i < len(original); i += 2 {
		outputA = append(outputA, original[i])
		outputB = append(outputB, original[i+1])
	}
	return outputA, outputB
}
