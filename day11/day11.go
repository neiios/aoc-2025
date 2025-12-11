package day11

import (
	"strings"
)

func parseInput(input string) (map[string][]string, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	devices := map[string][]string{}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		connections := strings.Split(parts[1], " ")
		devices[parts[0]] = connections
	}

	return devices, nil
}

func SolvePart1(input string) (int, error) {
	devices, _ := parseInput(input)

	memo := map[string]int{}
	pathCount := dfs(memo, devices, "you", "out")

	return pathCount, nil
}

func SolvePart2(input string) (int, error) {
	devices, _ := parseInput(input)

	svrToDac := dfs(map[string]int{}, devices, "svr", "dac")
	dacToFft := dfs(map[string]int{}, devices, "dac", "fft")
	fftToOut := dfs(map[string]int{}, devices, "fft", "dac")

	svrToFft := dfs(map[string]int{}, devices, "svr", "fft")
	fftToDac := dfs(map[string]int{}, devices, "fft", "dac")
	dacToOut := dfs(map[string]int{}, devices, "dac", "out")

	return svrToDac*dacToFft*fftToOut + svrToFft*fftToDac*dacToOut, nil
}

func dfs(memo map[string]int, devices map[string][]string, current string, target string) int {
	if current == target {
		return 1
	}

	if count, ok := memo[current]; ok {
		return count
	}

	total := 0
	connections := devices[current]
	for _, connection := range connections {
		total += dfs(memo, devices, connection, target)
		memo[current] = total
	}

	return total
}
