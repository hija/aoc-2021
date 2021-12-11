package main

import (
	"fmt"
	"os"
	"strings"
)

func getExpectedBracket(openingbracket rune) rune {
	if openingbracket == '{' {
		return '}'
	} else if openingbracket == '(' {
		return ')'
	} else if openingbracket == '<' {
		return '>'
	} else if openingbracket == '[' {
		return ']'
	}
	return ' '
}

func getPanaltyCost(wrongCharacter rune) int {
	if wrongCharacter == '}' {
		return 1197
	} else if wrongCharacter == ']' {
		return 57
	} else if wrongCharacter == '>' {
		return 25137
	} else if wrongCharacter == ')' {
		return 3
	}
	return -1
}

func solveTaskOne(lines []string) int {
	score := 0

	for _, line := range lines {
		opencharacters := []rune{}
		for _, character := range line {
			if character == '{' || character == '[' || character == '(' || character == '<' {
				opencharacters = append(opencharacters, character)
			} else {
				// Get last opencharacter since this one should be closed now
				lc := opencharacters[len(opencharacters)-1]
				opencharacters = opencharacters[:len(opencharacters)-1]

				if character != getExpectedBracket(lc) {
					score += getPanaltyCost(character)
				}
			}
		}
	}
	return score
}

func main() {
	inputAsBytes, err := os.ReadFile("10/input.txt")

	if err != nil {
		fmt.Println("Error reading input >> Exit")
		os.Exit(-1)
	}

	inputAsStr := string(inputAsBytes)
	lines := strings.Split(inputAsStr, "\n")

	fmt.Print("Task 1: ")
	fmt.Println(solveTaskOne(lines))

}
