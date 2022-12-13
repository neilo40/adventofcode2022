package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

	pair := 1
	sum := 0
	for scanner.Scan() {
		p1 := scanner.Text()
		scanner.Scan()
		p2 := scanner.Text()
		scanner.Scan() // blank line
		// remove enclosing brackets
		p1 = strings.TrimSuffix(strings.TrimPrefix(p1, "["), "]")
		p2 = strings.TrimSuffix(strings.TrimPrefix(p2, "["), "]")
		if compare(p1, p2) == -1 {
			sum += pair
		}
		pair++
	}
	log.Printf("Sum of pair indices: %d\n", sum)
}

type bySize []string

func (b bySize) Len() int           { return len(b) }
func (b bySize) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b bySize) Less(i, j int) bool { return compare(b[i], b[j]) == -1 }

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	packets := make([]string, 0, 150)
	for scanner.Scan() {
		t := scanner.Text()
		if t != "" {
			packets = append(packets, strings.TrimSuffix(strings.TrimPrefix(t, "["), "]"))
		}
	}
	packets = append(packets, "[2]")
	packets = append(packets, "[6]")

	sort.Sort(bySize(packets))

	idx2 := 0
	idx6 := 0
	for i, p := range packets {
		if p == "[2]" {
			idx2 = i + 1
		}
		if p == "[6]" {
			idx6 = i + 1
		}
	}

	log.Printf("Decoder Key: %d\n", idx2*idx6)
}

// -1 left is smaller
// 0 same
// 1 right is smaller
func compare(p1 string, p2 string) int {

	// get the elements of the list
	l := splitPacket(p1)
	r := splitPacket(p2)

	if l[0] == "" && r[0] != "" {
		return -1
	} else if r[0] == "" && l[0] != "" {
		return 1
	}

	for i := 0; i < len(l); i++ {
		res := 0
		if i >= len(r) {
			// we got to the end of r, r is smaller
			return 1
		} else if strings.HasPrefix(l[i], "[") && !strings.HasPrefix(r[i], "[") {
			// left is array, right is int
			res = compare(strings.TrimSuffix(strings.TrimPrefix(l[i], "["), "]"), r[i])
		} else if strings.HasPrefix(r[i], "[") && !strings.HasPrefix(l[i], "[") {
			// right is array, left is int
			res = compare(l[i], strings.TrimSuffix(strings.TrimPrefix(r[i], "["), "]"))
		} else if strings.HasPrefix(l[i], "[") && strings.HasPrefix(r[i], "[") {
			// both are arrays
			res = compare(strings.TrimSuffix(strings.TrimPrefix(l[i], "["), "]"), strings.TrimSuffix(strings.TrimPrefix(r[i], "["), "]"))
		} else {
			// both are ints
			lNum, _ := strconv.Atoi(l[i])
			rNum, _ := strconv.Atoi(r[i])
			if lNum < rNum {
				return -1
			} else if lNum > rNum {
				return 1
			}
		}
		if res != 0 {
			return res
		}

		if i == len(l)-1 && i < len(r)-1 {
			// we got to the end of l, l is smaller
			return -1
		}
	}

	return 0
}

func splitPacket(p string) []string {
	depth := 0
	parts := make([]string, 0)
	part := ""
	for _, r := range p {
		switch r {
		case '[':
			depth++
			part += string(r)
		case ']':
			depth--
			part += string(r)
		case ',':
			if depth == 0 {
				// split
				parts = append(parts, part)
				part = ""
			} else {
				part += string(r)
			}
		default:
			part += string(r)
		}
	}
	parts = append(parts, part) // p will never end in , so this call is needed at the end
	return parts
}
