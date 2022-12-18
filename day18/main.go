package main

import (
	"bufio"
	"log"
	"math"
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

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	cubes := make(map[int]map[int]map[int]bool)
	for scanner.Scan() {
		t := scanner.Text()
		n := strings.Split(t, ",")
		x, _ := strconv.Atoi(n[0])
		_, ok := cubes[x]
		if !ok {
			cubes[x] = make(map[int]map[int]bool)
		}
		y, _ := strconv.Atoi(n[1])
		_, ok = cubes[x][y]
		if !ok {
			cubes[x][y] = make(map[int]bool)
		}
		z, _ := strconv.Atoi(n[2])
		cubes[x][y][z] = true
	}

	surface := 0
	for x := range cubes {
		for y := range cubes[x] {
			for z := range cubes[x][y] {
				neighbours := 0
				_, ok := cubes[x][y][z-1]
				if ok {
					neighbours++
				}
				_, ok = cubes[x][y][z+1]
				if ok {
					neighbours++
				}
				_, ok = cubes[x][y-1][z]
				if ok {
					neighbours++
				}
				_, ok = cubes[x][y+1][z]
				if ok {
					neighbours++
				}
				_, ok = cubes[x-1][y][z]
				if ok {
					neighbours++
				}
				_, ok = cubes[x+1][y][z]
				if ok {
					neighbours++
				}
				surface += (6 - neighbours)
			}
		}
	}

	log.Printf("Surface area: %d\n", surface)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	cubes := make(map[int]map[int]map[int]bool)
	for scanner.Scan() {
		t := scanner.Text()
		n := strings.Split(t, ",")
		x, _ := strconv.Atoi(n[0])
		_, ok := cubes[x]
		if !ok {
			cubes[x] = make(map[int]map[int]bool)
		}
		y, _ := strconv.Atoi(n[1])
		_, ok = cubes[x][y]
		if !ok {
			cubes[x][y] = make(map[int]bool)
		}
		z, _ := strconv.Atoi(n[2])
		cubes[x][y][z] = true
	}

	minX := math.MaxInt
	maxX := 0
	minY := math.MaxInt
	maxY := 0
	minZ := math.MaxInt
	maxZ := 0

	for x := range cubes {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		for y := range cubes[x] {
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
			for z := range cubes[x][y] {
				if z > maxZ {
					maxZ = z
				}
				if z < minZ {
					minZ = z
				}
			}
		}
	}

	seen := make(map[int]map[int]map[int]bool)
	for x := minX - 1; x <= maxX+1; x++ {
		seen[x] = make(map[int]map[int]bool)
		for y := minY - 1; y <= maxY+1; y++ {
			seen[x][y] = make(map[int]bool)
			for z := minZ - 1; z <= maxZ+1; z++ {
				seen[x][y][z] = false
			}
		}
	}

	// start outside the droplet
	startX := minX - 1
	startY := minY - 1
	startZ := minZ - 1

	log.Printf("Exterior surface area: %d\n", getSurfaces(cubes, seen, startX, startY, startZ))
}

func getSurfaces(cubes map[int]map[int]map[int]bool, seen map[int]map[int]map[int]bool, x int, y int, z int) int {
	surfaces := 0
	seen[x][y][z] = true
	// test each neighbour
	for _, coords := range [][]int{{x, y, z - 1}, {x, y, z + 1}, {x, y - 1, z}, {x, y + 1, z}, {x - 1, y, z}, {x + 1, y, z}} {
		nx := coords[0]
		ny := coords[1]
		nz := coords[2]
		// outside the scope of seen?  skip
		_, xok := seen[nx]
		_, yok := seen[nx][ny]
		_, zok := seen[nx][ny][nz]
		if !xok || !yok || !zok {
			continue
		}
		// if cube, increment surface count
		_, xok = cubes[nx]
		_, yok = cubes[nx][ny]
		_, zok = cubes[nx][ny][nz]
		if xok && yok && zok {
			surfaces++
		} else {
			// empty space, if not seen before, descend into it
			if !seen[nx][ny][nz] {
				surfaces += getSurfaces(cubes, seen, nx, ny, nz)
			}
		}

	}

	return surfaces
}
