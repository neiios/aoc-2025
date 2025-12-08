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
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	junctionBoxes := []JunctionBox{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid input format: %s", line)
		}
		x, errX := strconv.Atoi(parts[0])
		y, errY := strconv.Atoi(parts[1])
		z, errZ := strconv.Atoi(parts[2])
		if errX != nil || errY != nil || errZ != nil {
			return nil, fmt.Errorf("coordinate parsing error: %v, %v, %v", errX, errY, errZ)
		}
		junctionBoxes = append(junctionBoxes, JunctionBox{position: Position{x, y, z}, connected: map[int]bool{}, distances: make([]float64, len(lines))})
	}
	for i := range junctionBoxes {
		for j := range junctionBoxes {
			junctionBoxes[i].distances[j] = calculateEuclideanDistance(junctionBoxes[i].position, junctionBoxes[j].position)
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
		closestUnconnectedNodes := make([]DistancePair, len(junctionBoxes))
		for i := range len(closestUnconnectedNodes) {
			closestUnconnectedNodes[i] = DistancePair{math.Inf(1), -1}
		}

		for i, junctionBox := range junctionBoxes {
			for j, distance := range junctionBox.distances {
				if i != j && !junctionBox.connected[j] && closestUnconnectedNodes[i].distance > distance {
					closestUnconnectedNodes[i] = DistancePair{distance, j}
				}
			}
		}

		minConnectionCandidate := DistancePair{math.Inf(1), -1}
		sourceNodeIndex := -1
		for i, pair := range closestUnconnectedNodes {
			if minConnectionCandidate.distance > pair.distance {
				minConnectionCandidate = pair
				sourceNodeIndex = i
			}
		}
		if sourceNodeIndex == -1 {
			break
		}

		targetNodeIndex := minConnectionCandidate.index
		junctionBoxes[sourceNodeIndex].connected[targetNodeIndex] = true
		junctionBoxes[targetNodeIndex].connected[sourceNodeIndex] = true
	}

	visited := make([]bool, len(junctionBoxes))
	circuits := [][]int{}

	for i := range junctionBoxes {
		if !visited[i] {
			var currentCircuit []int
			dfs(i, junctionBoxes, visited, &currentCircuit)
			circuits = append(circuits, currentCircuit)
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

	var finalConnection []int
	for {
		closestUnconnectedNodes := make([]DistancePair, len(junctionBoxes))
		for i := range len(closestUnconnectedNodes) {
			closestUnconnectedNodes[i] = DistancePair{math.Inf(1), -1}
		}

		for i, junctionBox := range junctionBoxes {
			for j, distance := range junctionBox.distances {
				if i != j && !junctionBox.connected[j] && closestUnconnectedNodes[i].distance > distance {
					closestUnconnectedNodes[i] = DistancePair{distance, j}
				}
			}
		}

		minConnectionCandidate := DistancePair{math.Inf(1), -1}
		sourceNodeIndex := -1
		for i, pair := range closestUnconnectedNodes {
			if minConnectionCandidate.distance > pair.distance {
				minConnectionCandidate = pair
				sourceNodeIndex = i
			}
		}
		if sourceNodeIndex == -1 {
			break
		}

		targetNodeIndex := minConnectionCandidate.index
		junctionBoxes[sourceNodeIndex].connected[targetNodeIndex] = true
		junctionBoxes[targetNodeIndex].connected[sourceNodeIndex] = true

		visited := make([]bool, len(junctionBoxes))
		circuits := [][]int{}
		for i := range junctionBoxes {
			if !visited[i] {
				var currentCircuit []int
				dfs(i, junctionBoxes, visited, &currentCircuit)
				circuits = append(circuits, currentCircuit)
			}
		}
		if len(circuits) == 1 {
			finalConnection = []int{sourceNodeIndex, targetNodeIndex}
			break
		}
	}

	return junctionBoxes[finalConnection[0]].position.x * junctionBoxes[finalConnection[1]].position.x, nil
}

func dfs(nodeIndex int, junctionBoxes []JunctionBox, visited []bool, circuit *[]int) {
	visited[nodeIndex] = true
	*circuit = append(*circuit, nodeIndex)

	for connectedIndex := range junctionBoxes[nodeIndex].connected {
		if !visited[connectedIndex] {
			dfs(connectedIndex, junctionBoxes, visited, circuit)
		}
	}
}

func calculateEuclideanDistance(p1 Position, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2))
}
