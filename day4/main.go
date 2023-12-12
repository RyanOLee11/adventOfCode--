package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	totals := []int{}
	ScratchCards := []ScratchCard{}

	for scanner.Scan() {
		line := scanner.Text()
		index := strings.Index(line, ":")
		line = line[index+1:]
		index = strings.Index(line, "|")
		winningNumbers, myNumbers := line[:index], line[index+1:]
        winningNumbersSlice := strings.Fields(winningNumbers)
        myNumbersSlice := strings.Fields(myNumbers)

		currentTotal := 0
		winningNumberCount := 0
		for _, myNumber := range myNumbersSlice {
            if contains(winningNumbersSlice, myNumber) {
                if currentTotal == 0 {
					currentTotal = 1
				}else{
					currentTotal = currentTotal * 2
				}
				winningNumberCount++
            }
        }
		totals = append(totals, currentTotal)
		ScratchCards = append(ScratchCards, ScratchCard{winningNumberCount, 1})
	}
	sum := 0
	for _, total := range totals {
		sum += total
	}
	//fmt.Println(sum)
	fmt.Println(totals)
	increaseInstanceCountBasedOnWinningNumbers(ScratchCards)
	//fmt.Println(ScratchCards)
	sum = 0
	for _, scratchCard := range ScratchCards {
		sum += scratchCard.instanceCount
	}
	fmt.Println(sum)
}

func contains(slice []string, item string) bool {
    for _, a := range slice {
        if a == item {
            return true
        }
    }
    return false
}
func increaseInstanceCountBasedOnWinningNumbers(scratchCards []ScratchCard) {
    for i := 0; i < len(scratchCards)-1; i++ {
		for x := 1; x <= scratchCards[i].winningNumberCount; x++ {
			scratchCards[i + x].instanceCount = scratchCards[i + x].instanceCount + scratchCards[i].instanceCount
		}
       
    }
}

type ScratchCard struct {
	winningNumberCount int
	instanceCount int
}