package main

import (
	"fmt"
	"os"
	"sort"
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

type VisitedPoint struct {
	x, y int
}

func (p VisitedPoint) containedIn(points []VisitedPoint) bool {
	for _, point := range points {
		if p == point {
			return true
		}
	}
	return false
}

func mergeVisitedPoints(a1 []VisitedPoint, a2 []VisitedPoint) []VisitedPoint {
	mergedPoints := make([]VisitedPoint, len(a1))
	copy(mergedPoints, a1)

	for _, v := range a2 {
		if !v.containedIn(mergedPoints) {
			mergedPoints = append(mergedPoints, v)
		}
	}
	return mergedPoints
}

func getNeighboursNotNineRecursively(datapoints [][]int, visitedPoints []VisitedPoint, x int, y int) []VisitedPoint {
	visitedPoints = append(visitedPoints, VisitedPoint{x: x, y: y})

	// Get left neighbour
	if x > 0 {
		left_neighbour_point := VisitedPoint{x: x - 1, y: y}
		if !left_neighbour_point.containedIn(visitedPoints) {
			left_neighbour := datapoints[y][x-1]
			if left_neighbour != 9 {
				visitedPoints = getNeighboursNotNineRecursively(datapoints, visitedPoints, x-1, y)
			}
		}

	}
	// Get right neighbour
	if x < len(datapoints[0])-1 {
		right_neighbour_point := VisitedPoint{x: x + 1, y: y}
		if !right_neighbour_point.containedIn(visitedPoints) {
			right_neighbour := datapoints[y][x+1]
			if right_neighbour != 9 {
				visitedPoints = getNeighboursNotNineRecursively(datapoints, visitedPoints, x+1, y)
			}
		}
	}

	// Get top neigbour
	if y > 0 {
		top_neighbour_point := VisitedPoint{x: x, y: y - 1}
		if !top_neighbour_point.containedIn(visitedPoints) {
			top_neighbour := datapoints[y-1][x]
			if top_neighbour != 9 {
				visitedPoints = getNeighboursNotNineRecursively(datapoints, visitedPoints, x, y-1)
			}
		}
	}

	// Get bottom neighbour
	if y < len(datapoints)-1 {
		bottom_neighbour_point := VisitedPoint{x: x, y: y + 1}
		if !bottom_neighbour_point.containedIn(visitedPoints) {
			bottom_neighbour := datapoints[y+1][x]
			if bottom_neighbour != 9 {
				visitedPoints = getNeighboursNotNineRecursively(datapoints, visitedPoints, x, y+1)
			}
		}
	}
	return visitedPoints
}

func getBasin(datapoints [][]int, x int, y int) int {
	emptyVisitedPoints := make([]VisitedPoint, 0)
	basin := getNeighboursNotNineRecursively(datapoints, emptyVisitedPoints, x, y)
	return len(basin)
}

func solveTaskTwo(datapoints [][]int) int {
	basins := []int{}
	for y := 0; y < len(datapoints); y++ {
		for x := 0; x < len(datapoints[0]); x++ {
			neigbours := getNeighbours(datapoints, x, y)
			if datapoints[y][x] < smallestIntValue(neigbours) {
				basins = append(basins, getBasin(datapoints, x, y))
			}
		}
	}

	product := 1

	sort.Ints(basins)
	for i := len(basins) - 1; i >= len(basins)-3; i-- {
		product *= basins[i]
	}

	return product
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
	fmt.Print("Task 2: ")
	fmt.Println(solveTaskTwo(datapoints))

}
