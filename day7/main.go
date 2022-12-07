package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name    string
	parent  *dir
	subdirs map[string]*dir
	files   map[string]int
	size    int
}

func main() {
	part1()
	part2()
}

func part1() {

	// create a root dir
	root := &dir{
		name:    "/",
		subdirs: make(map[string]*dir),
		files:   make(map[string]int),
	}
	currentDir := root

	// read the input
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// build the filesystem tree
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)

		switch f[0] {
		case "$":
			switch f[1] {
			case "cd":
				switch f[2] {
				case "..":
					currentDir = currentDir.parent
				case "/":
					currentDir = root
				default:
					currentDir = currentDir.subdirs[f[2]]
				}
			case "ls":
				continue
			}
		case "dir":
			currentDir.subdirs[f[1]] = &dir{
				name:    f[1],
				parent:  currentDir,
				subdirs: make(map[string]*dir),
				files:   make(map[string]int),
			}
		default:
			fs, _ := strconv.Atoi(f[0])
			currentDir.files[f[1]] = fs
		}
	}

	// traverse the tree and calculate directory sizes
	size := dirSize(root)

	log.Printf("Part 1 - total size: %d\n", totalSize)

	freeSpace := 70000000 - size
	neededFreeSpace := 30000000 - freeSpace
	min := math.MaxInt
	for _, s := range dirSizes {
		if s >= neededFreeSpace {
			if s < min {
				min = s
			}
		}
	}
	log.Printf("Part 2 - smallest folder which will work: %d\n", min)
}

var totalSize = 0
var dirSizes = make([]int, 0)

func dirSize(d *dir) int {
	sizeOfFiles := 0
	size := 0
	for _, s := range d.files {
		sizeOfFiles += s
	}
	if len(d.subdirs) == 0 {
		size = sizeOfFiles
	} else {
		sizeOfDirs := 0
		for _, sd := range d.subdirs {
			sizeOfDirs += dirSize(sd)
		}
		size = sizeOfFiles + sizeOfDirs
	}
	if size <= 100000 {
		totalSize += size
	}
	dirSizes = append(dirSizes, size)
	d.size = size
	return size
}

func part2() {

}
