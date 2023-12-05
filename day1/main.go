package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fileName := "numbers.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// we will use this slice to store the numbers
	numbers := []int{}

	//s initialize the lineNumbers struct
	// we will use this to store the numbers from each line
	lineNumbers := lineNumbers{}

	for scanner.Scan() {
		// get the line from the file
		line := scanner.Text()
		//parse the line and replace words with numbers kinda
		fmtLine := parseLineStringNumbers(line)
		for _, c := range fmtLine {
			// if is number
			if unicode.IsDigit(c) {
				if lineNumbers.firstNumber == "" {
					lineNumbers.firstNumber = string(c)
				} else {
					lineNumbers.secondNumber = string(c)
				}
			}
		}
		// make sure we have two numbers
		if lineNumbers.secondNumber == "" {
			lineNumbers.secondNumber = lineNumbers.firstNumber
		}

		//fmt.Println(lineNumbers.firstNumber + lineNumbers.secondNumber)

		// Convert string to integer
		number, err := strconv.Atoi(lineNumbers.firstNumber + lineNumbers.secondNumber)
		if err != nil {
			log.Fatalf("Failed converting string to integer: %s", err)
		}

		// Add the number to the slice
		numbers = append(numbers, number)

		lineNumbers.firstNumber = ""  // Reset the numberText for the next line
		lineNumbers.secondNumber = "" // Reset the numberText for the next line
	}

	// sum the numbers
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	// print the sum
	fmt.Println("Sum: ", sum)
}

func parseLineStringNumbers(s string) string {
	//fmt.Println(s)

	// replace the words with numbers
	// we do the t2wo because of the twoone style case

	stringNumbersMap := make(map[string]string)
	stringNumbersMap["zero"] = "z0ero"
	stringNumbersMap["one"] = "o1ne"
	stringNumbersMap["two"] = "t2wo"
	stringNumbersMap["three"] = "t3hree"
	stringNumbersMap["four"] = "f4our"
	stringNumbersMap["five"] = "f5ive"
	stringNumbersMap["six"] = "s6ix"
	stringNumbersMap["seven"] = "s7even"
	stringNumbersMap["eight"] = "e8ight"
	stringNumbersMap["nine"] = "n9ine"
	// do the replace
	for key, value := range stringNumbersMap {
		s = strings.ReplaceAll(s, key, value)
	}

	//fmt.Println(s)

	return s
}

type lineNumbers struct {
	firstNumber  string
	secondNumber string
}
