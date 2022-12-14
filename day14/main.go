package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2022/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

type coord struct {
	col int
	row int
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// create the grid with all rock locations
	grid := make([][]bool, 200)
	for r := 0; r < 200; r++ {
		grid[r] = make([]bool, 1000)
	}

	deepestRock := 0
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)
		path := make([]coord, 0)
		for _, p := range f {
			if p == "->" {
				continue
			}
			coords := strings.Split(p, ",")
			col, _ := strconv.Atoi(coords[0])
			row, _ := strconv.Atoi(coords[1])
			if row > deepestRock {
				deepestRock = row
			}
			path = append(path, coord{col, row})
		}
		for i := 0; i < len(path)-1; i++ {
			start := path[i]
			end := path[i+1]
			if end.col > start.col {
				//l to r
				for col := start.col; col <= end.col; col++ {
					grid[start.row][col] = true
				}
			} else if end.col < start.col {
				//r to l
				for col := end.col; col <= start.col; col++ {
					grid[start.row][col] = true
				}
			} else if end.row > start.row {
				// down
				for row := start.row; row <= end.row; row++ {
					grid[row][start.col] = true
				}
			} else {
				// up, or single block?
				for row := end.row; row <= start.row; row++ {
					grid[row][start.col] = true
				}
			}
		}
	}

	// iterate the sand falling rules until sand y coord exceeds max rock y coord
	sandCount := 0
	sand := coord{row: 0, col: 500}
	for {
		if !grid[sand.row+1][sand.col] {
			// free space below
			sand.row++
			if sand.row > deepestRock {
				break
			}
			continue
		} else if !grid[sand.row+1][sand.col-1] {
			// below is blocked, free space to lower left
			sand.row++
			sand.col--
			if sand.row > deepestRock {
				break
			}
			continue
		} else if !grid[sand.row+1][sand.col+1] {
			// lower left is blocked, free space to lower right
			sand.row++
			sand.col++
			if sand.row > deepestRock {
				break
			}
			continue
		} else {
			// stopped
			sandCount++
			grid[sand.row][sand.col] = true
			sand.col = 500
			sand.row = 0
		}
	}

	log.Printf("Units of sand at rest: %d\n", sandCount)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// create the grid with all rock locations
	grid := make([][]bool, 200)
	for r := 0; r < 200; r++ {
		grid[r] = make([]bool, 1000)
	}

	deepestRock := 0
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)
		path := make([]coord, 0)
		for _, p := range f {
			if p == "->" {
				continue
			}
			coords := strings.Split(p, ",")
			col, _ := strconv.Atoi(coords[0])
			row, _ := strconv.Atoi(coords[1])
			if row > deepestRock {
				deepestRock = row
			}
			path = append(path, coord{col, row})
		}
		for i := 0; i < len(path)-1; i++ {
			start := path[i]
			end := path[i+1]
			if end.col > start.col {
				//l to r
				for col := start.col; col <= end.col; col++ {
					grid[start.row][col] = true
				}
			} else if end.col < start.col {
				//r to l
				for col := end.col; col <= start.col; col++ {
					grid[start.row][col] = true
				}
			} else if end.row > start.row {
				// down
				for row := start.row; row <= end.row; row++ {
					grid[row][start.col] = true
				}
			} else {
				// up, or single block?
				for row := end.row; row <= start.row; row++ {
					grid[row][start.col] = true
				}
			}
		}
	}

	// add a floor at deepestRock + 2
	for i := 0; i < 1000; i++ {
		grid[deepestRock+2][i] = true
	}

	// iterate the sand falling rules until sand reaches 500,0
	sandCount := 0
	sand := coord{row: 0, col: 500}
	for {
		if !grid[sand.row+1][sand.col] {
			// free space below
			sand.row++
			continue
		} else if !grid[sand.row+1][sand.col-1] {
			// below is blocked, free space to lower left
			sand.row++
			sand.col--
			continue
		} else if !grid[sand.row+1][sand.col+1] {
			// lower left is blocked, free space to lower right
			sand.row++
			sand.col++
			continue
		} else {
			// stopped
			sandCount++
			grid[sand.row][sand.col] = true
			if sand.col == 500 && sand.row == 0 {
				break
			}
			sand.col = 500
			sand.row = 0
		}
	}

	log.Printf("Part 2 : Units of sand at rest: %d\n", sandCount)
}
