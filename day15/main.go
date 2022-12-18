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

type coord struct {
	x int
	y int
}
type pair struct {
	sensor   coord
	beacon   coord
	distance int // manhattan distance between sensor and beacon
}

func part1() {
	filename := "example.txt"
	targetY := 10  // 2000000
	xOffset := 100 // 1000000

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	pairs := make([]pair, 0, 33)
	maxX := 0
	for scanner.Scan() {
		t := scanner.Text()
		fields := strings.Fields(t)
		coords := make([]int, 0, 4)
		for _, f := range fields {
			if strings.Contains(f, "=") {
				n := strings.TrimRight(strings.TrimLeft(f, "xy="), ",:")
				num, _ := strconv.Atoi(n)
				coords = append(coords, num)
			}
		}
		p := pair{
			sensor: coord{x: coords[0], y: coords[1]},
			beacon: coord{x: coords[2], y: coords[3]},
		}
		distance := math.Abs(float64(p.sensor.x-p.beacon.x)) + math.Abs(float64(p.sensor.y-p.beacon.y))
		p.distance = int(distance)
		if (p.sensor.x + p.distance) > maxX {
			maxX = p.sensor.x + p.distance
		}
		pairs = append(pairs, p)
	}

	targetRow := make([]bool, maxX+xOffset) // avoid negative slice indices
	for _, p := range pairs {
		// is p within range of target?  only consider sensors that have a chance of extending to the target row
		deltaY := int(math.Abs(float64(targetY - p.sensor.y)))
		if deltaY <= p.distance {
			// how many spaces left after subtracting delta-y from distance?
			remaining := p.distance - deltaY
			for x := 0; x <= remaining; x++ {
				if p.beacon.x > p.sensor.x {
					// beacon is to the right
					targetRow[xOffset+p.sensor.x+x] = true
				} else {
					// beacon is to the left (or at same x coord)
					targetRow[xOffset+p.sensor.x-x] = true
				}
			}
		}
	}

	illegalPos := 0
	for _, v := range targetRow {
		if v {
			illegalPos++
		}
	}
	log.Printf("Part1: %d locations cannot contain a beacon on row %d\n", illegalPos, targetY)
	// 3595343 too low

}

func part2() {
}
