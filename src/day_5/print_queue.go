package main

import (
	"advent_of_code_2024/src/utils"
	"fmt"
)

type chainLink struct {
	value int
}

func main() {
	fileName := "src/day_5/problem_input_1.txt"
	data := utils.ReadFileToRows(fileName)
	for _, line := range data {
		fmt.Println(line)
	}
}
