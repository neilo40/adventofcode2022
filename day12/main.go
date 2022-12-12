package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/neilo40/adventofcode2022/helper"
)

type node struct {
	id         string
	height     int
	isEnd      bool
	neighbours []*node
}

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

var end *node

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var start *node

	// start by making a grid of nodes
	grid := make([][]*node, 41)
	for r := 0; r < 41; r++ {
		grid[r] = make([]*node, 113)
		for c := 0; c < 113; c++ {
			grid[r][c] = &node{}
		}
	}
	row := 0
	for scanner.Scan() {
		t := scanner.Text()
		for col, h := range t {
			n := &node{id: fmt.Sprintf("%d_%d", row, col), neighbours: make([]*node, 0, 4)}
			if h == 'S' {
				n.height = int('a')
				start = n
			} else if h == 'E' {
				n.height = int('z')
				n.isEnd = true
				end = n
			} else {
				n.height = int(h)
			}
			grid[row][col] = n
		}
		row++
	}

	// now add all the valid neighbours to each node to form a graph
	for r := 0; r < 41; r++ {
		for c := 0; c < 113; c++ {
			this := grid[r][c]
			// can go up 0 or 1, can go down any
			if r > 0 {
				// up
				next := grid[r-1][c]
				if next.height-this.height <= 1 {
					this.neighbours = append(this.neighbours, next)
				}
			}
			if r < 40 {
				// down
				next := grid[r+1][c]
				if next.height-this.height <= 1 {
					this.neighbours = append(this.neighbours, next)
				}
			}
			if c > 0 {
				// left
				next := grid[r][c-1]
				if next.height-this.height <= 1 {
					this.neighbours = append(this.neighbours, next)
				}
			}
			if c < 112 {
				// right
				next := grid[r][c+1]
				if next.height-this.height <= 1 {
					this.neighbours = append(this.neighbours, next)
				}
			}
		}
	}

	// Recurse through each neighbour keeping track of node counts until we get to E

	minPath := getPath(map[string]bool{}, start)
	log.Printf("Shortest path from S to E is %d\n", len(minPath))
}

// maps are passed by reference, so we need to pass a clone to each neighbour
// unfortunately this slows things down considerably
func cloneMap(m map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func getPath(seen map[string]bool, n *node) map[string]bool {
	// seen: nodes that have been seen to date on this path
	// n : current node
	seen[n.id] = true

	if n.isEnd {
		// we got there
		return seen
	}

	// get path lengths through all un-seen neighbours
	paths := make(map[string]map[string]bool)
	for _, neighbour := range n.neighbours {
		_, ok := seen[neighbour.id]
		if !ok {
			paths[neighbour.id] = getPath(cloneMap(seen), neighbour)
		}
	}

	// return shortest path
	min := math.MaxInt
	var minPath map[string]bool
	for _, p := range paths {
		_, ok := p[end.id]
		if !ok {
			continue // path didn't get to the end
		}
		if len(p) < min {
			min = len(p)
			minPath = p
		}
	}

	return minPath
}

func part2() {
}
