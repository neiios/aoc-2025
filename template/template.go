package template

import (
	"aoc/lib"
	"bufio"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	nums := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nums, nil
}

func SolvePart1(input string) (int, error) {
	nums, _ := parseInput(input)
	return lib.Sum(nums), nil
}

func SolvePart2(input string) (int, error) {
	nums, _ := parseInput(input)
	return lib.Sum(nums) * 2, nil
}
