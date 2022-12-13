package main

import (
	"fmt"
	"strconv"
	"strings"
)

// type stacks []stack

type stack struct {
	crates []crate
}

type crate string

var allStacks []stack

func main() {
	// lines, _ := readLines("example.txt")
	// startingLine := 6
	// var stack1, stack2, stack3 stack
	// stack1.crates = append(stack1.crates, "Z")
	// stack1.crates = append(stack1.crates, "N")
	// stack2.crates = append(stack2.crates, "M")
	// stack2.crates = append(stack2.crates, "C")
	// stack2.crates = append(stack2.crates, "D")
	// stack3.crates = append(stack3.crates, "P")
	lines, _ := readLines("input.txt")
	startingLine := 11
	var stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9 stack
	stack1.crates = append(stack1.crates, "H")
	stack1.crates = append(stack1.crates, "T")
	stack1.crates = append(stack1.crates, "Z")
	stack1.crates = append(stack1.crates, "D")
	stack2.crates = append(stack2.crates, "Q")
	stack2.crates = append(stack2.crates, "R")
	stack2.crates = append(stack2.crates, "W")
	stack2.crates = append(stack2.crates, "T")
	stack2.crates = append(stack2.crates, "G")
	stack2.crates = append(stack2.crates, "C")
	stack2.crates = append(stack2.crates, "S")
	stack3.crates = append(stack3.crates, "P")
	stack3.crates = append(stack3.crates, "B")
	stack3.crates = append(stack3.crates, "F")
	stack3.crates = append(stack3.crates, "Q")
	stack3.crates = append(stack3.crates, "N")
	stack3.crates = append(stack3.crates, "R")
	stack3.crates = append(stack3.crates, "C")
	stack3.crates = append(stack3.crates, "H")
	stack4.crates = append(stack4.crates, "L")
	stack4.crates = append(stack4.crates, "C")
	stack4.crates = append(stack4.crates, "N")
	stack4.crates = append(stack4.crates, "F")
	stack4.crates = append(stack4.crates, "H")
	stack4.crates = append(stack4.crates, "Z")
	stack5.crates = append(stack5.crates, "G")
	stack5.crates = append(stack5.crates, "L")
	stack5.crates = append(stack5.crates, "F")
	stack5.crates = append(stack5.crates, "Q")
	stack5.crates = append(stack5.crates, "S")
	stack6.crates = append(stack6.crates, "V")
	stack6.crates = append(stack6.crates, "P")
	stack6.crates = append(stack6.crates, "W")
	stack6.crates = append(stack6.crates, "Z")
	stack6.crates = append(stack6.crates, "B")
	stack6.crates = append(stack6.crates, "R")
	stack6.crates = append(stack6.crates, "C")
	stack6.crates = append(stack6.crates, "S")
	stack7.crates = append(stack7.crates, "Z")
	stack7.crates = append(stack7.crates, "F")
	stack7.crates = append(stack7.crates, "J")
	stack8.crates = append(stack8.crates, "D")
	stack8.crates = append(stack8.crates, "L")
	stack8.crates = append(stack8.crates, "V")
	stack8.crates = append(stack8.crates, "Z")
	stack8.crates = append(stack8.crates, "R")
	stack8.crates = append(stack8.crates, "H")
	stack8.crates = append(stack8.crates, "Q")
	stack9.crates = append(stack9.crates, "B")
	stack9.crates = append(stack9.crates, "H")
	stack9.crates = append(stack9.crates, "G")
	stack9.crates = append(stack9.crates, "N")
	stack9.crates = append(stack9.crates, "F")
	stack9.crates = append(stack9.crates, "Z")
	stack9.crates = append(stack9.crates, "L")
	stack9.crates = append(stack9.crates, "D")
	allStacks = append(allStacks, stack1)
	allStacks = append(allStacks, stack2)
	allStacks = append(allStacks, stack3)
	allStacks = append(allStacks, stack4)
	allStacks = append(allStacks, stack5)
	allStacks = append(allStacks, stack6)
	allStacks = append(allStacks, stack7)
	allStacks = append(allStacks, stack8)
	allStacks = append(allStacks, stack9)
	_ = lines
	buildStacks(startingLine, lines)
	returnTopCrates()
}

func buildStacks(startingLine int, lines []string) []stack {
	// build allStacks from the first set of lines until we find a blank line
	var stacks []stack
	for i := startingLine - 1; i < len(lines); i++ {
		// fmt.Println(lines[i])
		amount, originalStack, destinationStack := interpretInstructions(lines[i])
		movingCrates := getCrates(amount, originalStack)
		// fmt.Println("Moving the following crates: ", movingCrates)
		// fmt.Println("Destination stack: ", destinationStack)
		placeCrates(movingCrates, destinationStack)
		fmt.Println("Full Stackset: ", allStacks)
	}
	return stacks
}

func interpretInstructions(line string) (amount, originalStack, destinationStack int) {
	substrings := strings.Split(line, " ")
	amount, _ = strconv.Atoi(substrings[1])
	originalStack, _ = strconv.Atoi(substrings[3])
	destinationStack, _ = strconv.Atoi(substrings[5])
	return amount, originalStack - 1, destinationStack - 1
}

func getCrates(amount, originatingStack int) (crates []crate) {
	poppedCrates := allStacks[originatingStack].crates[(len(allStacks[originatingStack].crates) - amount):]
	allStacks[originatingStack].crates = allStacks[originatingStack].crates[0 : len(allStacks[originatingStack].crates)-amount]
	fmt.Println("poppedCrates: ", poppedCrates)
	return poppedCrates
}

func placeCrates(crates []crate, destinationStack int) {
	allStacks[destinationStack].crates = append(allStacks[destinationStack].crates, crates...)
}

func returnTopCrates() {
	for i := 0; i < len(allStacks); i++ {
		fmt.Printf("%s", allStacks[i].crates[len(allStacks[i].crates)-1])
	}
}
