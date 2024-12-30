package main

import (
	"bufio"
	"fmt"
	"os"
)

var xmas []byte = []byte{'X', 'M', 'A', 'S'}

var (
	north     = direction{0, -1}
	northeast = direction{1, -1}
	east      = direction{1, 0}
	southeast = direction{1, 1}
	south     = direction{0, 1}
	southwest = direction{-1, 1}
	west      = direction{-1, 0}
	northwest = direction{-1, -1}

	directions = []direction{
		north, northeast, east, southeast, south, southwest, west, northwest,
	}
)

func main() {
	graph := toGraph(readInput())
	var total int
	for x := 0; x < len(graph.Vertices); x++ {
		for y := 0; y < len(graph.Vertices[x]); y++ {
			total += graph.Search(point{x, y})
		}
	}
	fmt.Println(total)
}

type node struct {
	Letter byte
}

type graph struct {
	Vertices [][]node
}

type point struct {
	X, Y int
}

type direction struct {
	X, Y int
}

func pointDirection(p point, d direction) []point {
	var result []point
	for i := 0; i < 4; i++ {
		result = append(result, point{p.X + i*d.X, p.Y + i*d.Y})
	}
	return result
}

func (g graph) Node(p point) *node {
	if p.X >= 0 && p.Y >= 0 {
		if p.X < len(g.Vertices) && p.Y < len(g.Vertices[p.X]) {
			return &g.Vertices[p.X][p.Y]
		}
	}
	return nil
}

// Search returns the number of words found starting from the given point.
func (g graph) Search(pt point) int {
	var total int
	words := g.makeSlices(pt)
	for _, word := range words {
		var good bool
		for i := 0; i < 4; i++ {
			if word[i] != xmas[i] {
				good = false
				break
			}
			good = true
		}
		if good {
			total++
		}
	}
	return total
}

func (g graph) makeSlices(p point) [][]byte {
	result := make([][]byte, 0)

	for _, dir := range directions {
		points := pointDirection(p, dir)
		var word []byte
		for _, pt := range points {
			if g.Node(pt) == nil {
				break
			}
			word = append(word, g.Node(pt).Letter)
		}
		if len(word) == 4 {
			result = append(result, word)
		}
	}
	return result
}

func toGraph(input [][]byte) graph {
	var result graph
	result.Vertices = make([][]node, len(input))
	for i, row := range input {
		result.Vertices[i] = make([]node, len(row))
		for j, letter := range row {
			result.Vertices[i][j] = node{Letter: letter}
		}
	}
	return result
}

func readInput() [][]byte {
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	var result [][]byte
	scan := bufio.NewScanner(inFile)
	for scan.Scan() {
		line := scan.Text()
		result = append(result, []byte(line))
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}
	return result
}
