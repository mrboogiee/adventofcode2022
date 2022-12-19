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
	for i := 13; i < len(line); i++ {
		if line[i] != line[i-13] &&
			line[i] != line[i-12] &&
			line[i] != line[i-11] &&
			line[i] != line[i-10] &&
			line[i] != line[i-9] &&
			line[i] != line[i-8] &&
			line[i] != line[i-7] &&
			line[i] != line[i-6] &&
			line[i] != line[i-5] &&
			line[i] != line[i-4] &&
			line[i] != line[i-3] &&
			line[i] != line[i-2] &&
			line[i] != line[i-1] &&
			line[i-1] != line[i-13] &&
			line[i-1] != line[i-12] &&
			line[i-1] != line[i-11] &&
			line[i-1] != line[i-10] &&
			line[i-1] != line[i-9] &&
			line[i-1] != line[i-8] &&
			line[i-1] != line[i-7] &&
			line[i-1] != line[i-6] &&
			line[i-1] != line[i-5] &&
			line[i-1] != line[i-4] &&
			line[i-1] != line[i-3] &&
			line[i-2] != line[i-13] &&
			line[i-2] != line[i-12] &&
			line[i-2] != line[i-11] &&
			line[i-2] != line[i-10] &&
			line[i-2] != line[i-9] &&
			line[i-2] != line[i-8] &&
			line[i-2] != line[i-7] &&
			line[i-2] != line[i-6] &&
			line[i-2] != line[i-5] &&
			line[i-2] != line[i-4] &&
			line[i-2] != line[i-3] &&
			line[i-3] != line[i-13] &&
			line[i-3] != line[i-12] &&
			line[i-3] != line[i-11] &&
			line[i-3] != line[i-10] &&
			line[i-3] != line[i-9] &&
			line[i-3] != line[i-8] &&
			line[i-3] != line[i-7] &&
			line[i-3] != line[i-6] &&
			line[i-3] != line[i-5] &&
			line[i-4] != line[i-13] &&
			line[i-4] != line[i-12] &&
			line[i-4] != line[i-11] &&
			line[i-4] != line[i-10] &&
			line[i-4] != line[i-9] &&
			line[i-4] != line[i-8] &&
			line[i-4] != line[i-7] &&
			line[i-4] != line[i-6] &&
			line[i-4] != line[i-5] &&
			line[i-5] != line[i-13] &&
			line[i-5] != line[i-12] &&
			line[i-5] != line[i-11] &&
			line[i-5] != line[i-10] &&
			line[i-5] != line[i-9] &&
			line[i-5] != line[i-8] &&
			line[i-5] != line[i-7] &&
			line[i-5] != line[i-6] &&
			line[i-6] != line[i-13] &&
			line[i-6] != line[i-12] &&
			line[i-6] != line[i-11] &&
			line[i-6] != line[i-10] &&
			line[i-6] != line[i-9] &&
			line[i-6] != line[i-8] &&
			line[i-6] != line[i-7] &&
			line[i-7] != line[i-13] &&
			line[i-7] != line[i-12] &&
			line[i-7] != line[i-11] &&
			line[i-7] != line[i-10] &&
			line[i-7] != line[i-9] &&
			line[i-7] != line[i-8] &&
			line[i-8] != line[i-13] &&
			line[i-8] != line[i-12] &&
			line[i-8] != line[i-11] &&
			line[i-8] != line[i-10] &&
			line[i-8] != line[i-9] &&
			line[i-9] != line[i-13] &&
			line[i-9] != line[i-12] &&
			line[i-9] != line[i-11] &&
			line[i-9] != line[i-10] &&
			line[i-10] != line[i-13] &&
			line[i-10] != line[i-12] &&
			line[i-10] != line[i-11] &&
			line[i-11] != line[i-13] &&
			line[i-11] != line[i-12] {
			fmt.Println("found it on line", i+1)
			return i + 1
		}
	}
	return 0
}
