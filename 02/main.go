package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var testInput = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func main() {
	input := readInput()

	var safe int
	for _, report := range input {
		revReport := make([]int, len(report))
		copy(revReport, report)
		slices.Reverse(revReport)
		result := isSafeDampened([]int{}, report) || isSafeDampened([]int{}, revReport)
		// result := IsSafe(report)  // part 1
		if result {
			safe += 1
		}
	}
	fmt.Println(safe)
}

func isSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}
	return safeInc(report[0], report[1]) && isSafe(report[1:])
}

func isSafeDampened(prev []int, report []int) bool {
	if len(report) < 2 {
		return true
	}
	if safeInc(report[0], report[1]) {
		return isSafeDampened([]int{report[0]}, report[1:])
	}
	return isSafe(append(prev, report[1:]...)) || isSafe(append([]int{report[0]}, report[2:]...))
}

func safeInc(a, b int) bool {
	return a < b && b-a >= 1 && b-a <= 3
}

func readInput() [][]int {
	ret := make([][]int, 0)

	inFile, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scLines := bufio.NewScanner(inFile)
	for scLines.Scan() {
		var report []int
		words := strings.Split(scLines.Text(), " ")
		for _, word := range words {
			level, err := strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
			report = append(report, level)
		}
		ret = append(ret, report)
	}
	return ret
}
