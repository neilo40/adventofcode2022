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

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var start *node
	seen := make(map[string]bool)
	distance := make(map[string]int)
	idToNode := make(map[string]*node)

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
			seen[n.id] = false
			distance[n.id] = math.MaxInt
			idToNode[n.id] = n
			if h == 'S' {
				n.height = int('a')
				start = n
				distance[n.id] = 0
			} else if h == 'E' {
				n.height = int('z')
				n.isEnd = true
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
			neighbours := make([]*node, 0, 4)
			if r > 0 { // up
				neighbours = append(neighbours, grid[r-1][c])
			}
			if r < 40 { // down
				neighbours = append(neighbours, grid[r+1][c])
			}
			if c > 0 { // left
				neighbours = append(neighbours, grid[r][c-1])
			}
			if c < 112 { // right
				neighbours = append(neighbours, grid[r][c+1])
			}
			for _, n := range neighbours {
				// can go up 0 or 1, can go down any
				if n.height-this.height <= 1 {
					this.neighbours = append(this.neighbours, n)
				}
			}
		}
	}

	// Now we Dijkstra...
	currentNode := start
	for {
		if currentNode.isEnd {
			log.Printf("Shortest path from S to E is %d\n", distance[currentNode.id])
			break
		}

		d := distance[currentNode.id] + 1 // all edges have a weight of 1
		for _, n := range currentNode.neighbours {
			if !seen[n.id] && d < distance[n.id] {
				distance[n.id] = d
			}
		}
		seen[currentNode.id] = true
		currentNode = getNextNode(seen, distance, idToNode)
	}

}

func getNextNode(seen map[string]bool, dist map[string]int, idToNode map[string]*node) *node {
	minDist := math.MaxInt
	minId := ""
	for nid, d := range dist {
		if !seen[nid] && d < minDist {
			minDist = d
			minId = nid
		}
	}
	return idToNode[minId]
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var start *node
	seen := make(map[string]bool)
	distance := make(map[string]int)
	idToNode := make(map[string]*node)

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
			seen[n.id] = false
			distance[n.id] = math.MaxInt
			idToNode[n.id] = n
			if h == 'E' { // start at the end and work back
				n.height = int('z')
				start = n
				distance[n.id] = 0
			} else if h == 'S' {
				n.height = int('a')
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
			neighbours := make([]*node, 0, 4)
			if r > 0 { // up
				neighbours = append(neighbours, grid[r-1][c])
			}
			if r < 40 { // down
				neighbours = append(neighbours, grid[r+1][c])
			}
			if c > 0 { // left
				neighbours = append(neighbours, grid[r][c-1])
			}
			if c < 112 { // right
				neighbours = append(neighbours, grid[r][c+1])
			}
			for _, n := range neighbours {
				// in reverse, can go up any, but only down 0 or 1
				if this.height-n.height <= 1 {
					this.neighbours = append(this.neighbours, n)
				}
			}
		}
	}

	// Now we Dijkstra...
	currentNode := start
	for {
		if currentNode.height == int('a') {
			log.Printf("Shortest path from E to any a is %d\n", distance[currentNode.id])
			break
		}

		d := distance[currentNode.id] + 1 // all edges have a weight of 1
		for _, n := range currentNode.neighbours {
			if !seen[n.id] && d < distance[n.id] {
				distance[n.id] = d
			}
		}
		seen[currentNode.id] = true
		currentNode = getNextNode(seen, distance, idToNode)
	}
}
