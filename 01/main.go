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
	for i := 0; i < len(input); i++ {
		left = append(left, input[i][0])
		right = append(right, input[i][1])
	}
	sort.Ints(left)
	sort.Ints(right)

	var sum int
	for i := 0; i < len(input); i++ {
		dist := left[i] - right[i]
		if dist < 0 {
			dist *= -1
		}
		sum += dist
	}
	fmt.Println(sum)
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
