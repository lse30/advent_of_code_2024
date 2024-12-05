package utils

import (
	"fmt"
	"os"
	"strings"
)

// AbsValue calculates the absolute value
func AbsValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadFileToString(fileName string) [][]string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	var output [][]string
	fileStr := string(b)
	var line []string
	for i := 0; i < len(fileStr); i++ {
		if fileStr[i] != '\n' {
			line = append(line, string(fileStr[i]))
		} else {
			output = append(output, line)
			line = []string{}
		}

	}
	output = append(output, line)
	return output
}

func ReadFileToRows(fileName string) []string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	fileStr := string(b)
	output := strings.Split(fileStr, "\n")
	return output
}
