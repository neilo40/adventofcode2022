package main

import (
	"bufio"
	"fmt"
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

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	x := 1
	cycle := 1
	signalStrength := 0
	targetCycle := 20
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)
		newX := x
		if f[0] == "noop" {
			cycle++
		} else {
			num, _ := strconv.Atoi(f[1])
			newX += num
			cycle += 2
		}
		// did we cross the target?
		if cycle == targetCycle {
			signalStrength += (newX * targetCycle)
			targetCycle += 40
		} else if cycle == targetCycle+1 {
			signalStrength += (x * targetCycle)
			targetCycle += 40
		}
		x = newX
	}

	log.Printf("Signal Strength: %d\n", signalStrength)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	x := 1
	cycle := 0
	screen := make([]rune, 240)
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)

		// noop and addx both take at least one cycle
		screen[cycle] = getPixel(x, cycle)
		cycle++

		// if addx, do it and take another cycle
		if f[0] == "addx" {
			screen[cycle] = getPixel(x, cycle)
			num, _ := strconv.Atoi(f[1])
			x += num
			cycle++
		}
	}

	for i := 0; i < 240; i++ {
		if i > 0 && i%40 == 0 {
			fmt.Println()
		}
		fmt.Printf("%s", string(screen[i]))
	}
}

func getPixel(x int, cycle int) rune {
	// cycle runs from 0 to 239, but x is from 1 to 38
	beam := cycle - 40*(cycle/40)

	if beam >= x-1 && beam <= x+1 {
		return '#'
	}
	return ' '
}
