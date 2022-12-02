package main

import (
	"log"
	"sort"
	"strconv"
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	var elves []int
	currentCarrying, elveIndex := 0, 0
	for i, line := range lines {
		if line != "" { // non-empty line so elve is carrying something
			intValue, _ := strconv.Atoi(line)
			currentCarrying += intValue
		} else { // empty line, so new elve starts
			elves = append(elves, currentCarrying)
			elveIndex++
			currentCarrying = 0
		}
		if i == len(lines)-1 { // last line
			elves = append(elves, currentCarrying)
		}
	}
	sort.Ints(elves)
	log.Println("Total calories carried by the top 3: ", elves[len(elves)-1]+elves[len(elves)-2]+elves[len(elves)-3])
}
