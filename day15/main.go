package main

import (
	"bufio"
	"fmt"
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
	filename := "input.txt"
	targetY := 2000000
	xOffset := 1000000

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	pairs := make([]pair, 0, 33)
	maxX := 0
	beaconsOnTargetRow := make(map[int]bool)
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
		if p.beacon.y == targetY {
			beaconsOnTargetRow[p.beacon.x] = true
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
				targetRow[xOffset+p.sensor.x+x] = true
				targetRow[xOffset+p.sensor.x-x] = true
			}
		}
	}

	illegalPos := 0
	illegalPos -= len(beaconsOnTargetRow)
	for _, v := range targetRow {
		if v {
			illegalPos++
		}
	}
	log.Printf("Part1: %d locations cannot contain a beacon on row %d\n", illegalPos, targetY)

}

var max int = 4000000

func part2() {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sensors := make([]pair, 0, 33)
	beacons := make(map[string]bool)
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
		beacons[fmt.Sprintf("%d_%d", coords[2], coords[3])] = true
		distance := math.Abs(float64(p.sensor.x-p.beacon.x)) + math.Abs(float64(p.sensor.y-p.beacon.y))
		p.distance = int(distance)
		sensors = append(sensors, p)
	}

	// only one space exists for the beacon, so it must be just outside the range of one of the sensors
	// compare points with the distance to all other sensors, if not within range of any, it's our location
	for _, s := range sensors {
		for dx := 0; dx < s.distance+2; dx++ {
			dy := (s.distance + 1) - dx
			for _, m := range [][]int{{-1, 1}, {1, -1}, {-1, -1}, {1, 1}} {
				x := s.sensor.x + (dx * m[0])
				y := s.sensor.y + (dy * m[1])
				_, isBeacon := beacons[fmt.Sprintf("%d_%d", x, y)]
				if isBeacon || x < 0 || x > max || y < 0 || y > max {
					continue
				}
				if possible(x, y, sensors) {
					log.Printf("Part2: tuning frequency is %d\n", (x*4000000)+y)
					return
				}
			}
		}
	}
}

func possible(x int, y int, sensors []pair) bool {
	for _, ns := range sensors {
		d := math.Abs(float64(x-ns.sensor.x)) + math.Abs(float64(y-ns.sensor.y))
		if d <= float64(ns.distance) {
			return false
		}
	}
	return true
}
