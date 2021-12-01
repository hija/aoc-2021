package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getNumIncreasing(list []int) int {

	prev := 0
	numIncreased := 0

	for i, v := range list {
		if i == 1 {
			continue
		}

		if v > prev {
			numIncreased = numIncreased + 1
		}

		prev = v
	}

	return numIncreased
}

func sum(list []int) int {
	summe := 0
	for _, v := range list {
		summe = summe + v
	}
	return summe
}

func getNumIncreasingOffset(list []int, offset int) int {
	numIncreased := 0

	for i, _ := range list {
		if i < offset+1 {
			continue
		}

		if sum(list[i-(offset+1):i-1]) < sum(list[i-offset:i]) {
			numIncreased = numIncreased + 1
		}

	}

	return numIncreased

}

func main() {
	fmt.Println("Hello World!")

	data, err := ioutil.ReadFile("01/input.txt")

	if err != nil {
		fmt.Println("Error reading input file. Exit!")
		os.Exit(-1)
	}

	input := string(data)
	singleNumbers := strings.Split(input, "\n")

	inputAsInt := make([]int, len(singleNumbers))

	for i, v := range singleNumbers {
		inputAsInt[i], _ = strconv.Atoi(v)
	}

	fmt.Println(getNumIncreasingOffset(inputAsInt, 3))
}
