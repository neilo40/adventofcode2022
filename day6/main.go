package main

import (
	"bufio"
	"log"
	"os"
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

	scanner.Scan()
	t := scanner.Text()
	for i := 0; i < len(t)-4; i++ {
		freq := make(map[byte]bool)
		dupeFound := false
		for j := 0; j < 4; j++ {
			_, ok := freq[t[i+j]]
			if ok {
				dupeFound = true
				break
			} else {
				freq[t[i+j]] = true
			}
		}
		if !dupeFound {
			log.Printf("Start of packet marker at %d\n", i+4)
			break
		}
	}
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	t := scanner.Text()
	for i := 0; i < len(t)-14; i++ {
		freq := make(map[byte]bool)
		dupeFound := false
		for j := 0; j < 14; j++ {
			_, ok := freq[t[i+j]]
			if ok {
				dupeFound = true
				break
			} else {
				freq[t[i+j]] = true
			}
		}
		if !dupeFound {
			log.Printf("Start of message marker at %d\n", i+14)
			break
		}
	}
}
