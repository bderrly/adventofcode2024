package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const inputFile = "input.txt"

var testInput = [][]int{
	{3, 4},
	{4, 3},
	{2, 5},
	{1, 3},
	{3, 9},
	{3, 3},
}

func main() {
	var left []int
	var right []int

	input := inputSlice()
	// input := testInput
	for i := 0; i < len(input); i++ {
		left = append(left, input[i][0])
		right = append(right, input[i][1])
	}
	sort.Ints(left)
	sort.Ints(right)

	var distance int
	for i := 0; i < len(input); i++ {
		dist := left[i] - right[i]
		if dist < 0 {
			dist *= -1
		}
		distance += dist
	}
	fmt.Println(distance)

	rightCounts := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0}
	for _, r := range right {
		rightCounts[r] += 1
	}

	var similarity int
	for _, l := range left {
		similarity += l * rightCounts[l]
	}
	fmt.Println(similarity)
}

func inputSlice() [][]int {
	ret := make([][]int, 0)

	fReader, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	var left, right int
	scan := bufio.NewScanner(fReader)
	for scan.Scan() {
		n, err := fmt.Sscanf(scan.Text(), "%d%d", &left, &right)
		if n != 2 {
			panic("did not find two inputs")
		}
		if err != nil {
			panic(err)
		}
		ret = append(ret, []int{left, right})
	}
	return ret
}
