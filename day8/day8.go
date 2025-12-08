package day8

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

type JunctionBox struct {
	position  Position
	connected map[int]bool
	distances []float64
}

type DistancePair struct {
	distance float64
	index    int
}

func parseInput(input string) ([]JunctionBox, error) {
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	junctionBoxes := []JunctionBox{}
	for _, row := range rows {
		parts := strings.Split(row, ",")
		x, errX := strconv.Atoi(parts[0])
		y, errY := strconv.Atoi(parts[1])
		z, errZ := strconv.Atoi(parts[2])
		if errX != nil || errY != nil || errZ != nil {
			return nil, fmt.Errorf("%v, %v, %v", errX, errY, errZ)
		}
		junctionBoxes = append(junctionBoxes, JunctionBox{position: Position{x, y, z}, connected: map[int]bool{}, distances: make([]float64, len(rows))})
	}
	for _, junctionBox := range junctionBoxes {
		for i := range junctionBoxes {
			junctionBox.distances[i] = findEucDistance(junctionBox.position, junctionBoxes[i].position)
		}
	}
	return junctionBoxes, nil
}

func SolvePart1(input string) (int, error) {
	junctionBoxes, parseErr := parseInput(input)
	if parseErr != nil {
		return 0, parseErr
	}

	for range 1000 {
		minNotConnectedDistances := make([]DistancePair, len(junctionBoxes))
		for i := range len(minNotConnectedDistances) {
			minNotConnectedDistances[i] = DistancePair{math.Inf(1), -1}
		}

		for i, junctionBox := range junctionBoxes {
			for j, distance := range junctionBox.distances {
				if !junctionBox.connected[j] && distance != 0 && minNotConnectedDistances[i].distance > distance {
					minNotConnectedDistances[i] = DistancePair{distance, j}
				}
			}
		}

		minNotConnectedDistance := DistancePair{math.Inf(1), -1}
		minIndex := -1
		for i, pair := range minNotConnectedDistances {
			if minNotConnectedDistance.distance > pair.distance {
				minNotConnectedDistance = pair
				minIndex = i
			}
		}
		if minIndex == -1 {
			break
		}

		junctionBoxes[minIndex].connected[minNotConnectedDistance.index] = true
		junctionBoxes[minNotConnectedDistance.index].connected[minIndex] = true
	}

	visited := make([]bool, len(junctionBoxes))
	circuits := [][]int{}

	for i := range junctionBoxes {
		if !visited[i] {
			circuit := []int{}
			dfs(i, junctionBoxes, visited, &circuit)
			circuits = append(circuits, circuit)
		}
	}

	circuitSizes := make([]int, len(circuits))
	for i, circuit := range circuits {
		circuitSizes[i] = len(circuit)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))

	result := 1
	for i := 0; i < 3 && i < len(circuitSizes); i++ {
		result *= circuitSizes[i]
	}

	return result, nil
}

func SolvePart2(input string) (int, error) {
	junctionBoxes, parseErr := parseInput(input)
	if parseErr != nil {
		return 0, parseErr
	}

	unifyingConnection := []int{}
	for {
		minNotConnectedDistances := make([]DistancePair, len(junctionBoxes))
		for i := range len(minNotConnectedDistances) {
			minNotConnectedDistances[i] = DistancePair{math.Inf(1), -1}
		}

		for i, junctionBox := range junctionBoxes {
			for j, distance := range junctionBox.distances {
				if !junctionBox.connected[j] && distance != 0 && minNotConnectedDistances[i].distance > distance {
					minNotConnectedDistances[i] = DistancePair{distance, j}
				}
			}
		}

		minNotConnectedDistance := DistancePair{math.Inf(1), -1}
		minIndex := -1
		for i, pair := range minNotConnectedDistances {
			if minNotConnectedDistance.distance > pair.distance {
				minNotConnectedDistance = pair
				minIndex = i
			}
		}
		if minIndex == -1 {
			break
		}

		junctionBoxes[minIndex].connected[minNotConnectedDistance.index] = true
		junctionBoxes[minNotConnectedDistance.index].connected[minIndex] = true

		visited := make([]bool, len(junctionBoxes))
		circuits := [][]int{}
		for i := range junctionBoxes {
			if !visited[i] {
				circuit := []int{}
				dfs(i, junctionBoxes, visited, &circuit)
				circuits = append(circuits, circuit)
			}
		}
		if len(circuits) == 1 {
			unifyingConnection = []int{minIndex, minNotConnectedDistance.index}
			break
		}
	}

	return junctionBoxes[unifyingConnection[0]].position.x * junctionBoxes[unifyingConnection[1]].position.x, nil
}

func dfs(index int, junctionBoxes []JunctionBox, visited []bool, circuit *[]int) {
	visited[index] = true
	*circuit = append(*circuit, index)

	for connectedIndex := range junctionBoxes[index].connected {
		if !visited[connectedIndex] {
			dfs(connectedIndex, junctionBoxes, visited, circuit)
		}
	}
}

func findEucDistance(a Position, b Position) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}
