package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
	"aoc/day9"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

type PartSolver func(input string) (int, error)

type Day struct {
	Day        int
	SolvePart1 PartSolver
	SolvePart2 PartSolver
}

func main() {
	day := flag.Int("day", 0, "day to run")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

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
	case 5:
		solver = Day{5, day5.SolvePart1, day5.SolvePart2}
	case 6:
		solver = Day{6, day6.SolvePart1, day6.SolvePart2}
	case 7:
		solver = Day{7, day7.SolvePart1, day7.SolvePart2}
	case 8:
		solver = Day{8, day8.SolvePart1, day8.SolvePart2}
	case 9:
		solver = Day{9, day9.SolvePart1, day9.SolvePart2}
	default:
		log.Fatalf("invalid day %d", *day)
	}

	content, err := os.ReadFile(fmt.Sprintf("./input/day%d.txt", solver.Day))
	if err != nil {
		log.Fatalf("file not found: %v", err)
	}

	result1, err := solver.SolvePart1(string(content))
	if err != nil {
		log.Fatalf("solving part 1 failed: %v", err)
	}

	result2, err := solver.SolvePart2(string(content))
	if err != nil {
		log.Fatalf("solving part 2 failed: %v", err)
	}

	fmt.Println(result1, result2)
}
