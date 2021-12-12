package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Cave struct {
	isBigCave   bool
	name        string
	connections []*Cave
}

func (c Cave) containedIn(list []*Cave) bool {
	for _, cave := range list {
		if cave.equals(c) {
			return true
		}
	}
	return false
}

func getCaveByName(str string, caves []*Cave) *Cave {
	for caveindex := range caves {
		if caves[caveindex].name == str {
			return caves[caveindex]
		}
	}
	return &Cave{}
}

func (c1 Cave) equals(c2 Cave) bool {
	return c1.name == c2.name
}

func isUppercaseOnly(str string) bool {
	for _, v := range str {
		if !unicode.IsUpper(v) {
			return false
		}
	}
	return true
}

type Path struct {
	caves []*Cave
}

func (p Path) printPath() {
	for i, v := range p.caves {
		if i != 0 {
			fmt.Print(" --> ")
		}
		fmt.Print(v.name)
	}
	fmt.Println()
}

func countOccurences(caves []*Cave, cave Cave) int {
	count := 0
	for _, v := range caves {
		if v.name == cave.name {
			count++
		}
	}
	return count
}

func (p Path) isValidPath(maxvisits int) bool {
	// Check if a part is valid, i.e. if a small cave has not been visited more than once

	visitedSmallCaves := []*Cave{}
	maxAllowedOccurence := maxvisits

	for _, element := range p.caves {
		if !element.isBigCave {
			count := countOccurences(visitedSmallCaves, *element)
			if count >= maxAllowedOccurence {
				return false
			}
			if count == maxvisits-1 {
				maxAllowedOccurence = 1
			}
			visitedSmallCaves = append(visitedSmallCaves, element)
		}
	}
	return true
}

func findPathToEnd(p Path, maxVisitsSmallCave int) []Path {
	lastElement := p.caves[len(p.caves)-1]
	nextPaths := []Path{}
	for _, nextStep := range lastElement.connections {

		if nextStep.name == "start" {
			continue
		}

		currentPathCaves := make([]*Cave, len(p.caves)+1)
		copy(currentPathCaves, p.caves)
		currentPathCaves[len(currentPathCaves)-1] = nextStep

		newPath := Path{caves: currentPathCaves}

		if !newPath.isValidPath(maxVisitsSmallCave) {
			continue
		}

		if nextStep.name != "end" {
			nextPaths = append(nextPaths, findPathToEnd(newPath, maxVisitsSmallCave)...)
		} else {
			newPath.printPath()
			nextPaths = append(nextPaths, newPath)
		}
	}

	return nextPaths

}

func solveTaskOne(network []*Cave) int {

	startCave := getCaveByName("start", network)
	paths := findPathToEnd(Path{caves: []*Cave{startCave}}, 1)

	return len(paths)
}

func solveTaskTwo(network []*Cave) int {

	startCave := getCaveByName("start", network)
	paths := findPathToEnd(Path{caves: []*Cave{startCave}}, 2)

	return len(paths)
}

func main() {
	dataAsBytes, err := os.ReadFile("12/input.txt")
	if err != nil {
		fmt.Println("Error loading input >> Exit")
		os.Exit(-1)
	}

	dataAsStr := string(dataAsBytes)
	lines := strings.Split(dataAsStr, "\n")

	caves := []*Cave{}

	// Construct the network
	for _, line := range lines {
		strC1 := strings.Split(line, "-")[0]
		strC2 := strings.Split(line, "-")[1]

		c1IsBigCave := isUppercaseOnly(strC1)
		c2IsBigCave := isUppercaseOnly(strC2)

		cave1 := Cave{isBigCave: c1IsBigCave, name: strC1, connections: []*Cave{}}
		cave2 := Cave{isBigCave: c2IsBigCave, name: strC2, connections: []*Cave{}}

		if !cave1.containedIn(caves) {
			caves = append(caves, &cave1)
		}

		if !cave2.containedIn(caves) {
			caves = append(caves, &cave2)
		}

		pCave1 := getCaveByName(strC1, caves)
		pCave2 := getCaveByName(strC2, caves)

		if !pCave1.containedIn(pCave2.connections) {
			pCave2.connections = append(pCave2.connections, pCave1)
		}

		if !pCave2.containedIn(pCave1.connections) {
			pCave1.connections = append(pCave1.connections, pCave2)
		}
	}

	fmt.Println("Task 1:", solveTaskOne(caves))
	fmt.Println("Task 2:", solveTaskTwo(caves))
}
