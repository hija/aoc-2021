package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoField struct {
	field_values  [][]int
	fields_checks [][]bool
}

func (bf *BingoField) addFieldValue(value int) {
	for column := 0; column < 5; column++ {
		for row := 0; row < 5; row++ {

			fieldvalue := bf.field_values[row][column]
			if fieldvalue == value {
				bf.fields_checks[row][column] = true
			}
		}
	}
}

func (bf *BingoField) sumUncheckedFields() int {
	sum := 0
	// Then check columns
	for column := 0; column < 5; column++ {

		for row := 0; row < 5; row++ {
			fieldvalue := bf.fields_checks[row][column]
			if !fieldvalue {
				sum += bf.field_values[row][column]
			}
		}
	}
	return sum
}

func (bf *BingoField) printBoard() {
	for _, row := range bf.field_values {
		fmt.Println(row)
	}
}

func (bf *BingoField) hasBingo() bool {
	// First check rows
	for _, row := range bf.fields_checks {
		isBingo := true

		for _, fieldvalue := range row {
			if !fieldvalue {
				isBingo = false
			}
		}
		if isBingo {
			return true
		}
	}

	// Then check columns
	for column := 0; column < 5; column++ {
		isBingo := true

		for row := 0; row < 5; row++ {
			fieldvalue := bf.fields_checks[row][column]
			if !fieldvalue {
				isBingo = false
			}
		}
		if isBingo {
			return true
		}
	}
	return false
}

func lineToBoardLine(line string) []int {
	arr := make([]int, 5)

	for i := 0; i < 5; i = i + 1 {
		arr[i], _ = strconv.Atoi(strings.Trim(line[(i*3):(i*3+2)], " "))
	}
	return arr
}

func strArrToIntArr(arr []string) []int {
	newArr := make([]int, len(arr))
	for i, v := range arr {
		newArr[i], _ = strconv.Atoi(v)
	}
	return newArr
}

func main() {
	dataAsByte, err := os.ReadFile("04/input.txt")
	if err != nil {
		fmt.Println("Error reading file >> Exit")
	}

	dataAsStr := string(dataAsByte)
	lines := strings.Split(dataAsStr, "\n")

	numbersAsStr := lines[0]
	lines = lines[2:]

	var bingoFields []BingoField

	for {
		bingoFieldData := make([][]int, 5)
		bingoFieldBoolData := make([][]bool, 5)

		for i := 0; i < 5; i++ {
			bingoFieldData[i] = lineToBoardLine(lines[i])
			bingoFieldBoolData[i] = []bool{false, false, false, false, false} // ugly but should work
		}

		bingofield := BingoField{field_values: bingoFieldData, fields_checks: bingoFieldBoolData}
		bingoFields = append(bingoFields, bingofield)
		if len(lines) >= 6 {
			lines = lines[6:]
		} else {
			break
		}
	}

	numbersAsInt := strArrToIntArr(strings.Split(numbersAsStr, ","))
	for _, number := range numbersAsInt {
		endLoop := false
		for _, bingofield := range bingoFields {
			bingofield.addFieldValue(number)

			if bingofield.hasBingo() {
				fmt.Print("Task 1: ")
				fmt.Println(bingofield.sumUncheckedFields() * number)
				endLoop = true
			}
		}
		if endLoop {
			break
		}
	}

	// Task 2
	for _, number := range numbersAsInt {
		i := 0

		for {
			if i >= len(bingoFields) {
				break
			}
			bingofield := bingoFields[i]
			bingofield.addFieldValue(number)

			if bingofield.hasBingo() {
				if len(bingoFields) == 1 {
					// last board has bingo
					fmt.Print("Task 2:")
					fmt.Println(bingofield.sumUncheckedFields() * number)
				}

				// remove the bingo field from the list of bingo fields, since it has been solved now
				bingoFields = append(bingoFields[:i], bingoFields[i+1:]...)
				continue
			}

			i += 1
		}
	}
}
