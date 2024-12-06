package print_queue

import (
	"advent_of_code_2024/src/utils"
	"strconv"
	"strings"
)

// ChainLink Defines the structure of the target, the numbers it must be ahead and behind of.
type ChainLink struct {
	Target      int
	MustPrecede []int
	MustFollow  []int
}

// checkInList returns the index if the target in within the list else -1
func checkInList(chains []ChainLink, target int) int {
	for i := 0; i < len(chains); i++ {
		if chains[i].Target == target {
			return i
		}
	}
	return -1
}

// getAB extracts two numbers from the line of format "X|Y"
func getAB(line string) (int, int) {
	items := strings.Split(line, "|")
	a, _ := strconv.Atoi(items[0])
	b, _ := strconv.Atoi(items[1])
	return a, b
}

// buildOrder Creates a new order from the file line
func buildOrder(line string) []int {
	var output []int
	items := strings.Split(line, ",")
	for _, item := range items {
		value, _ := strconv.Atoi(item)
		output = append(output, value)
	}
	return output
}

// contains if an int is contained inside []int
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// checkIsValid Checks that a given index is valid
func checkIsValid(order []int, rule ChainLink, index int) bool {
	// is there a number in front of target that is not allowed to be there?
	for _, x := range order[:index] {
		if contains(rule.MustPrecede, x) {
			return false
		}
	}

	// is there a number behind target that must go before it?
	for _, x := range order[index:] {
		if contains(rule.MustFollow, x) {
			return false
		}
	}
	return true
}

// CheckOrder Asserts if a given order meets the criteria of the ruleset
func CheckOrder(order []int, rules []ChainLink) bool {
	for i := 0; i < len(order); i++ {
		// Get the rule for this index (if exists)
		index := checkInList(rules, order[i])
		if index != -1 {
			isValid := checkIsValid(order, rules[index], i)
			if !isValid {
				return false
			}
		}
	}
	return true
}

// analyseOrders Builds and checks all given orders, adds the middle value of correct orders.
func analyseOrders(data []string, rules []ChainLink) int {
	var output int
	for _, line := range data {
		if strings.Contains(line, ",") {
			order := buildOrder(line)
			if CheckOrder(order, rules) {
				toAdd := order[(len(order))/2]
				output += toAdd
			}
		}
	}
	return output
}

// insertIntoList adds a new value into a list at a given index
func insertIntoList(list []int, target int, index int) []int {
	var output []int
	for i := 0; i < len(list); i++ {
		if i == index {
			output = append(output, target)
		}
		output = append(output, list[i])
	}
	return output
}

// placeNextNumber determines where to place the next number in the sequence
func placeNextNumber(current []int, rule ChainLink) []int {
	for i := 0; i < len(current); i++ {
		if contains(rule.MustPrecede, current[i]) {
			// This number must be ahead of this current number, insert it now.
			return insertIntoList(current, rule.Target, i)
		}
	}
	current = append(current, rule.Target)
	return current
}

// reorderOrder builds the order in the correct order
func reorderOrder(order []int, rules []ChainLink) []int {
	output := []int{order[0]}
	for i := 1; i < len(order); i++ {
		output = placeNextNumber(output, rules[checkInList(rules, order[i])])
	}
	return output
}

// repairOrders fixes all broken orders and returns the sum of all middle values of all broken orders after correction
func repairOrders(data []string, rules []ChainLink) int {
	var output int
	for _, line := range data {
		if strings.Contains(line, ",") {
			order := buildOrder(line)
			if !CheckOrder(order, rules) {
				// fix the order
				fixedOrder := reorderOrder(order, rules)
				toAdd := fixedOrder[(len(order))/2]
				output += toAdd
			}
		}
	}
	return output
}

// buildRuleset creates a slice containing data for all the number rules
func buildRuleset(data []string) []ChainLink {
	var output []ChainLink
	for _, line := range data {
		if strings.Contains(line, "|") {
			a, b := getAB(line)
			index := checkInList(output, a)
			if index == -1 {
				output = append(output, ChainLink{Target: a, MustPrecede: []int{b}})
			} else {
				output[index].MustPrecede = append(output[index].MustPrecede, b)
			}

			index2 := checkInList(output, b)
			if index2 == -1 {
				output = append(output, ChainLink{Target: b, MustFollow: []int{a}})
			} else {
				output[index2].MustFollow = append(output[index2].MustFollow, a)
			}
		}
	}
	return output
}

// PrintQueue script entrance function
func PrintQueue(fileName string, ptTwo bool) int {
	//fileName := "src/day_5/problem_input_1.txt"
	data := utils.ReadFileToRows(fileName)
	rules := buildRuleset(data)
	var result int
	if ptTwo {
		result = repairOrders(data, rules)
	} else {
		result = analyseOrders(data, rules)
	}
	return result
}
