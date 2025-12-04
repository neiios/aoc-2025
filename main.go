package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"flag"
	"fmt"
	"log"
	"os"
)

type PartSolver func(input string) (int, error)

type Day struct {
	Day        int
	SolvePart1 PartSolver
	SolvePart2 PartSolver
}

func main() {
	day := flag.Int("day", 0, "day to run")
	flag.Parse()

	var solver Day
	switch *day {
	case 1:
		solver = Day{1, day1.SolvePart1, day1.SolvePart2}
	case 2:
		solver = Day{2, day2.SolvePart1, day2.SolvePart2}
	case 3:
		solver = Day{3, day3.SolvePart1, day3.SolvePart2}
	case 4:
		solver = Day{4, day4.SolvePart1, day4.SolvePart2}
	default:
		log.Fatalf("invalid day %d", *day)
	}

	content, err := os.ReadFile(fmt.Sprintf("./input/day%d.txt", solver.Day))
	if err != nil {
		log.Fatalf("failed reading input: %v", err)
	}

	result1, err := solver.SolvePart1(string(content))
	if err != nil {
		log.Fatalf("failed reading input: %v", err)
	}

	result2, err := solver.SolvePart2(string(content))
	if err != nil {
		log.Fatalf("failed reading input: %v", err)
	}

	fmt.Println(result1, result2)
}
