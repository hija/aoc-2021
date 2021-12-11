package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

func getRuneByCount(signalPatterns []string, relevantCount int) rune {
	if relevantCount == 7 || relevantCount == 8 {
		return ' '
	}
	countMap := map[rune]int{}
	for _, signalPattern := range signalPatterns {
		for _, signalRune := range signalPattern {
			countMap[signalRune] = countMap[signalRune] + 1
		}
	}

	for k, v := range countMap {
		if v == relevantCount {
			return k
		}
	}

	return ' '
}

func getNumberWithMapping(mapping map[rune]rune, input string) string {
	validSignals := []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}
	decodedSignal := ""
	for _, rune := range input {
		decodedSignal += string(mapping[rune])
	}

	decodedSignal = SortString(decodedSignal)

	for i, v := range validSignals {
		if v == decodedSignal {
			number := i
			return strconv.Itoa(number)
		}
	}
	return ""
}

func mapHasValue(mapOfConcern map[rune]rune, value rune) bool {
	for k, _ := range mapOfConcern {
		if k == value {
			return true
		}
	}
	return false
}

func getMapFromSignalPatterns(signalPatterns []string) map[rune]rune {
	returnMap := make(map[rune]rune)
	returnMap[getRuneByCount(signalPatterns, 9)] = 'f'
	returnMap[getRuneByCount(signalPatterns, 6)] = 'b'
	returnMap[getRuneByCount(signalPatterns, 4)] = 'e'

	// Get 1
	for _, v := range signalPatterns {
		if len(v) == 2 {
			for _, singleRune := range v {
				if !mapHasValue(returnMap, singleRune) {
					returnMap[singleRune] = 'c'
				}
			}
		}
	}

	// Get 4
	for _, v := range signalPatterns {
		if len(v) == 4 {
			for _, singleRune := range v {
				if !mapHasValue(returnMap, singleRune) {
					returnMap[singleRune] = 'd'
				}
			}
		}
	}

	// Get 7
	for _, v := range signalPatterns {
		if len(v) == 3 {
			for _, singleRune := range v {
				if !mapHasValue(returnMap, singleRune) {
					returnMap[singleRune] = 'a'
				}
			}
		}
	}

	// Get last value
	for _, v := range signalPatterns {
		if len(v) == 7 {
			for _, singleRune := range v {
				if !mapHasValue(returnMap, singleRune) {
					returnMap[singleRune] = 'g'
				}
			}
		}
	}
	return returnMap
}

func SortString(w string) string {
	// Taken from https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func solveLine(line string) int {
	signalPatterns := strings.Split(strings.Split(line, " | ")[0], " ")
	numbersToRetrieve := strings.Split(strings.Split(line, " | ")[1], " ")

	finalNumberAsStr := ""

	mapping := getMapFromSignalPatterns(signalPatterns)
	for _, numberToRetrieve := range numbersToRetrieve {
		decodedNumber := getNumberWithMapping(mapping, numberToRetrieve)
		finalNumberAsStr = finalNumberAsStr + decodedNumber
	}

	finalNumber, _ := strconv.Atoi(finalNumberAsStr)
	return finalNumber
}

func solveTwo(lines []string) int {
	result := 0

	for _, line := range lines {
		result += solveLine(line)
	}
	return result
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

	fmt.Print("Task 2: ")
	fmt.Println(solveTwo(lines))
}
