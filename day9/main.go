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

	hX := 0
	hY := 0
	tX := 0
	tY := 0
	visited := make(map[string]bool)
	visited["(0,0)"] = true

	for scanner.Scan() {
		t := scanner.Text()
		parts := strings.Fields(t)
		dir := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {

			// move the head according to the instructions
			switch dir {
			case "R":
				hX++
			case "U":
				hY++
			case "L":
				hX--
			case "D":
				hY--
			}

			// now move the tail
			if hX > tX+1 {
				tX++
				if hY > tY {
					tY++
				} else if hY < tY {
					tY--
				}
			} else if hX < tX-1 {
				tX--
				if hY > tY {
					tY++
				} else if hY < tY {
					tY--
				}
			} else if hY > tY+1 {
				tY++
				if hX > tX {
					tX++
				} else if hX < tX {
					tX--
				}
			} else if hY < tY-1 {
				tY--
				if hX > tX {
					tX++
				} else if hX < tX {
					tX--
				}
			}

			// store the tail's visitation history
			visited[fmt.Sprintf("(%d,%d)", tX, tY)] = true
		}
	}

	log.Printf("Tail visited %d positions\n", len(visited))
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	positions := make([]map[string]int, 10)
	for i := 0; i < 10; i++ {
		positions[i] = map[string]int{"x": 0, "y": 0}
	}
	visited := make(map[string]bool)
	visited["(0,0)"] = true

	for scanner.Scan() {
		t := scanner.Text()
		parts := strings.Fields(t)
		dir := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {

			// move the head according to the instructions
			switch dir {
			case "R":
				positions[0]["x"]++
			case "U":
				positions[0]["y"]++
			case "L":
				positions[0]["x"]--
			case "D":
				positions[0]["y"]--
			}

			// now move the rest of the "snake"
			for j := 1; j < 10; j++ {
				h := positions[j-1] // head is the knot in front
				t := positions[j]   // tail is the current knot
				if h["x"] > t["x"]+1 {
					t["x"]++
					if h["y"] > t["y"] {
						t["y"]++
					} else if h["y"] < t["y"] {
						t["y"]--
					}
				} else if h["x"] < t["x"]-1 {
					t["x"]--
					if h["y"] > t["y"] {
						t["y"]++
					} else if h["y"] < t["y"] {
						t["y"]--
					}
				} else if h["y"] > t["y"]+1 {
					t["y"]++
					if h["x"] > t["x"] {
						t["x"]++
					} else if h["x"] < t["x"] {
						t["x"]--
					}
				} else if h["y"] < t["y"]-1 {
					t["y"]--
					if h["x"] > t["x"] {
						t["x"]++
					} else if h["x"] < t["x"] {
						t["x"]--
					}
				}
			}

			// store the tail's visitation history
			visited[fmt.Sprintf("(%d,%d)", positions[9]["x"], positions[9]["y"])] = true
		}
	}

	log.Printf("Tail visited %d positions\n", len(visited))
}
