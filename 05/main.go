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

type orderRule struct {
	First, Last int
}

type pageUpdates []int

func main() {
	var sum int
	orderingRules, pages := readInput()
	for _, page := range pages {
		if validUpdate(orderingRules, page) {
			sum += page[len(page)/2]
		}
	}
	fmt.Println(sum)
}

func validUpdate(rules []orderRule, pages pageUpdates) bool {
	for _, rule := range rules {
		if slices.Contains(pages, rule.First) && slices.Contains(pages, rule.Last) {
			if slices.Index(pages, rule.First) > slices.Index(pages, rule.Last) {
				return false
			}
		}
	}
	return true
}

func readInput() ([]orderRule, []pageUpdates) {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	var ordering []orderRule
	var updates []pageUpdates

	scan := bufio.NewScanner(inFile)
	for scan.Scan() {
		line := scan.Text()
		if strings.Contains(line, "|") {
			ordering = append(ordering, parseOrdering(line))
		} else if strings.Contains(line, ",") {
			updates = append(updates, parseUpdate(line))
		}
	}
	return ordering, updates
}

func parseOrdering(line string) (rule orderRule) {
	fmt.Sscanf(line, "%d|%d", &rule.First, &rule.Last)
	return
}

func parseUpdate(line string) (update pageUpdates) {
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
