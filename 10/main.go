package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

func getCompletionCost(expectedCharacter rune) int {
	if expectedCharacter == '}' {
		return 3
	} else if expectedCharacter == ']' {
		return 2
	} else if expectedCharacter == '>' {
		return 4
	} else if expectedCharacter == ')' {
		return 1
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

func solveTaskTwo(lines []string) int {
	scores := []int{}

	for _, line := range lines {
		score := 0
		opencharacters := []rune{}

		evaluateOpenCharacters := true
		for _, character := range line {
			if character == '{' || character == '[' || character == '(' || character == '<' {
				opencharacters = append(opencharacters, character)
			} else {
				lc := opencharacters[len(opencharacters)-1]
				opencharacters = opencharacters[:len(opencharacters)-1]

				if character != getExpectedBracket(lc) {
					evaluateOpenCharacters = false
					break // We don't care, we want the incomplete lines
				}
			}
		}

		for evaluateOpenCharacters && len(opencharacters) > 0 {
			// incomplete line
			lc := opencharacters[len(opencharacters)-1]
			opencharacters = opencharacters[:len(opencharacters)-1]
			expectedCharacter := getExpectedBracket(lc)
			score = score*5 + getCompletionCost(expectedCharacter)
		}
		if evaluateOpenCharacters {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	medianElement := int(math.Floor(float64(len(scores) / 2)))

	return scores[medianElement]
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
	fmt.Print("Task 2: ")
	fmt.Println(solveTaskTwo(lines))

}
