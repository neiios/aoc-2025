package day4

import (
	"bufio"
	"strings"
)

func parseInput(input string) (*Grid, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	rows := [][]*Cell{}
	rowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		row := []*Cell{}
		colIndex := 0
		for _, char := range chars {
			var cell *Cell
			if char == "@" {
				cell = &Cell{state: Occupied, row: rowIndex, col: colIndex}
			} else {
				cell = &Cell{state: Empty, row: rowIndex, col: colIndex}
			}
			row = append(row, cell)
			colIndex++
		}
		rows = append(rows, row)
		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Grid{
		rowLen: len(rows),
		colLen: len(rows[0]),
		rows:   rows,
	}, nil
}

func SolvePart1(input string) (int, error) {
	grid, _ := parseInput(input)
	totalMarked := removeAccesibleRolls(grid)
	return totalMarked, nil
}

func SolvePart2(input string) (int, error) {
	grid, _ := parseInput(input)
	totalMarked := 0

	for {
		marked := removeAccesibleRolls(grid)
		totalMarked += marked
		if marked == 0 {
			break
		}
	}

	return totalMarked, nil
}

func removeAccesibleRolls(grid *Grid) int {
	marked := 0

	for i := range grid.rows {
		row := grid.rows[i]
		for j := range row {
			isSurrounded := checkIsSurrounded(row[j], grid)
			if !isSurrounded && row[j].state == Occupied {
				row[j] = &Cell{state: Marked, col: row[j].col, row: row[j].row}
			}
		}
	}

	for i := range grid.rows {
		row := grid.rows[i]
		for j := range row {
			if row[j].state == Marked {
				marked++
				row[j] = &Cell{state: Empty, col: row[j].col, row: row[j].row}
			}
		}
	}

	return marked
}

func checkIsSurrounded(cell *Cell, grid *Grid) bool {
	var total = 0

	for _, offset := range offsets {
		nCol := cell.col + offset[0]
		nRow := cell.row + offset[1]
		if nCol < 0 || nRow < 0 {
			continue
		} else if nCol >= grid.colLen || nRow >= grid.rowLen {
			continue
		} else if grid.rows[nRow][nCol].state != Empty {
			total++
		} else {
			continue
		}
	}

	return total > 3
}

type CellState int

const (
	Empty CellState = iota
	Occupied
	Marked
)

type Grid struct {
	rowLen int
	colLen int
	rows   [][]*Cell
}

type Cell struct {
	state CellState
	row   int
	col   int
}

var offsets = [8][2]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}
