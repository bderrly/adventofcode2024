package main

import (
	"bufio"
	"fmt"
	"os"
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
	// input := testInput
	input := getInput()

	var results []bool
	var safe int
	for _, report := range input {
		result := IsSafe(report)
		if result {
			safe += 1
		}
		results = append(results, result)
	}
	fmt.Println(safe)
}

func IsSafe(report []int) bool {
	var inc bool = report[1]-report[0] > 0

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if inc {
			if diff <= 0 {
				return false
			} else if diff < 1 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 {
				return false
			} else if diff > -1 || diff < -3 {
				return false
			}
		}
	}
	return true
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
