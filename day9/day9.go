package day9

import (
	"aoc/lib"
	"fmt"
	"strconv"
	"strings"
)

type Tile struct {
	X, Y int
}

func calcArea(a, b Tile) int {
	return (lib.Abs(a.X-b.X) + 1) * (lib.Abs(a.Y-b.Y) + 1)
}

func parseInput(input string) ([]Tile, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	tiles := []Tile{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, errX := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, errY := strconv.Atoi(strings.TrimSpace(parts[1]))
		if errX != nil || errY != nil {
			return nil, fmt.Errorf("coordinate parsing error: %v, %v", errX, errY)
		}
		tiles = append(tiles, Tile{x, y})
	}
	return tiles, nil
}

func SolvePart1(input string) (int, error) {
	tiles, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	maxArea := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			area := calcArea(tiles[i], tiles[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func SolvePart2(input string) (int, error) {
	tiles, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	n := len(tiles)
	maxArea := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isValidRectangle(tiles, i, j) {
				area := calcArea(tiles[i], tiles[j])
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea, nil
}

func isValidRectangle(tiles []Tile, i, j int) bool {
	n := len(tiles)
	r1, r2 := tiles[i], tiles[j]

	for k := 0; k < n; k++ {
		next := (k + 1) % n
		if k == i || k == j || next == i || next == j {
			continue
		}

		p1, p2 := tiles[k], tiles[next]
		if isLineInside(r1, r2, p1, p2) {
			return false
		}
	}
	return true
}

func isLineInside(r1, r2, p1, p2 Tile) bool {
	if isPointInside(r1, r2, p1) || isPointInside(r1, r2, p2) {
		return true
	}

	minX, maxX := min(r1.X, r2.X), max(r1.X, r2.X)
	minY, maxY := min(r1.Y, r2.Y), max(r1.Y, r2.Y)

	if p1.Y == p2.Y {
		yInside := p1.Y >= minY && p1.Y <= maxY
		spansX := (p1.X < minX && p2.X > maxX) || (p2.X < minX && p1.X > maxX)
		return yInside && spansX
	}

	if p1.X == p2.X {
		xInside := p1.X >= minX && p1.X <= maxX
		spansY := (p1.Y < minY && p2.Y > maxY) || (p2.Y < minY && p1.Y > maxY)
		return xInside && spansY
	}

	return false
}

func isPointInside(r1, r2, p Tile) bool {
	minX, maxX := min(r1.X, r2.X), max(r1.X, r2.X)
	minY, maxY := min(r1.Y, r2.Y), max(r1.Y, r2.Y)
	return p.X >= minX && p.X <= maxX && p.Y >= minY && p.Y <= maxY
}
