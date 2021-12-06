package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringSliceTointSlice(arr []string) []int {
	intArr := make([]int, len(arr))
	for i, v := range arr {
		intArr[i], _ = strconv.Atoi(v)
	}
	return intArr
}

func simulateDay(arr []int) []int {
	newDay := make([]int, len(arr))

	for i, v := range arr {
		if v == 0 {
			newDay[i] = 6
			newDay = append(newDay, 8)
		} else {
			newDay[i] = arr[i] - 1
		}
	}

	return newDay
}

func simulateDayWithMap(organismsAsMap map[int]int) map[int]int {
	newDayMap := make(map[int]int)

	for k, v := range organismsAsMap {

		if k > 0 && k < 7 {
			// Easy case - just use previous value
			newDayMap[k-1] = v
		} else if k == 0 {
			// If k == 0 we need to create v new items (--> they get lifetime of 8)
			// Additionally, we need to set the value of 6 manually, since the "parents" now have a lifespan of 6
			newDayMap[8] = v
			newDayMap[6] = organismsAsMap[7] + v
		} else if k == 8 {
			// Remaining case
			newDayMap[7] = v
		}

		// This is an edge case. If we dont have and parents which get children now, k == 0 is not called above
		// and as a result, newDayMap[6] is not filled. So we need to fill it here manually
		if k == 7 && organismsAsMap[0] == 0 {
			newDayMap[6] = organismsAsMap[7]
		}
	}

	return newDayMap
}

func intSliceToMap(inputArr []int) map[int]int {
	resultMap := make(map[int]int)

	for _, v := range inputArr {
		if val, ok := resultMap[v]; ok {
			resultMap[v] = val + 1
		} else {
			resultMap[v] = 1
		}
	}

	return resultMap
}

func main() {

	DAYS_TO_SIMULATE_TASK1 := 80
	DAYS_TO_SIMULATE_TASK2 := 256

	dataAsBytes, err := os.ReadFile("06/input.txt")
	if err != nil {
		fmt.Println("Error reading file >> Exit")
		os.Exit(-1)
	}

	dataAsStr := string(dataAsBytes)
	numbersAsStrArray := strings.Split(dataAsStr, ",")
	numbersAsIntArray := stringSliceTointSlice(numbersAsStrArray)

	for i := 0; i < DAYS_TO_SIMULATE_TASK1; i++ {
		numbersAsIntArray = simulateDay(numbersAsIntArray)
	}

	fmt.Print("Task 1: ")
	fmt.Println(len(numbersAsIntArray))

	// Task 2
	// Using more muscle memory :D
	// Idea: Use a dictionary instead of a list
	// Idea 2 (too lazy right now): Use formula description to calculate

	numbersAsIntArray = stringSliceTointSlice(numbersAsStrArray)
	countMap := intSliceToMap(numbersAsIntArray)

	for i := 0; i < DAYS_TO_SIMULATE_TASK2; i++ {
		countMap = simulateDayWithMap(countMap)
	}

	sumOfOrganisms := 0
	for _, v := range countMap {
		sumOfOrganisms += v
	}
	fmt.Print("Task 2: ")
	fmt.Println(sumOfOrganisms)
}
