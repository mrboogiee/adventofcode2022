package main

import (
	"log"
	"strconv"
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	var elves []int
	currentCarrying, elveIndex := 0, 0
	bigThreeCarriers := []int{0, 0, 0} // big three carriers of calories
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
	for _, elve := range elves {
		if elve > bigThreeCarriers[0] { // if the elve is bigger than the number one, shift all one down
			bigThreeCarriers[2] = bigThreeCarriers[1]
			bigThreeCarriers[1] = bigThreeCarriers[0]
			bigThreeCarriers[0] = elve
		} else if elve > bigThreeCarriers[1] { // if the elve is bigger than the number 2, shift 2 and 3 down
			bigThreeCarriers[2] = bigThreeCarriers[1]
			bigThreeCarriers[1] = elve
			bigThreeCarriers[1] = elve
		} else if elve > bigThreeCarriers[2] { // if the elve is bigger than the number 3, replace number 3
			bigThreeCarriers[2] = elve
		}
	}
	log.Println("Total calories carried by the top 3: ", bigThreeCarriers[0]+bigThreeCarriers[1]+bigThreeCarriers[2])
}
