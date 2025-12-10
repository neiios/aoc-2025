package day10

import (
	"aoc/lib"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	diagram  []bool
	buttons  [][]int
	joltages []int
}

func parseInput(input string) ([]Machine, error) {
	machines := []Machine{}
	lineExtractor, _ := regexp.Compile(`\[(.+)\] (\(.+\)) {(.+)}`)
	buttonExtractor, _ := regexp.Compile(`\(([^)]+)\)`)
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		machine := Machine{}
		parts := lineExtractor.FindStringSubmatch(line)

		diagramStrParts := strings.SplitSeq(parts[1], "")
		for part := range diagramStrParts {
			switch part {
			case ".":
				machine.diagram = append(machine.diagram, false)
			case "#":
				machine.diagram = append(machine.diagram, true)
			}
		}

		buttons := buttonExtractor.FindAllStringSubmatch(parts[2], -1)
		for _, buttonMatch := range buttons {
			button := []int{}
			buttonStr := buttonMatch[1]
			buttonParts := strings.SplitSeq(buttonStr, ",")
			for indexStr := range buttonParts {
				index, _ := strconv.Atoi(strings.TrimSpace(indexStr))
				button = append(button, index)
			}
			machine.buttons = append(machine.buttons, button)
		}

		joltagesStrParts := strings.SplitSeq(parts[3], ",")
		for joltageStr := range joltagesStrParts {
			joltage, _ := strconv.Atoi(joltageStr)
			machine.joltages = append(machine.joltages, joltage)
		}

		machines = append(machines, machine)
	}

	return machines, nil
}

func SolvePart1(input string) (int, error) {
	machines, _ := parseInput(input)
	fewestButtonPresses := []int{}

	for _, machine := range machines {
		fewestButtonPresses = append(fewestButtonPresses, findMinimalButtonPresses(machine.diagram, machine.buttons))
	}

	return lib.Sum(fewestButtonPresses), nil
}

func findMinimalButtonPresses(diagram []bool, buttons [][]int) int {
	queue := lib.Queue[Item]{}
	queue.Push(Item{depth: 0, state: make([]bool, len(diagram))})
	for {
		current, _ := queue.Pop()
		if slices.Equal(current.state, diagram) {
			return current.depth
		}

		for _, button := range buttons {
			queue.Push(Item{depth: current.depth + 1, state: pressButton(current.state, button)})
		}
	}
}

func pressButton(state []bool, button []int) []bool {
	newState := make([]bool, len(state))
	copy(newState, state)
	for _, i := range button {
		newState[i] = !newState[i]
	}
	return newState
}

func SolvePart2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	return 0, nil
}

type Item struct {
	depth int
	state []bool
}
