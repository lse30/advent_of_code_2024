package plutonian_pebbles

import (
	"advent_of_code_2024/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type stones struct {
	value int
	count int
}

func performBlinkRule(number int) []int {
	if number == 0 {
		return []int{1}
	}
	strNum := strconv.Itoa(number)
	if len(strNum)%2 == 0 {
		numA, _ := strconv.Atoi(strNum[:len(strNum)/2])
		numB, _ := strconv.Atoi(strNum[len(strNum)/2:])
		return []int{numA, numB}
	}
	return []int{number * 2024}
}

func bunchStones(inputA []stones) []stones {
	// Create a map to hold aggregated counts
	grouped := make(map[int]int)

	// Aggregate counts for each value
	for _, pair := range inputA {
		grouped[pair.value] += pair.count
	}

	// Convert the map back to a slice of tuples
	result := make([]stones, 0, len(grouped))
	for val, cnt := range grouped {
		result = append(result, struct {
			value, count int
		}{value: val, count: cnt})
	}
	return result
}

func blink(numbers []stones) []stones {
	var output []stones
	for _, pebble := range numbers {
		newPebbles := performBlinkRule(pebble.value)
		for i := 0; i < len(newPebbles); i++ {
			data := stones{newPebbles[i], pebble.count}
			output = append(output, data)
		}
	}
	output = bunchStones(output)
	return output
}

func countStones(stones []stones) int {
	var output int
	for _, item := range stones {
		output += item.count
	}
	return output
}

func Solve(fileName string, blinks int) int {
	data := strings.Split(utils.ReadFileToString(fileName), " ")
	var pebbles []stones
	for i := 0; i < len(data); i++ {
		num, _ := strconv.Atoi(data[i])
		pebbles = append(pebbles, stones{value: num, count: 1})
	}
	for j := 0; j < blinks; j++ {
		fmt.Println(j)
		pebbles = blink(pebbles)
	}
	return countStones(pebbles)
}
