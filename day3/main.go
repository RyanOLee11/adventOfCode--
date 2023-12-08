package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("schematic.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	correctNumbers := []int{}
	gearRatios := []int{}

	matrix := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		line := scanner.Text()

		// Initialize a new row slice
		row := []string{}

		for _, c := range line {
			// Append the character to the row slice
			row = append(row, string(c))
		}

		// Append the row slice to the matrix
		matrix = append(matrix, row)
	}
	currentNumber := ""
	symbolFound := false
	symbol := ""
	for rowIndex, row := range matrix {
		for colIndex, col := range row {
			//fmt.Print(col)
			if unicode.IsDigit([]rune(col)[0]) {
				currentNumber += string(col)
				if !symbolFound {
					symbolFound, symbol = checkAroundChar(matrix, rowIndex, colIndex)
				}
			} else {
				if currentNumber != "" {
					number, err := strconv.Atoi(currentNumber)
					if err != nil {
						log.Fatalf("Failed converting string to integer: %s", err)
					}
					if symbolFound {
						correctNumbers = append(correctNumbers, number)
						if len(currentNumber) == 2 {
							gearRatios = append(gearRatios, number)
						}
					}
					currentNumber = ""
					symbolFound = false
				}

			}

			if col == "*" {
				checkAroundStar(matrix, rowIndex, colIndex)
			}
		}
	}

	//fmt.Println(correctNumbers)
	sum := 0
	for _, number := range correctNumbers {
		sum += number
	}
	fmt.Sprintf(symbol)
	fmt.Println(sum)
	gearSum := 0
	for _, number := range gearRatios {
		gearSum += number
	}
	fmt.Println(gearSum)

}

func checkAroundChar(pMatrix [][]string, pRow int, pCol int) (bool, string) {
	// Check the row above

	if pRow > 0 {
		if checkForSpecialChar(pMatrix[pRow-1][pCol]) {
			return true, pMatrix[pRow-1][pCol]
		}
		if (pCol-1) >= 0 && checkForSpecialChar(pMatrix[pRow-1][pCol-1]) {
			return true, pMatrix[pRow-1][pCol-1]
		}
		if (pCol+1) <= len(pMatrix[pRow])-1 && checkForSpecialChar(pMatrix[pRow-1][pCol+1]) {
			return true, pMatrix[pRow-1][pCol+1]
		}
	}

	// Check the row below
	if pRow < len(pMatrix)-1 {
		if checkForSpecialChar(pMatrix[pRow+1][pCol]) {
			return true, pMatrix[pRow+1][pCol]
		}
		if (pCol-1) >= 0 && checkForSpecialChar(pMatrix[pRow+1][pCol-1]) {
			return true, pMatrix[pRow+1][pCol-1]
		}
		if (pCol+1) <= len(pMatrix[pRow])-1 && checkForSpecialChar(pMatrix[pRow+1][pCol+1]) {
			return true, pMatrix[pRow+1][pCol+1]
		}
	}

	// Check the column to the left
	if pCol > 0 {
		if checkForSpecialChar(pMatrix[pRow][pCol-1]) {
			return true, pMatrix[pRow][pCol-1]
		}
	}

	// Check the column to the right
	if pCol < len(pMatrix[pRow])-1 {
		if checkForSpecialChar(pMatrix[pRow][pCol+1]) {
			return true, pMatrix[pRow][pCol+1]
		}
	}

	return false, ""
}

func checkForSpecialChar(s string) bool {
	b := (unicode.IsLetter([]rune(s)[0]) || unicode.IsNumber([]rune(s)[0]) || s == ".")
	if !b {
		fmt.Println(s)
	}
	return !b

}
