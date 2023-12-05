package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("games.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := []Game{}
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		games = append(games, destructureGameLine(line))
	}

	//fmt.Println(games)
	//fmt.Println(getPossibleGames(games, 12, 14, 13))
	possibleGames := getPossibleGames(games, 12, 14, 13)
	sum := 0
	// for sum of possible games
	for i := 0; i < len(possibleGames); i++ {
		sum += possibleGames[i].id
	}
	fmt.Println("sum of possible games:" + strconv.Itoa(sum))
	sum = 0
	for i := 0; i < len(games); i++ {
		sum += (games[i].maxRed * games[i].maxBlue * games[i].maxGreen)
	}
	fmt.Println("sum of power of games:" + strconv.Itoa(sum))
}

func destructureGameLine(line string) Game {
	// Split string on multiple characters
	splitFunc := func(c rune) bool {
		return c == ',' || c == ';' || c == ':' // Add more delimiters if needed
	}
	parts := strings.FieldsFunc(line, splitFunc)

	var gameHighColors = make(map[string]int)
	entry := strings.Split(parts[0], " ")
	gameID, err := strconv.Atoi(entry[1])
	if err != nil {
		log.Fatalf("Failed converting string to integer: %s", err)
	}
	gameHighColors["green"] = 0
	gameHighColors["blue"] = 0
	gameHighColors["red"] = 0

	for i := 1; i < len(parts); i++ {
		entry := strings.Split(parts[i], " ")
		number, err := strconv.Atoi(entry[1])
		if err != nil {
			log.Fatalf("Failed converting string to integer: %s", err)
		}
		//fmt.Println(entry[2])
		if number >= gameHighColors[entry[2]] {
			gameHighColors[entry[2]] = number
		}
	}

	//fmt.Println(gameHighColors)
	//	fmt.Println(gameID)
	// Process parts here

	return Game{gameID, gameHighColors["green"], gameHighColors["blue"], gameHighColors["red"]}
}

func getPossibleGames(pGames []Game, pRed int, pBlue int, pGreen int) []Game {
	possibleGames := []Game{}

	for i := 0; i < len(pGames); i++ {
		if pGames[i].maxRed <= pRed && pGames[i].maxBlue <= pBlue && pGames[i].maxGreen <= pGreen {
			possibleGames = append(possibleGames, pGames[i])
		}
	}
	return possibleGames
}

type Game struct {
	id       int
	maxGreen int
	maxBlue  int
	maxRed   int
}
