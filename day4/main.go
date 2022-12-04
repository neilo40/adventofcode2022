package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

	overlaps := 0
	for scanner.Scan() {
		t := scanner.Text()
		elfRanges := strings.Split(t, ",")

		elf1 := strings.Split(elfRanges[0], "-")
		elf1Start, _ := strconv.Atoi(elf1[0])
		elf1End, _ := strconv.Atoi(elf1[1])

		elf2 := strings.Split(elfRanges[1], "-")
		elf2Start, _ := strconv.Atoi(elf2[0])
		elf2End, _ := strconv.Atoi(elf2[1])

		if elf1Start <= elf2Start && elf1End >= elf2End {
			// elf1 contains elf2
			overlaps++
		} else if elf2Start <= elf1Start && elf2End >= elf1End {
			// elf2 contains elf1
			overlaps++
		}
	}

	log.Printf("Full overlaps: %d\n", overlaps)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	overlaps := 0
	for scanner.Scan() {
		t := scanner.Text()
		elfRanges := strings.Split(t, ",")

		elf1 := strings.Split(elfRanges[0], "-")
		elf1Start, _ := strconv.Atoi(elf1[0])
		elf1End, _ := strconv.Atoi(elf1[1])

		elf2 := strings.Split(elfRanges[1], "-")
		elf2Start, _ := strconv.Atoi(elf2[0])
		elf2End, _ := strconv.Atoi(elf2[1])

		if elf1Start <= elf2Start && elf1End >= elf2Start {
			overlaps++
		} else if elf1Start > elf2Start && elf1Start <= elf2End {
			overlaps++
		}
	}

	log.Printf("Overlaps: %d\n", overlaps)
}
