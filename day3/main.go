package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file , err := os.Open("schematic.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	matrix := [][]string{}
	row := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
		line := scanner.Text()
		for col, c := range line {
			matrix[row][col] = string(c)
		}
		row++
	}

	fmt.Print(matrix)
}