package main

import (
	"log"
	"strconv"
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	highestValue, currentValue := 0, 0
	for i, line := range lines {
		if line != "" {
			intValue, _ := strconv.Atoi(line)
			currentValue += intValue
			if currentValue > highestValue {
				highestValue = currentValue
			}
		} else {
			currentValue = 0
		}
		if i == len(lines)-1 {
			log.Println("highestValue: ", highestValue)
		}
	}
}
