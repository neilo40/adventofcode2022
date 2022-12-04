package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var scoring = map[string]map[string]int{
	"A": { //rock
		"X": 1 + 3, // rock - draw
		"Y": 2 + 6, // paper - win
		"Z": 3,     // scissors - lose
	},
	"B": { // paper
		"X": 1,
		"Y": 2 + 3,
		"Z": 3 + 6,
	},
	"C": { // scissors
		"X": 1 + 6,
		"Y": 2,
		"Z": 3 + 3,
	},
}

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

	score := 0
	for scanner.Scan() {
		t := scanner.Text()
		plays := strings.Fields(t)
		score += scoring[plays[0]][plays[1]]
	}
	log.Printf("Score: %d\n", score)
}

var scoring2 = map[string]map[string]int{
	"A": { //rock
		"X": 3,
		"Y": 1 + 3,
		"Z": 2 + 6,
	},
	"B": { // paper
		"X": 1,
		"Y": 2 + 3,
		"Z": 3 + 6,
	},
	"C": { // scissors
		"X": 2,
		"Y": 3 + 3,
		"Z": 1 + 6,
	},
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	score := 0
	for scanner.Scan() {
		t := scanner.Text()
		plays := strings.Fields(t)
		score += scoring2[plays[0]][plays[1]]
	}
	log.Printf("Score (part 2): %d\n", score)
}
