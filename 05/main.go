package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

var orderingRules map[int][]int

func main() {
	orderingRules = make(map[int][]int)

	var sum, fixedSum int
	pageUpdates := readInput()
	for _, pageUpdate := range pageUpdates {
		if validPageOrdering(pageUpdate) {
			sum += pageUpdate[len(pageUpdate)/2]
		} else {
			fixedPage := fixPageOrdering(pageUpdate)
			fixedSum += fixedPage[len(fixedPage)/2]
		}
	}
	fmt.Println(sum)
	fmt.Println(fixedSum)
}

func fixPageOrdering(pages []int) []int {
	updatedPages := make([]int, len(pages))
	copy(updatedPages, pages)

	var orderedUpdateRules []int
	for k := range orderingRules {
		orderedUpdateRules = append(orderedUpdateRules, k)
	}
	slices.Sort(orderedUpdateRules)

	triggeredUpdates := make([][]int, 0)

	for _, rule := range orderedUpdateRules {
		for _, update := range orderingRules[rule] {
			if slices.Contains(updatedPages, rule) && slices.Contains(updatedPages, update) {
				ruleIdx := slices.Index(updatedPages, rule)
				updateIdx := slices.Index(updatedPages, update)
				if ruleIdx < updateIdx {
					continue
				}
				triggeredUpdates = append(triggeredUpdates, []int{rule, update})
				updatedPages = slices.Delete(updatedPages, ruleIdx, ruleIdx+1)
				if updateIdx < 0 {
					updateIdx = 0
				}
				updatedPages = slices.Insert(updatedPages, updateIdx, rule)
			}
		}
	}
	if !validPageOrdering(updatedPages) {
		panic(fmt.Sprintf("fixed but invalid: %v", updatedPages))
	}
	return updatedPages
}

func validPageOrdering(pages []int) bool {
	for k, vals := range orderingRules {
		for _, val := range vals {
			if slices.Contains(pages, k) && slices.Contains(pages, val) {
				if slices.Index(pages, k) > slices.Index(pages, val) {
					return false
				}
			}
		}
	}
	return true
}

func readInput() [][]int {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	var updates [][]int

	scan := bufio.NewScanner(inFile)
	for scan.Scan() {
		line := scan.Text()
		if strings.Contains(line, "|") {
			parseOrdering(line)
		} else if strings.Contains(line, ",") {
			updates = append(updates, parseUpdate(line))
		}
	}
	return updates
}

func parseOrdering(line string) {
	var key, value int
	fmt.Sscanf(line, "%d|%d", &key, &value)
	orderingRules[key] = append(orderingRules[key], value)
}

func parseUpdate(line string) (update []int) {
	digitScanner := bufio.NewScanner(strings.NewReader(line))
	digitScanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		start := 0
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if r == ',' {
				return i + width, data[start:i], nil
			}
		}
		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}
		return start, nil, nil
	})

	for digitScanner.Scan() {
		page, err := strconv.Atoi(digitScanner.Text())
		if err != nil {
			panic(err)
		}
		update = append(update, page)
	}
	return
}
