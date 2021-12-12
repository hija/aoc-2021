package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Neighbour struct {
	x, y int
}

func (n Neighbour) isValidNeighbour() bool {
	if n.x < 0 || n.x > 9 {
		return false
	}
	if n.y < 0 || n.y > 9 {
		return false
	}
	return true
}

func flashOctupus(gameField [][]int, flashingField [][]bool, x int, y int) ([][]int, [][]bool) {
	flashingField[y][x] = true

	top := Neighbour{y: y - 1, x: x}
	topleft := Neighbour{y: y - 1, x: x - 1}
	topright := Neighbour{y: y - 1, x: x + 1}

	left := Neighbour{y: y, x: x - 1}
	right := Neighbour{y: y, x: x + 1}

	bottom := Neighbour{y: y + 1, x: x}
	bottomleft := Neighbour{y: y + 1, x: x - 1}
	bottomright := Neighbour{y: y + 1, x: x + 1}

	neighbours := []Neighbour{top, topleft, topright, left, right, bottom, bottomleft, bottomright}

	for _, neighbour := range neighbours {
		if !neighbour.isValidNeighbour() {
			continue
		}

		// Increase its value
		gameField[neighbour.y][neighbour.x]++
		if gameField[neighbour.y][neighbour.x] > 9 && !flashingField[neighbour.y][neighbour.x] {
			gameField, flashingField = flashOctupus(gameField, flashingField, neighbour.x, neighbour.y)
		}

	}

	return gameField, flashingField
}

func tickGamefield(gameField [][]int) (newGameField [][]int, flashes int) {
	// One gametick

	flashes = 0

	flashingField := make([][]bool, 10)

	// First, the energy level of each octopus increases by 1.
	for y := 0; y < 10; y++ {
		flashingField[y] = make([]bool, 10) // Initialize
		for x := 0; x < 10; x++ {
			gameField[y][x]++
		}
	}

	// Then, any octopus with an energy level greater than 9 flashes.
	// This increases the energy level of all adjacent octopuses by 1,
	// including octopuses that are diagonally adjacent.
	// If this causes an octopus to have an energy level greater than 9, it also flashes.
	// This process continues as long as new octopuses keep having their energy level increased beyond 9.

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if gameField[y][x] > 9 && !flashingField[y][x] {
				gameField, flashingField = flashOctupus(gameField, flashingField, x, y)
			}
		}
	}

	// Finally, any octopus that flashed during this step has its energy level set to 0,
	// as it used all of its energy to flash.
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if flashingField[y][x] {
				gameField[y][x] = 0
				flashes++
			}
		}
	}

	return gameField, flashes
}

func main() {
	inputAsBytes, err := os.ReadFile("11/input.txt")
	if err != nil {
		fmt.Println("Error reading file >> Exit")
		os.Exit(-1)
	}

	inputAsStr := string(inputAsBytes)
	lines := strings.Split(inputAsStr, "\n")
	gameField := make([][]int, 10)

	for y, line := range lines {
		gameField[y] = make([]int, 10)
		for x, element := range line {
			gameField[y][x], _ = strconv.Atoi(string(element))
		}
	}

	flashCounter := 0
	firstTickAllFlashed := -1
	for tick := 0; tick < 100; tick++ {
		newGameField, flashes := tickGamefield(gameField)
		flashCounter += flashes
		if flashes == 100 && firstTickAllFlashed == -1 {
			firstTickAllFlashed = tick + 1 // + 1 bc answer should not be 0 based
		}
		gameField = newGameField
	}

	// Task 2: If we are still not done with all flashing, continue until we do
	tick := 100
	for firstTickAllFlashed == -1 {
		newGameField, flashes := tickGamefield(gameField)
		if flashes == 100 && firstTickAllFlashed == -1 {
			firstTickAllFlashed = tick + 1 // + 1 bc answer should not be 0 based
		}
		tick++
		gameField = newGameField
	}

	fmt.Print("Task 1: ")
	fmt.Println(flashCounter)
	fmt.Print("Task 2: ")
	fmt.Println(firstTickAllFlashed)
}
