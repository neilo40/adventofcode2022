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

	elfCalories := 0
	maxCalories := 0
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			if elfCalories > maxCalories {
				maxCalories = elfCalories
			}
			elfCalories = 0
		} else {
			calories, _ := strconv.Atoi(t)
			elfCalories += calories
		}
	}
	log.Printf("Max Calories : %d\n", maxCalories)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	elfCalories := 0
	maxCalories := []int{0, 0, 0}
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			maxCalories = topThree(elfCalories, maxCalories)
			elfCalories = 0
		} else {
			calories, _ := strconv.Atoi(t)
			elfCalories += calories
		}
	}
	log.Printf("Sum of top three Calories : %d\n", maxCalories[0]+maxCalories[1]+maxCalories[2])
}

func topThree(val int, max []int) []int {
	newMax := []int{max[0], max[1], max[2]}
	// max[0] is the highest
	if val > max[2] {
		if val > max[1] {
			if val > max[0] {
				newMax[2] = max[1]
				newMax[1] = max[0]
				newMax[0] = val
			} else {
				newMax[2] = max[1]
				newMax[1] = val
			}
		} else {
			newMax[2] = val
		}
	}
	return newMax
}
