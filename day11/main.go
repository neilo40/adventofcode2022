package main

import (
	"log"

	"github.com/neilo40/adventofcode2022/helper"
)

type monkey struct {
	inspected int
	items     []int
	operation func(int) int
	test      int
	onTrue    int
	onFalse   int
}

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

func part1() {
	// capture problem starting condition
	monkeys := make([]monkey, 8)
	monkeys[0] = monkey{0, []int{64, 89, 65, 95}, func(old int) int { return old * 7 }, 3, 4, 1}
	monkeys[1] = monkey{0, []int{76, 66, 74, 87, 70, 56, 51, 66}, func(old int) int { return old + 5 }, 13, 7, 3}
	monkeys[2] = monkey{0, []int{91, 60, 63}, func(old int) int { return old * old }, 2, 6, 5}
	monkeys[3] = monkey{0, []int{92, 61, 79, 97, 79}, func(old int) int { return old + 6 }, 11, 2, 6}
	monkeys[4] = monkey{0, []int{93, 54}, func(old int) int { return old * 11 }, 5, 1, 7}
	monkeys[5] = monkey{0, []int{60, 79, 92, 69, 88, 82, 70}, func(old int) int { return old + 8 }, 17, 4, 0}
	monkeys[6] = monkey{0, []int{64, 57, 73, 89, 55, 53}, func(old int) int { return old + 1 }, 19, 0, 5}
	monkeys[7] = monkey{0, []int{62}, func(old int) int { return old + 4 }, 7, 3, 2}

	// iterate for 20 rounds
	for r := 0; r < 20; r++ {
		for m := 0; m < 8; m++ {
			for _, i := range monkeys[m].items {
				monkeys[m].inspected++
				worry := monkeys[m].operation(i)
				worry = worry / 3
				if worry%monkeys[m].test == 0 {
					monkeys[monkeys[m].onTrue].items = append(monkeys[monkeys[m].onTrue].items, worry)
				} else {
					monkeys[monkeys[m].onFalse].items = append(monkeys[monkeys[m].onFalse].items, worry)
				}
			}
			monkeys[m].items = make([]int, 0)
		}
	}

	// calculate monkey business level
	max := 0
	max2 := 0
	for _, m := range monkeys {
		if m.inspected > max {
			max2 = max
			max = m.inspected
		} else if m.inspected > max2 {
			max2 = m.inspected
		}
	}
	monkeyBusiness := max * max2
	log.Printf("Level of monkey business: %d\n", monkeyBusiness)

}

func part2() {
	// capture problem starting condition
	monkeys := make([]monkey, 8)
	monkeys[0] = monkey{0, []int{64, 89, 65, 95}, func(old int) int { return old * 7 }, 3, 4, 1}
	monkeys[1] = monkey{0, []int{76, 66, 74, 87, 70, 56, 51, 66}, func(old int) int { return old + 5 }, 13, 7, 3}
	monkeys[2] = monkey{0, []int{91, 60, 63}, func(old int) int { return old * old }, 2, 6, 5}
	monkeys[3] = monkey{0, []int{92, 61, 79, 97, 79}, func(old int) int { return old + 6 }, 11, 2, 6}
	monkeys[4] = monkey{0, []int{93, 54}, func(old int) int { return old * 11 }, 5, 1, 7}
	monkeys[5] = monkey{0, []int{60, 79, 92, 69, 88, 82, 70}, func(old int) int { return old + 8 }, 17, 4, 0}
	monkeys[6] = monkey{0, []int{64, 57, 73, 89, 55, 53}, func(old int) int { return old + 1 }, 19, 0, 5}
	monkeys[7] = monkey{0, []int{62}, func(old int) int { return old + 4 }, 7, 3, 2}

	// iterate for 10000 rounds
	for r := 0; r < 10000; r++ {
		for m := 0; m < 8; m++ {
			for _, i := range monkeys[m].items {
				monkeys[m].inspected++
				worry := monkeys[m].operation(i)

				// need to do something here to keep the numbers down, but not change the test outcomes
				// modular arithmetic.  I had to look this up.
				worry = worry % (2 * 7 * 13 * 3 * 19 * 5 * 17 * 11) // product of all test values

				if worry%monkeys[m].test == 0 {
					monkeys[monkeys[m].onTrue].items = append(monkeys[monkeys[m].onTrue].items, worry)
				} else {
					monkeys[monkeys[m].onFalse].items = append(monkeys[monkeys[m].onFalse].items, worry)
				}
			}
			monkeys[m].items = make([]int, 0)
		}
	}

	// calculate monkey business level
	max := 0
	max2 := 0
	for _, m := range monkeys {
		if m.inspected > max {
			max2 = max
			max = m.inspected
		} else if m.inspected > max2 {
			max2 = m.inspected
		}
	}
	monkeyBusiness := max * max2
	log.Printf("Level of monkey business: %d\n", monkeyBusiness)
}
