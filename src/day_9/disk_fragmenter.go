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

func popValidEndElement(slice []memoryBlock) ([]memoryBlock, memoryBlock) {
	outputBlock := slice[len(slice)-1]
	for outputBlock.value == "." {
		slice = slice[:len(slice)-1]
		outputBlock = slice[len(slice)-1]
	}
	return slice[:len(slice)-1], outputBlock
}

func compressMemory(memory []memoryBlock) []memoryBlock {
	var output []memoryBlock
	var endValue memoryBlock

	for len(memory) > 0 {
		if memory[0].value != "." {
			output = append(output, memory[0])
		} else {
			lengthToFill := memory[0].length
			memory, endValue = popValidEndElement(memory)
			for lengthToFill > 0 {
				if endValue.length <= lengthToFill {
					output = append(output, endValue)
					lengthToFill -= endValue.length
					if lengthToFill > 0 {
						memory, endValue = popValidEndElement(memory)
					}
				} else {
					output = append(output, memoryBlock{value: endValue.value, length: lengthToFill})
					memory = append(memory, memoryBlock{value: endValue.value, length: endValue.length - lengthToFill})
					lengthToFill = 0
				}
			}
		}
		memory = memory[1:]
	}
	return output
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
	memory := buildMemoryBlock(data)
	displayMemory(memory)
	memory = compressMemory(memory)
	displayMemory(memory)
	output := calculateCheckSum(memory)
	return output
}

func SolvePartTwo(fileName string) int {
	data := utils.ReadFileToString(fileName)
	output := buildMemoryBlock(data)
	displayMemory(output)
	compressMemory(output)
	return 1
}
