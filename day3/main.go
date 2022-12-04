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

	sharedItems := make([]byte, 0, 300)
	for scanner.Scan() {
		t := scanner.Text()
		bag2Items := make(map[byte]bool)
		for i := len(t) / 2; i < len(t); i++ {
			bag2Items[t[i]] = true
		}

		for i := 0; i < len(t)/2; i++ {
			_, ok := bag2Items[t[i]]
			if ok {
				sharedItems = append(sharedItems, t[i])
				break
			}
		}
	}

	prioSum := 0
	for _, item := range sharedItems {
		prioSum += getPriority(item)
	}
	log.Printf("Sum of Priorities: %d\n", prioSum)
}

func getPriority(item byte) int {
	if item > 96 {
		// lowercase
		return int(item - 96)
	} else {
		// uppercase
		return int(item-64) + 26
	}
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	badges := make([]byte, 0, 100)
	for g := 0; g < 300; g += 3 {
		scanner.Scan()
		g1 := scanner.Text()
		scanner.Scan()
		g2 := scanner.Text()
		scanner.Scan()
		g3 := scanner.Text()

		commonItems := make([]rune, 0, 20)
		for _, i := range g1 {
			for _, j := range g2 {
				if i == j {
					commonItems = append(commonItems, i)
				}
			}
		}

		found := false
		for _, k := range commonItems {
			if found {
				break
			}
			for _, l := range g3 {
				if k == l {
					badges = append(badges, byte(k))
					found = true
					break
				}
			}
		}
	}

	prioSum := 0
	for _, item := range badges {
		prioSum += getPriority(item)
	}
	log.Printf("Sum of Badge Priorities: %d\n", prioSum)
}
