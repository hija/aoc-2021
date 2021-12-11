package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lineToIntSlice(line string) []int {
	ints := make([]int, len(line))
	for i, v := range line {
		ints[i], _ = strconv.Atoi(string(v))
	}
	return ints
}

func getNeighbours(datapoints [][]int, x int, y int) []int {
	var neighbourValues []int

	// Get left neighbour
	if x > 0 {
		neighbourValues = append(neighbourValues, datapoints[y][x-1])
	}
	// Get right neighbour
	if x < len(datapoints[0])-1 {
		neighbourValues = append(neighbourValues, datapoints[y][x+1])
	}
	// Get top neigbour
	if y > 0 {
		neighbourValues = append(neighbourValues, datapoints[y-1][x])
	}
	// Get bottom neighbour
	if y < len(datapoints)-1 {
		neighbourValues = append(neighbourValues, datapoints[y+1][x])
	}
	return neighbourValues
}

func smallestIntValue(input []int) int {
	smallestValue := input[0]
	for i := 1; i < len(input); i++ {
		if smallestValue > input[i] {
			smallestValue = input[i]
		}
	}
	return smallestValue
}

func solveTaskOne(datapoints [][]int) int {
	sum := 0
	for y := 0; y < len(datapoints); y++ {
		for x := 0; x < len(datapoints[0]); x++ {
			neigbours := getNeighbours(datapoints, x, y)
			if datapoints[y][x] < smallestIntValue(neigbours) {
				sum += datapoints[y][x] + 1
			}
		}
	}
	return sum
}

func main() {
	dataAsByte, err := os.ReadFile("09/input.txt")
	if err != nil {
		fmt.Println("Error loading input file >> Exit")
		os.Exit(-1)
	}

	dataAsStr := string(dataAsByte)
	lines := strings.Split(dataAsStr, "\n")
	datapoints := make([][]int, len(lines))

	for i, line := range lines {
		datapoints[i] = lineToIntSlice(line)
	}

	fmt.Print("Task 1: ")
	fmt.Println(solveTaskOne(datapoints))
}
