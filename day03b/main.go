package main

import (
	"fmt"
	"strings"
)

type rucksack struct {
	items string
}

type rucksacks []rucksack

var (
	allRucksacks   []rucksack
	rucksackGroups []rucksacks
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	_ = lines
	for _, line := range lines {
		var newRucksack rucksack
		newRucksack.items = line
		allRucksacks = append(allRucksacks, newRucksack)
	}
	splitRucksacksByGroup()
	score := 0
	fmt.Println(len(rucksackGroups))
	for i, rucksackGroup := range rucksackGroups {
		fmt.Println(i, len(rucksackGroup))
		commonItem := findCommonItem(rucksackGroup[0], rucksackGroup[1], rucksackGroup[2])
		score += getItemValue(commonItem)
	}

	fmt.Println(score)
}

func splitRucksacksByGroup() {
	counter := 0
	var newRucksackGroup rucksacks
	rucksackGroups = append(rucksackGroups, newRucksackGroup)
	for i, rucksack := range allRucksacks {
		rucksackGroups[counter] = append(rucksackGroups[counter], rucksack)
		if i%3 == 2 {
			counter++
			if counter < len(allRucksacks)/3 {
				rucksackGroups = append(rucksackGroups, newRucksackGroup)
			}
		}
	}
}

func findCommonItem(rucksack1, rucksack2, rucksack3 rucksack) string {
	for _, item := range rucksack1.items {
		if strings.Contains(rucksack2.items, string(item)) && strings.Contains(rucksack3.items, string(item)) {
			return string(item)
		}
	}
	return ""
}

func compareRucksackItems(compartment1, compartment2 []string) int {
	for _, item := range compartment1 {
		for _, item2 := range compartment2 {
			if item == item2 {
				return getItemValue(item)
			}
		}
	}
	return 0
}

func getItemValue(item string) int {
	a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z := 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z := 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52
	switch item {
	case "a":
		return a
	case "b":
		return b
	case "c":
		return c
	case "d":
		return d
	case "e":
		return e
	case "f":
		return f
	case "g":
		return g
	case "h":
		return h
	case "i":
		return i
	case "j":
		return j
	case "k":
		return k
	case "l":
		return l
	case "m":
		return m
	case "n":
		return n
	case "o":
		return o
	case "p":
		return p
	case "q":
		return q
	case "r":
		return r
	case "s":
		return s
	case "t":
		return t
	case "u":
		return u
	case "v":
		return v
	case "w":
		return w
	case "x":
		return x
	case "y":
		return y
	case "z":
		return z
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "D":
		return D
	case "E":
		return E
	case "F":
		return F
	case "G":
		return G
	case "H":
		return H
	case "I":
		return I
	case "J":
		return J
	case "K":
		return K
	case "L":
		return L
	case "M":
		return M
	case "N":
		return N
	case "O":
		return O
	case "P":
		return P
	case "Q":
		return Q
	case "R":
		return R
	case "S":
		return S
	case "T":
		return T
	case "U":
		return U
	case "V":
		return V
	case "W":
		return W
	case "X":
		return X
	case "Y":
		return Y
	case "Z":
		return Z
	}
	return 0
}
