package mull_it_over

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Token struct {
	keyword string
	pattern string
}

var (
	multiplication = Token{"mul", `mul\(\d+,\d+\)`}
	dos            = Token{"do()", `do\(\)`}
	donts          = Token{"don't()", `don't\(\)`}
)

// combinePatterns Combines multiple regex patterns
func combinePatterns(tokens []Token) string {
	var output string
	for _, token := range tokens {
		output += "|" + token.pattern
	}
	return output[1:]
}

// tokeniser Finds all occurrences of a target token
func multiTokeniser(inputString string, tokens []Token) []string {
	combinedPattern := regexp.MustCompile(combinePatterns(tokens))
	// Find all matches
	matches := combinedPattern.FindAllString(inputString, -1)
	return matches
}

// extractNumbers gets the two numbers contained in the brackets
func extractNumbers(match string) (int, int) {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(match, -1)
	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[1])
	return a, b
}

// readFile Opens and reads a file then converts each column to a string
func readFileToString(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	fileStr := string(b)
	return fileStr
}

// SolvePartOne Aim is to calculate the total absolute value of each column once they have been ordered.
func SolvePartOne(fileName string) int {
	fileContents := readFileToString(fileName)
	output := 0
	matches := multiTokeniser(fileContents, []Token{multiplication})
	for _, match := range matches {
		a, b := extractNumbers(match)
		output += a * b
	}
	return output
}

// identifyToken Reads what token the match is
func identifyToken(match string, tokens []Token) string {
	for _, token := range tokens {
		if token.keyword == match[0:len(token.keyword)] {
			return token.keyword
		}
	}
	return "No Token"
}

// solveMatchList reads through the identified tokens and evaluates the output
func solveMatchList(matches []string, tokens []Token) int {
	doFlag := 1
	output := 0
	for _, match := range matches {
		current := identifyToken(match, tokens)
		if current == multiplication.keyword {
			a, b := extractNumbers(match)
			output += a * b * doFlag
		} else if current == dos.keyword {
			doFlag = 1
		} else if current == donts.keyword {
			doFlag = 0
		}
	}
	return output
}

// SolvePartTwo Aim to parse the tokens and evaluate the tokens in the correct order
func SolvePartTwo(fileName string) int {
	fileContents := readFileToString(fileName)
	tokenList := []Token{multiplication, dos, donts}
	matches := multiTokeniser(fileContents, tokenList)
	output := solveMatchList(matches, tokenList)
	return output
}
