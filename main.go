package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"flag"
	"fmt"
	"os"
)

func main() {
	var dFlag = flag.Int("day", 0, "day to run")
	flag.Parse()

	switch *dFlag {
	case 1:
		content, _ := os.ReadFile("./input/day1.txt")
		part1, _ := day1.SolvePart1(string(content))
		part2, _ := day1.SolvePart2(string(content))
		fmt.Println(part1, part2)
	case 2:
		content, _ := os.ReadFile("./input/day2.txt")
		part1, _ := day2.SolvePart1(string(content))
		part2, _ := day2.SolvePart2(string(content))
		fmt.Println(part1, part2)
	case 3:
		content, _ := os.ReadFile("./input/day3.txt")
		part1, _ := day3.SolvePart1(string(content))
		part2, _ := day3.SolvePart2(string(content))
		fmt.Println(part1, part2)
	case 4:
		content, _ := os.ReadFile("./input/day4.txt")
		part1, _ := day4.SolvePart1(string(content))
		part2, _ := day4.SolvePart2(string(content))
		fmt.Println(part1, part2)
	default:
		fmt.Println("Invalid day:", *dFlag)
	}
}
