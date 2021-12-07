package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

func strSliceToFloatSlice(slice []string) []float64 {
	newSlice := make([]float64, len(slice))

	for i, v := range slice {
		newSlice[i], _ = strconv.ParseFloat(v, 64)
	}

	return newSlice
}

func sumDistanceToMedian(slice []float64) int {
	median, _ := stats.Median(slice)

	diff := 0
	for _, v := range slice {
		diff += int(math.Abs(median - v))
	}

	return diff
}

func sumDistanceToMean(slice []float64) int {
	mean, _ := stats.Mean(slice)
	meanFloor := math.Floor(mean)
	meanCeil := math.Ceil(mean)

	totaldiffFloor := 0
	for _, v := range slice {
		diff := int(math.Abs(meanFloor - v))
		totaldiffFloor += int((diff * (diff + 1)) / 2)
	}

	totaldiffCeil := 0
	for _, v := range slice {
		diff := int(math.Abs(meanCeil - v))
		totaldiffCeil += int((diff * (diff + 1)) / 2)
	}

	if totaldiffCeil < totaldiffFloor {
		return totaldiffCeil
	} else {
		return totaldiffFloor
	}
}

func main() {
	dataAsByte, err := os.ReadFile("07/input.txt")
	if err != nil {
		fmt.Println("Could not read input >> Exit")
		os.Exit(-1)
	}

	dataAsStr := string(dataAsByte)
	numbersAsStrSlice := strings.Split(dataAsStr, ",")
	numbersAsFloatSlice := strSliceToFloatSlice(numbersAsStrSlice)

	// Task 1
	// Idea is easy: The median lies in "the middle" of all values
	// thus we calculate the difference between all values and their median
	fmt.Print("Task 1: ")
	fmt.Println(sumDistanceToMedian(numbersAsFloatSlice))

	// Task 2
	// The fuel burned is equal to Gauss-Sum (diff*diff+1) / 2
	// To keep the overall sum small it makes sense to minimize the distance between n and x
	// Just use the mean?
	fmt.Print("Task 2: ")
	fmt.Println(sumDistanceToMean(numbersAsFloatSlice))

}
