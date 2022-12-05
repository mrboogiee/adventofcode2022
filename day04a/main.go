package main

import (
	"fmt"
	"strconv"
	"strings"
)

type elvenPairWorkload struct {
	one, two []int
}

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	_ = lines
	var allElvenWorkloads []elvenPairWorkload
	for _, line := range lines {
		workloads := strings.Split(line, ",")
		var team elvenPairWorkload
		for i, workload := range workloads {
			workloadStrings := strings.Split(workload, "-")
			startingWorkload, _ := strconv.Atoi(workloadStrings[0])
			endingWorkload, _ := strconv.Atoi(workloadStrings[1])
			definedWorkload := []int{startingWorkload, endingWorkload}
			expandedWorkload := expandWorkload(definedWorkload)
			if i == 0 {
				team.one = expandedWorkload
			} else {
				team.two = expandedWorkload
			}
		}
		allElvenWorkloads = append(allElvenWorkloads, team)
	}
	var countedContainedWorkloads int
	for _, elvenWorkload := range allElvenWorkloads {
		if subslice(elvenWorkload.one, elvenWorkload.two) {
			countedContainedWorkloads++
		} else if subslice(elvenWorkload.two, elvenWorkload.one) {
			countedContainedWorkloads++
		}
	}
	fmt.Println(countedContainedWorkloads)
}

func expandWorkload(workloads []int) []int {
	var expandedWorkload []int
	counter := workloads[0]
	for i := workloads[0]; i < workloads[1]+1; i++ {
		expandedWorkload = append(expandedWorkload, counter)
		counter++
	}
	return expandedWorkload
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func subslice(s1 []int, s2 []int) bool {
	fmt.Println(s1)
	fmt.Println(s2)
	if len(s1) > len(s2) {
		fmt.Println("false")
		return false
	}
	for _, e := range s1 {
		if !contains(s2, e) {
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("true")
	return true
}
