package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vent struct {
	x1, x2, y1, y2 int
}

func (v *Vent) overlapsWithPoint(x int, y int, checkdiagnoal bool) bool {

	xOverlap := false
	yOverlap := false
	diagonal := false

	if v.x1 < v.x2 {
		xOverlap = x >= v.x1 && x <= v.x2 && v.y1 == v.y2 && v.y1 == y
		diagonal = diagonal || (x >= v.x1 && x <= v.x2 && v.isOnDiagonal(x, y))
	} else {
		xOverlap = x >= v.x2 && x <= v.x1 && v.y1 == v.y2 && v.y1 == y
		diagonal = diagonal || (x >= v.x2 && x <= v.x1 && v.isOnDiagonal(x, y))
	}

	if v.y1 < v.y2 {
		yOverlap = y >= v.y1 && y <= v.y2 && v.x1 == v.x2 && v.x1 == x
		diagonal = diagonal || (y >= v.y1 && y <= v.y2 && v.isOnDiagonal(x, y))
	} else {
		yOverlap = y >= v.y2 && y <= v.y1 && v.x1 == v.x2 && v.x1 == x
		diagonal = diagonal || (y >= v.y2 && y <= v.y1 && v.isOnDiagonal(x, y))
	}

	return xOverlap || yOverlap || (checkdiagnoal && diagonal)
}

func (v *Vent) isOnDiagonal(x int, y int) bool {
	// Idea: Create a line between both points (== diagnonal line)
	// and check if the point lies on that line
	m_upper := (v.y2 - v.y1)
	m_divisor := (v.x2 - v.x1)
	m := 0

	if m_divisor == 0 {
		return false
	}

	m = m_upper / m_divisor

	if m == 0 {
		return false
	}
	b := v.y1 - m*v.x1

	return m*x+b == y
}

func parseInputLine(line string) (x1 int, y1 int, x2 int, y2 int) {
	lineParts := strings.Split(line, " -> ")
	firstPart := lineParts[0]
	secondPart := lineParts[1]

	x1, _ = strconv.Atoi(strings.Split(firstPart, ",")[0])
	y1, _ = strconv.Atoi(strings.Split(firstPart, ",")[1])

	x2, _ = strconv.Atoi(strings.Split(secondPart, ",")[0])
	y2, _ = strconv.Atoi(strings.Split(secondPart, ",")[1])

	return x1, y1, x2, y2
}

func main() {
	inputAsBytes, err := os.ReadFile("05/input.txt")
	if err != nil {
		fmt.Println("Error loading file >> Exit")
		os.Exit(-1)
	}

	inputAsStr := string(inputAsBytes)
	lines := strings.Split(inputAsStr, "\n")

	var vents []Vent
	xmax, ymax := 0, 0

	for _, line := range lines {
		// Parse inputline
		x1, y1, x2, y2 := parseInputLine(line)
		vent := Vent{x1: x1, x2: x2, y1: y1, y2: y2}
		vents = append(vents, vent)

		// Set xmax and ymax
		if xmax < x2 {
			xmax = x2
		} else if xmax < x1 {
			xmax = x1
		}
		if ymax < y2 {
			ymax = y2
		} else if ymax < y1 {
			ymax = y1
		}
	}

	overlapCounter := 0
	// Check overlapping points
	for yi := 0; yi <= ymax; yi++ {
		//fmt.Println()
		for xi := 0; xi <= xmax; xi++ {
			pointsInPosition := 0
			for _, vent := range vents {
				if vent.overlapsWithPoint(xi, yi, false) {
					pointsInPosition++
				}
			}

			if pointsInPosition >= 2 {
				overlapCounter++
				//fmt.Print(pointsInPosition)
			} else {
				//fmt.Print(".")
			}
		}
	}
	fmt.Print("Task 1: ")
	fmt.Println(overlapCounter)

	overlapCounter = 0
	// Check overlapping points
	for yi := 0; yi <= ymax; yi++ {
		//fmt.Println()
		for xi := 0; xi <= xmax; xi++ {
			pointsInPosition := 0
			for _, vent := range vents {
				if vent.overlapsWithPoint(xi, yi, true) {
					pointsInPosition++
				}
			}

			if pointsInPosition >= 2 {
				overlapCounter++
				//fmt.Print(pointsInPosition)
			} else {
				//fmt.Print(".")
			}
		}
	}
	fmt.Print("Task 2: ")
	fmt.Println(overlapCounter)

}
