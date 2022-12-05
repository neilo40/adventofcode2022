package main

import (
	"bufio"
	"fmt"
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

	// create 9 empty stacks to hold crates
	stacks := make([][]string, 10)
	for i := 0; i < 10; i++ {
		stacks[i] = make([]string, 0)
	}

	// read the first 8 lines into a slice (the stack setup)
	setupLines := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		scanner.Scan()
		setupLines = append(setupLines, scanner.Text())
	}

	// working backwards, add the crates to each stack
	for row := 7; row >= 0; row-- {
		idx := 1
		for col := 1; col < 10; col++ {
			crate := string(setupLines[row][idx])
			if crate != " " {
				stacks[col] = append(stacks[col], crate)
			}
			idx += 4
		}
	}

	// skip the next two lines
	scanner.Scan()
	scanner.Scan()

	// the remainder of the file is the crate move instructions
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)
		count, _ := strconv.Atoi(f[1])
		from, _ := strconv.Atoi(f[3])
		to, _ := strconv.Atoi(f[5])
		for c := 0; c < count; c++ {
			l := len(stacks[from])
			stacks[to] = append(stacks[to], stacks[from][l-1])
			stacks[from] = stacks[from][:l-1]
		}
	}

	// print the crates on the top of each stack
	for i := 1; i < 10; i++ {
		l := len(stacks[i])
		fmt.Print(stacks[i][l-1])
	}
	fmt.Println()
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// create 9 empty stacks to hold crates
	stacks := make([][]string, 10)
	for i := 0; i < 10; i++ {
		stacks[i] = make([]string, 0)
	}

	// read the first 8 lines into a slice (the stack setup)
	setupLines := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		scanner.Scan()
		setupLines = append(setupLines, scanner.Text())
	}

	// working backwards, add the crates to each stack
	for row := 7; row >= 0; row-- {
		idx := 1
		for col := 1; col < 10; col++ {
			crate := string(setupLines[row][idx])
			if crate != " " {
				stacks[col] = append(stacks[col], crate)
			}
			idx += 4
		}
	}

	// skip the next two lines
	scanner.Scan()
	scanner.Scan()

	// the remainder of the file is the crate move instructions
	for scanner.Scan() {
		t := scanner.Text()
		f := strings.Fields(t)
		count, _ := strconv.Atoi(f[1])
		from, _ := strconv.Atoi(f[3])
		to, _ := strconv.Atoi(f[5])
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-count:]...)
		stacks[from] = stacks[from][:len(stacks[from])-count]
	}

	// print the crates on the top of each stack
	for i := 1; i < 10; i++ {
		l := len(stacks[i])
		fmt.Print(stacks[i][l-1])
	}
	fmt.Println()
}
