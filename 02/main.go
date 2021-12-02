package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFinalPosition(commandos []string) (int, int) {
	horizontal := 0
	depth := 0

	for _, v := range commandos {
		if len(v) == 0 {
			continue
		}
		commandparts := strings.Split(v, " ")
		direction := commandparts[0]
		intensity, _ := strconv.Atoi(commandparts[1])

		if direction == "forward" {
			horizontal = horizontal + intensity
		} else if direction == "up" {
			depth = depth - intensity
		} else if direction == "down" {
			depth = depth + intensity
		}
	}

	return horizontal, depth
}

func getFinalPositionWithAim(commandos []string) (int, int) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, v := range commandos {
		if len(v) == 0 {
			continue
		}
		commandparts := strings.Split(v, " ")
		direction := commandparts[0]
		intensity, _ := strconv.Atoi(commandparts[1])

		if direction == "forward" {
			horizontal = horizontal + intensity
			depth = depth + (aim * intensity)
		} else if direction == "up" {
			aim = aim - intensity
		} else if direction == "down" {
			aim = aim + intensity
		}
	}

	return horizontal, depth
}

func main() {

	dataAsBytes, err := os.ReadFile("02/input.txt")
	if err != nil {
		fmt.Println("Error loading input file --> Exit!")
		os.Exit(-1)
	}

	dataAsStr := string(dataAsBytes)
	dataParts := strings.Split(dataAsStr, "\n")

	horizontal, depth := getFinalPosition(dataParts)
	fmt.Print("Solution for part 1 (without aim): ")
	fmt.Println(horizontal * depth)

	horizontal, depth = getFinalPositionWithAim(dataParts)
	fmt.Print("Solution for part 2 (with aim): ")
	fmt.Println(horizontal * depth)

	return
}
