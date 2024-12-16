package utils

import (
	"fmt"
	"os"
	"strings"
)

// Coord 2-dimensional coordinate type
type Coord struct {
	X int
	Y int
}

func (c Coord) AddCoord(other Coord) Coord {
	return Coord{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

// AbsValue calculates the absolute value
func AbsValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ReadFileToUnits reads data from the file split by rows (\n)
func ReadFileToUnits(fileName string) [][]string {
	fileStr := ReadFileToString(fileName)
	var output [][]string
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

// ReadFileToRows reads the file and splits the data by row
func ReadFileToRows(fileName string) []string {
	fileStr := ReadFileToString(fileName)
	return strings.Split(fileStr, "\n")
}

// ReadFileToString Reads all the file into one string
func ReadFileToString(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	fileStr := string(b)
	return fileStr
}

// InBounds checks if a given set of points is within a grid
func InBounds(x int, y int, xLimit int, yLimit int) bool {
	return x >= 0 && x < xLimit && y >= 0 && y < yLimit
}
