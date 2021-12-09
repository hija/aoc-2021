package main

import (
	"fmt"
	"os"
	"strings"
)

func solveOne(lines []string) int {
	result := 0
	for _, line := range lines {
		secondPart := strings.Split(line, "|")[1]
		allParts := strings.Split(secondPart, " ")
		for _, part := range allParts {
			if len(part) == 4 || len(part) == 2 || len(part) == 3 || len(part) == 7 {
				result++
			}
		}
	}
	return result
}

func getMapFromSignalPatterns(signalPatterns []string) map[int]string {
	returnMap := make(map[int]string)
	for _, signalPattern := range signalPatterns {
		returnMap[len(signalPattern)] = signalPattern
	}
	return returnMap
}

func determineNumber(signalPatterns []string, numberToRetrieve string) string {

	return ""
}

func solveLine(line string) int {
	signalPatterns := strings.Split(strings.Split(line, "|")[0], " ")
	numbersToRetrieve := strings.Split(strings.Split(line, "|")[1], " ")

	finalNumber := ""

	for _, numberToRetrieve := range numbersToRetrieve {
		if len(numberToRetrieve) == 4 {
			finalNumber += "4"
		} else if len(numberToRetrieve) == 2 {
			finalNumber += "1"
		} else if len(numberToRetrieve) == 3 {
			finalNumber += "7"
		} else if len(numberToRetrieve) == 7 {
			finalNumber += "8"
		} else {
			finalNumber += determineNumber(signalPatterns, numberToRetrieve)
		}

	}
}

func solveTwo(lines []string) int {
	result := 0

	for _, line := range lines {

	}
}

func main() {
	dataAsByte, err := os.ReadFile("08/input.txt")
	if err != nil {
		fmt.Println("Error reading input file >> Exit")
		os.Exit(-1)
	}

	dataAsStr := string(dataAsByte)
	lines := strings.Split(dataAsStr, "\n")

	fmt.Print("Task 1: ")
	fmt.Println(solveOne(lines))
}
