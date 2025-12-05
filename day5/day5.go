package day5

import (
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
}

func parseInput(input string) ([]Range, []int, error) {
	parts := strings.Split(input, "\n\n")

	rangeLines := strings.Split(parts[0], "\n")
	ranges := []Range{}

	for _, line := range rangeLines {
		rangeParts := strings.Split(line, "-")
		from, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			return nil, nil, err
		}
		to, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			return nil, nil, err
		}
		ranges = append(ranges, Range{from: from, to: to})
	}

	numLines := strings.Split(parts[1], "\n")
	nums := []int{}

	for _, line := range numLines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, nil, err
		}
		nums = append(nums, num)
	}

	return ranges, nums, nil
}

func SolvePart1(input string) (int, error) {
	ranges, nums, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	totalFresh := 0
	for _, num := range nums {
		for _, r := range ranges {
			if num >= r.from && num <= r.to {
				totalFresh++
				break
			}
		}
	}

	return totalFresh, nil
}

func SolvePart2(input string) (int, error) {
	ranges, _, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].from < ranges[j].from
	})

	merged := []Range{ranges[0]}

	for _, candidate := range ranges[1:] {
		for _, saved := range merged {
			if saved.to >= candidate.from && candidate.to > saved.to {
				candidate.from = saved.to + 1
			}
		}

		isContained := false
		for _, saved := range merged {
			if saved.from <= candidate.from && saved.to >= candidate.to {
				isContained = true
				break
			}
		}

		if !isContained {
			merged = append(merged, candidate)
		}
	}

	totalIds := 0
	for _, r := range merged {
		totalIds += r.to - r.from + 1
	}

	return totalIds, nil
}
