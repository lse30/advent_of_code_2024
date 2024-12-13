package disk_fragmenter

import (
	"advent_of_code_2024/src/utils"
	"fmt"
	"strconv"
)

type memoryBlock struct {
	value  string
	length int
}

func isEven(a int) bool {
	return a%2 == 0
}

func buildMemoryString(data string) []string {
	var output []string
	id := 0
	for i := 0; i < len(data); i++ {
		var idString string
		lengthValue, _ := strconv.Atoi(string(data[i]))
		if isEven(i) {
			idString = strconv.Itoa(id)
			id++
		} else {
			idString = "."
		}
		for j := 0; j < lengthValue; j++ {
			output = append(output, idString)
		}
	}
	return output
}

func buildMemoryBlock(data string) []memoryBlock {
	id := 0
	var output []memoryBlock
	for i := 0; i < len(data); i++ {
		dataValue, _ := strconv.Atoi(string(data[i]))
		if isEven(i) {
			idString := strconv.Itoa(id)
			newBlock := memoryBlock{value: idString, length: dataValue}
			output = append(output, newBlock)
			id++
		} else {
			newBlock := memoryBlock{value: ".", length: dataValue}
			output = append(output, newBlock)
		}
	}
	return output
}

func findValidGap(memory []memoryBlock, gapLength int) int {
	for i := 0; i < len(memory); i++ {
		if memory[i].value == "." {

		}
	}
	return -1
}

func compressMemory(memory []memoryBlock) []memoryBlock {
	var output []memoryBlock

	for k := len(memory) - 1; k >= 0; k-- {
		fmt.Println(memory[k].length)
		if memory[k].value == "." {
			continue
		} else {
			gapIndex := findValidGap(memory[:k], memory[k].length)

		}
	}

	return output
}

func compressMemoryV2(memory []string) []string {
	i := 0
	k := len(memory) - 1
	for i <= k {
		if memory[i] != "." {
			i++
		} else {
			// time to swap with k
			for memory[k] == "." {
				k--
			}
			memory[i] = memory[k]
			memory[k] = "."
			k--
			i++
		}
	}
	return memory
}

func displayMemory(memory []memoryBlock) {
	for _, block := range memory {
		for i := 0; i < block.length; i++ {
			fmt.Print(block.value)
			fmt.Print(",")
		}
	}
	fmt.Println()
}

func displayMemoryString(memory []string) {
	for _, block := range memory {
		fmt.Print(block)
		fmt.Print(",")
	}
	fmt.Println()
}

func calculateCheckSumString(memory []string) int {
	output := 0
	for i := 0; i < len(memory); i++ {
		if memory[i] == "." {
			return output
		} else {
			value, _ := strconv.Atoi(memory[i])
			output += i * value
		}
	}
	return output
}

func calculateCheckSum(memory []memoryBlock) int {
	index := 0
	output := 0
	for _, block := range memory {
		value, _ := strconv.Atoi(block.value)
		for i := 0; i < block.length; i++ {
			output += value * index
			index++
		}
	}
	return output
}

func SolvePartOne(fileName string) int {
	data := utils.ReadFileToString(fileName)
	memory := buildMemoryString(data)
	memory = compressMemoryV2(memory)
	memory = compressMemoryV2(memory)
	output := calculateCheckSumString(memory)
	return output
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToString(fileName)
	output := buildMemoryBlock(data)
	displayMemory(output)
	compressMemory(output)
	return 1
}
