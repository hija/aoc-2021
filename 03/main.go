package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func bitArrayToInt(arr []int) int {
	intVal := 0
	for i, _ := range arr {
		positionOfConcern := len(arr) - 1 - i
		intVal = intVal + (arr[positionOfConcern])*int(math.Pow(2, float64(i)))
	}
	return intVal
}

func reverseArray(arr []int) []int {
	reversedArray := make([]int, len(arr))
	for i, v := range arr {
		if v == 0 {
			reversedArray[i] = 1
		} else {
			reversedArray[i] = 0
		}
	}
	return reversedArray
}

func filterArray(arr []string, position int, value string) []string {
	index := 0
	for {
		if index >= len(arr) {
			break
		}
		currentelement := arr[index]
		if string(currentelement[position]) != value {
			arr = append(arr[:index], arr[index+1:]...)
			continue
		}
		index = index + 1
	}
	return arr
}

func getFrequentValue(arr []string, position int, tie string) string {
	trues := 0
	falses := 0
	for _, v := range arr {
		if string(v[position]) == "1" {
			trues = trues + 1
		} else {
			falses = falses + 1
		}
	}

	if trues == falses {
		return tie
	}

	if trues > falses {
		return "1"
	}

	return "0"
}

func strArrayToIntArray(arr string) []int {
	reArr := make([]int, len(arr))
	for i, v := range arr {
		reArr[i], _ = strconv.Atoi(string(v))
	}
	return reArr
}

func main() {
	inputAsBytes, err := os.ReadFile("03/input.txt")
	if err != nil {
		fmt.Println("Error loading input file >> Exit")
		os.Exit(-1)
	}

	inputAsStr := string(inputAsBytes)
	lines := strings.Split(inputAsStr, "\n")

	mostCommonPosition := make([][]int, 12)
	for i := range mostCommonPosition {
		mostCommonPosition[i] = make([]int, 2)
	}

	for _, v := range lines {
		for i, p := range v {
			if string(p) == "1" {
				mostCommonPosition[i][1] = mostCommonPosition[i][1] + 1
				//sumPositions[i] = sumPositions[i] + 1
			} else {
				mostCommonPosition[i][0] = mostCommonPosition[i][0] + 1
			}
		}
	}

	sumPositions := make([]int, len(mostCommonPosition))

	for i, _ := range mostCommonPosition {
		if mostCommonPosition[i][0] > mostCommonPosition[i][1] {
			sumPositions[i] = 0
		} else {
			sumPositions[i] = 1
		}
	}

	for i, _ := range sumPositions {
		sumPositions[i] = int(math.Mod(float64(sumPositions[i]), 2))
	}

	fmt.Print("Task 1:")
	fmt.Println(bitArrayToInt(sumPositions) * bitArrayToInt(reverseArray(sumPositions)))

	// Teil 2
	cpylist := make([]string, len(lines))
	copy(cpylist, lines)
	oxygen := 0
	index := 0

	for {
		if len(cpylist) == 1 {
			oxygen = bitArrayToInt(strArrayToIntArray(cpylist[0]))
			break
		}
		mostFrequent := getFrequentValue(cpylist, index, "1")
		cpylist = filterArray(cpylist, index, mostFrequent)
		index = index + 1
	}

	cpylist = make([]string, len(lines))
	copy(cpylist, lines)
	co2 := 0
	index = 0

	for {
		if len(cpylist) == 1 {
			co2 = bitArrayToInt(strArrayToIntArray(cpylist[0]))
			break
		}
		mostFrequent := getFrequentValue(cpylist, index, "tie")
		if mostFrequent == "1" || mostFrequent == "tie" {
			cpylist = filterArray(cpylist, index, "0")
		} else {
			cpylist = filterArray(cpylist, index, "1")
		}
		index = index + 1
	}

	fmt.Print("Task 2:")
	fmt.Println(oxygen * co2)

}
