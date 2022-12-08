package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
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

	var forest = make([][]int, 99)
	// make a 2D grid with all the tree heights in it
	for r := 0; r < 99; r++ {
		scanner.Scan()
		row := scanner.Text()
		forest[r] = make([]int, 0, 99)
		for _, b := range row {
			h, _ := strconv.Atoi(string(b))
			forest[r] = append(forest[r], h)
		}
	}

	// go through each tree location and check if it is visible
	// no need to check the edges, so start with a count of 98 * 4 (grid is 99 on a side)
	visible := 98 * 4

	// need a grid to store whether a tree is visible.  trees can be visible from multiple
	// directions, but it should only be counted once
	seen := make([][]bool, 99)
	for i := 0; i < 99; i++ {
		seen[i] = make([]bool, 99)
	}

	// looking from left side
	for row := 1; row < 98; row++ {
		maxHeight := forest[row][0]
		for col := 1; col < 98; col++ {
			if forest[row][col] > maxHeight {
				maxHeight = forest[row][col]
				seen[row][col] = true
			}
		}
	}

	// looking from right
	for row := 1; row < 98; row++ {
		maxHeight := forest[row][98]
		for col := 97; col > 0; col-- {
			if forest[row][col] > maxHeight {
				maxHeight = forest[row][col]
				seen[row][col] = true
			}
		}
	}

	// looking from top
	for col := 1; col < 98; col++ {
		maxHeight := forest[0][col]
		for row := 1; row < 98; row++ {
			if forest[row][col] > maxHeight {
				maxHeight = forest[row][col]
				seen[row][col] = true
			}
		}
	}

	// looking from bottom
	for col := 1; col < 98; col++ {
		maxHeight := forest[98][col]
		for row := 97; row > 0; row-- {
			if forest[row][col] > maxHeight {
				maxHeight = forest[row][col]
				seen[row][col] = true
			}
		}
	}

	// get all visible trees
	for row := 1; row < 98; row++ {
		for col := 1; col < 98; col++ {
			if seen[row][col] {
				visible++
			}
		}
	}

	log.Printf("Visible trees: %d\n", visible)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var forest = make([][]int, 99)
	// make a 2D grid with all the tree heights in it
	for r := 0; r < 99; r++ {
		scanner.Scan()
		row := scanner.Text()
		forest[r] = make([]int, 0, 99)
		for _, b := range row {
			h, _ := strconv.Atoi(string(b))
			forest[r] = append(forest[r], h)
		}
	}

	maxScenicScore := 0
	// test every tree.  no need to consider the edges
	for row := 1; row < 98; row++ {
		for col := 1; col < 98; col++ {
			scenicScore := 1
			// get visible trees from here to top
			count := 0
			for nRow := row - 1; nRow >= 0; nRow-- {
				count++
				if forest[nRow][col] >= forest[row][col] {
					break
				}
			}
			scenicScore = scenicScore * count

			// from here to bottom
			count = 0
			for nRow := row + 1; nRow <= 98; nRow++ {
				count++
				if forest[nRow][col] >= forest[row][col] {
					break
				}
			}
			scenicScore = scenicScore * count

			// to right
			count = 0
			for nCol := col + 1; nCol <= 98; nCol++ {
				count++
				if forest[row][nCol] >= forest[row][col] {
					break
				}
			}
			scenicScore = scenicScore * count

			// to left
			count = 0
			for nCol := col - 1; nCol >= 0; nCol-- {
				count++
				if forest[row][nCol] >= forest[row][col] {
					break
				}
			}
			scenicScore = scenicScore * count

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	log.Printf("Max scenic score is %d\n", maxScenicScore)
}
