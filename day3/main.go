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
					}
					currentNumber = ""
					symbolFound = false
				}

			}

			if col == "*" {
				l_temp :=  checkAroundStar(matrix, rowIndex, colIndex)
				if len(l_temp) == 2 {
					gearRatios = append(gearRatios, l_temp[0] * l_temp[1])
					//fmt.Println(l_temp[0] , l_temp[1])
				}
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
		//fmt.Println(s)
	}
	return !b
}

func checkAroundStar(pMatrix [][]string, pRow int, pCol int) []int {
	numbersAroundStar := []int{}

	if pRow > 0 {
		if unicode.IsNumber([]rune(pMatrix[pRow-1][pCol])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow-1, pCol))
		}
		if (pCol-1) >= 0 && unicode.IsNumber([]rune(pMatrix[pRow-1][pCol-1])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow-1, pCol-1))
		}
		if (pCol+1) <= len(pMatrix[pRow])-1 && unicode.IsNumber([]rune(pMatrix[pRow-1][pCol+1])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow-1, pCol+1))
		}
	}

	// Check the row below
	if pRow < len(pMatrix)-1 {
		if unicode.IsNumber([]rune(pMatrix[pRow+1][pCol])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow+1, pCol))
		}
		if (pCol-1) >= 0 && unicode.IsNumber([]rune(pMatrix[pRow+1][pCol-1])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow+1, pCol-1))
		}
		if (pCol+1) <= len(pMatrix[pRow])-1 && unicode.IsNumber([]rune(pMatrix[pRow+1][pCol+1])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow+1, pCol+1))
		}
	}

	// Check the column to the left
	if pCol > 0 {
		if unicode.IsNumber([]rune(pMatrix[pRow][pCol-1])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow, pCol-1))
		}
	}

	// Check the column to the right
	if pCol < len(pMatrix[pRow])-1 {
		if unicode.IsNumber([]rune(pMatrix[pRow][pCol+1])[0]) {
			numbersAroundStar = appendIfNotExists(numbersAroundStar, getNumberFromMatrix(pMatrix, pRow, pCol+1))
		}
	}
	fmt.Println(numbersAroundStar)
	return numbersAroundStar
}

func getNumberFromMatrix(pMatrix [][]string, pRow int, pCol int) int {
	currentNumber := "" + string(pMatrix[pRow][pCol])
	if pCol < len(pMatrix[pRow])-1 {
		for i := pCol + 1; i <= len(pMatrix[pRow]) -1; i++ {
			if unicode.IsNumber([]rune(pMatrix[pRow][i])[0]) {
				currentNumber += string(pMatrix[pRow][i])
			} else {
				break
			}
		}
	}
	if pCol > 0 {
		for i := pCol - 1; i >= 0; i-- {
			if unicode.IsNumber([]rune(pMatrix[pRow][i])[0]) {
				currentNumber = string(pMatrix[pRow][i]) + currentNumber
			} else {
				break
			}
		}
	}

	number, _ := strconv.Atoi(currentNumber)
	return number
}

func appendIfNotExists(slice []int, i int) []int {
    for _, ele := range slice {
        if ele == i {
            return slice
        }
    }
    return append(slice, i)
}