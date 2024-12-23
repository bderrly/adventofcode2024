package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	inputPairs := readInput()
	var total int
	for _, pair := range inputPairs {
		total += pair.A * pair.B
	}
	fmt.Println(total)
}

type mulPair struct {
	A, B int
}

func readInput() []mulPair {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	rgx := regexp.MustCompile(`do(?:n't)?\(\)|mul\((?<a>\d{1,3}),(?<b>\d{1,3})\)`)
	var pairs []mulPair
	sc := bufio.NewScanner(inFile)
	var do bool = true
	for sc.Scan() {
		for _, match := range rgx.FindAllStringSubmatch(sc.Text(), -1) {
			switch match[0] {
			case "do()":
				do = true
			case "don't()":
				do = false
			default:
				if do && len(match) == 3 {
					a, err := strconv.Atoi(match[1])
					if err != nil {
						panic(err)
					}
					b, err := strconv.Atoi(match[2])
					if err != nil {
						panic(err)
					}
					pairs = append(pairs, mulPair{a, b})
				}
			}
		}
	}
	return pairs
}
