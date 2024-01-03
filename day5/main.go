package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Seed struct {
	value uint64
}

type AtoB struct {
	destinationRangeStart uint64
	sourceRangeStart      uint64
	rangeLength           uint64
}

var seeds []Seed

var mappers [][]AtoB = make([][]AtoB, 7)
var minLocation uint64 = math.MaxUint64

var mu sync.Mutex

var re = regexp.MustCompile("[0-9]+")

func main() {
	lines := ReadFile("./puzzle.txt")

	// split by empty line
	split := strings.Split(strings.Join(lines, "\n"), "\n\n")

	seedsStr := split[0]
	seedsNumbersStr := re.FindAllString(seedsStr, -1)

	for _, seedNumberStr := range seedsNumbersStr {
		seedNumber, err := strconv.Atoi(seedNumberStr)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, Seed{uint64(seedNumber)})
	}

	var seedsPairs [][]Seed
	for i := 0; i < len(seeds); i += 2 {
		seedsPairs = append(seedsPairs, []Seed{seeds[i], seeds[i+1]})
	}

	for index, split := range split[1:] {
		buildMap(split, index+1)
	}

	var wg sync.WaitGroup
	wg.Add(len(seedsPairs))

	for _, seedPair := range seedsPairs {
		fmt.Println("Calculating for seed pair", seedPair[0].value, seedPair[1].value)
		go runForSeedPair(seedPair, &wg)
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Min location is", minLocation)
	fmt.Println("Done!")
}

func runForSeedPair(seedPair []Seed, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := seedPair[0].value; i <= seedPair[0].value+seedPair[1].value-1; i++ {
		calculateForSource(Seed{i}.value, 0)
	}
}

func calculateForSource(source uint64, index uint64) {
	nextSource := source
	for _, mapper := range mappers[index] {
		if isBetween(source, mapper.sourceRangeStart, mapper.sourceRangeStart+mapper.rangeLength-1) {
			nextSource = source + (mapper.destinationRangeStart - mapper.sourceRangeStart)
		}
	}

	if index < 6 {
		calculateForSource(nextSource, index+1)
	} else {
		if nextSource < minLocation {
			mu.Lock()
			minLocation = nextSource
			mu.Unlock()
		}
	}
}

func buildMap(split string, index int) {
	elementsWithoutTitle := strings.Split(split, "\n")[1:]
	build(elementsWithoutTitle, &mappers[index-1])
}

func build(lines []string, mapToUpdate *[]AtoB) {
	for _, line := range lines {
		elementStr := re.FindAllString(line, -1)

		destinationRangeStart, err := strconv.Atoi(elementStr[0])
		if err != nil {
			panic(err)
		}
		sourceRangeStart, err := strconv.Atoi(elementStr[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(elementStr[2])
		if err != nil {
			panic(err)
		}

		*mapToUpdate = append(*mapToUpdate, AtoB{uint64(destinationRangeStart), uint64(sourceRangeStart), uint64(rangeLength)})
	}
}

func isBetween(num, min, max uint64) bool {
	return num >= min && num <= max
}

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines
}