package main

import "fmt"

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	for _, line := range lines {
		fmt.Println(findMarker(line))
	}

}

func findMarker(line string) int {
	for i := 3; i < len(line); i++ {
		// fmt.Println(i - 4)
		if line[i] != line[i-3] &&
			line[i] != line[i-2] &&
			line[i] != line[i-1] &&
			line[i-1] != line[i-3] &&
			line[i-1] != line[i-2] &&
			line[i-2] != line[i-3] {
			fmt.Println("found it on line", i+1)
			return i + 1
		}
	}
	return 0
}
