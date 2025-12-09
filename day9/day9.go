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

type Rectangle struct {
	P1, P2 Tile
}

func (r Rectangle) Area() int {
	return (lib.Abs(r.P1.X-r.P2.X) + 1) * (lib.Abs(r.P1.Y-r.P2.Y) + 1)
}

func (r Rectangle) Contains(p Tile) bool {
	minX, maxX := min(r.P1.X, r.P2.X), max(r.P1.X, r.P2.X)
	minY, maxY := min(r.P1.Y, r.P2.Y), max(r.P1.Y, r.P2.Y)
	return p.X >= minX && p.X <= maxX && p.Y >= minY && p.Y <= maxY
}

// this was a pain to write...
func (r Rectangle) ContainsLine(p1, p2 Tile) bool {
	if r.Contains(p1) || r.Contains(p2) {
		return true
	}

	minX, maxX := min(r.P1.X, r.P2.X), max(r.P1.X, r.P2.X)
	minY, maxY := min(r.P1.Y, r.P2.Y), max(r.P1.Y, r.P2.Y)

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
			rect := Rectangle{tiles[i], tiles[j]}
			area := rect.Area()
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
			rect := Rectangle{tiles[i], tiles[j]}
			if isValidRectangle(rect, tiles, i, j) {
				area := rect.Area()
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea, nil
}

func isValidRectangle(rectangle Rectangle, tiles []Tile, i, j int) bool {
	n := len(tiles)

	for k := 0; k < n; k++ {
		next := (k + 1) % n
		p1, p2 := tiles[k], tiles[next]
		if k == i || k == j || next == i || next == j {
			continue
		}

		if rectangle.ContainsLine(p1, p2) {
			return false
		}
	}

	return true
}
