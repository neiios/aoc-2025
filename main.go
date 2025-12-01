package main

import (
	"flag"
	"fmt"
)

func main() {
	var dFlag = flag.Int("day", 0, "day to run")
	flag.Parse()

	switch *dFlag {
	case 1:
		part1, _ := RunDay1Part1("./input/day1.txt")
		part2, _ := RunDay1Part2("./input/day1.txt")
		fmt.Println(part1, part2)
	default:
		fmt.Println("Invalid day:", *dFlag)
	}
}
