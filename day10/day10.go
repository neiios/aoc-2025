package day10

import (
	"aoc/lib"
	"math"
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

type Item struct {
	depth int
	state []bool
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

func SolvePart1(input string) (int, error) {
	machines, _ := parseInput(input)
	fewestButtonPresses := []int{}

	for _, machine := range machines {
		fewestButtonPresses = append(fewestButtonPresses, findMinimalButtonPresses(machine.diagram, machine.buttons))
	}

	return lib.Sum(fewestButtonPresses), nil
}

func SolvePart2(input string) (int, error) {
	machines, _ := parseInput(input)
	grandTotal := 0
	const magicNumber = 1e-4

	for _, m := range machines {
		rows, cols := len(m.joltages), len(m.buttons)

		A := make([][]float64, rows)
		b := make([]float64, rows)
		for i := range rows {
			A[i] = make([]float64, cols)
			b[i] = float64(m.joltages[i])
		}

		for j, btn := range m.buttons {
			for _, r := range btn {
				if r < rows {
					A[r][j] = 1.0
				}
			}
		}

		var err error
		A, b, err = lib.GaussianElimination(A, b)
		if err != nil {
			return 0, err
		}

		// find pivots
		pivotColFor := make([]int, rows)
		isFree := make([]bool, cols)
		for i := range isFree {
			isFree[i] = true
		}
		for r := range rows {
			pivotColFor[r] = -1
			for c := range cols {
				if math.Abs(A[r][c]) > magicNumber {
					pivotColFor[r] = c
					isFree[c] = false
					break
				}
			}
		}

		var freeVars []int
		for c, free := range isFree {
			if free {
				freeVars = append(freeVars, c)
			}
		}

		minPresses := math.MaxInt
		sol := make([]float64, cols)

		// brute force free variables
		var search func(int)
		search = func(idx int) {
			if idx == len(freeVars) {
				currentTotal := 0
				for _, fv := range freeVars {
					currentTotal += int(sol[fv])
				}

				if currentTotal >= minPresses {
					return
				}

				for r := rows - 1; r >= 0; r-- {
					pCol := pivotColFor[r]

					if pCol == -1 {
						if math.Abs(b[r]) > magicNumber {
							return
						}
						continue
					}

					sum := 0.0
					for c := range cols {
						if c > pCol {
							sum += A[r][c] * sol[c]
						}
					}

					val := (b[r] - sum) / A[r][pCol]

					if val < -magicNumber || math.Abs(val-math.Round(val)) > magicNumber {
						return
					}

					intVal := int(math.Round(val))
					sol[pCol] = float64(intVal)
					currentTotal += intVal
				}

				minPresses = min(minPresses, currentTotal)
				return
			}

			fv := freeVars[idx]
			for v := range 200 { // another magic number
				sol[fv] = float64(v)
				search(idx + 1)
			}
		}

		search(0)

		if minPresses != math.MaxInt {
			grandTotal += minPresses
		}
	}

	return grandTotal, nil
}
