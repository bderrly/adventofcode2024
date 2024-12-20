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

type Result struct {
	s    []int
	safe bool
}

func main() {
	// input := testies
	input := getInput()
	// input := testInput

	var safe int
	for _, report := range input {
		revReport := make([]int, len(report))
		copy(revReport, report)
		slices.Reverse(revReport)
		result := IsSafeDampened(-1, report) || IsSafeDampened(-1, revReport)
		// result := IsSafe(report)
		if result {
			safe += 1
		}
		fmt.Println(Result{report, result})
	}
	fmt.Println(safe)
}

func IsSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}
	return safeInc(report[0], report[1]) && IsSafe(report[1:])
}

var testies = [][]int{
	// {4, 2, 4, 5, 6},    // P
	// {2, 4, 4, 5, 6},    // P
	// {2, 4, 4, 4, 5, 6}, // F
	// {5, 4, 6, 7, 8, 9}, // P
	// {5, 4, 3, 2, 3, 4}, // F
	// {5, 4, 3, 3, 2},    // P
	// {5, 4, 3, 3, 2, 3}, // F
	// {5, 6, 2, 7},       // P
	// {5, 6, 2, 7, 6},    // F
	// {5, 6, 2, 7, 8}, // P
	// {4, 4, 4, 4},       // F
	// {4, 5, 5, 5},       // F
	// {4, 5, 4, 5},       // F
	// {4, 5, 4, 3, 2}, // P
	// {6, 5, 6, 4, 3, 2},    // P
	// {6, 5, 6, 5, 4, 3, 2}, // F
	// {6, 5, 4, 3, 3},    // P
	// {6, 5, 4, 3, 3, 3}, // F
}

func IsSafeDampened(prev int, report []int) bool {
	if len(report) < 2 {
		return true
	}
	if safeInc(report[0], report[1]) {
		return IsSafeDampened(report[0], report[1:])
	} else {
		if prev == -1 {
			return IsSafe(report[1:])
		} else {
			return IsSafe(append(append([]int{}, prev), report[1:]...)) || IsSafe(append(append([]int{}, report[0]), report[2:]...))
		}
	}
}

func safeInc(a, b int) bool {
	return a < b && b-a >= 1 && b-a <= 3
}

func getInput() [][]int {
	ret := make([][]int, 0)

	inFile, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scLines := bufio.NewScanner(inFile)
	for scLines.Scan() {
		var report []int
		scWords := bufio.NewScanner(strings.NewReader(scLines.Text()))
		scWords.Split(bufio.ScanWords)
		for scWords.Scan() {
			val, err := strconv.ParseInt(scWords.Text(), 10, 32)
			if err != nil {
				panic(err)
			}
			report = append(report, int(val))
		}
		ret = append(ret, report)
	}
	return ret
}
