package day4

import (
	"bufio"
	"strings"
)

func parseInput(input string) (Grid, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	rows := [][]Cell{}
	rowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := []Cell{}
		for colIndex, char := range line {
			state := Empty
			if char == '@' {
				state = Occupied
			}
			row = append(row, Cell{state: state, row: rowIndex, col: colIndex})
		}
		rows = append(rows, row)
		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		return Grid{}, err
	}

	return Grid{
		rowLen: len(rows),
		colLen: len(rows[0]),
		rows:   rows,
	}, nil
}

func SolvePart1(input string) (int, error) {
	grid, _ := parseInput(input)
	totalMarked := removeAccesibleRolls(&grid)
	return totalMarked, nil
}

func SolvePart2(input string) (int, error) {
	grid, _ := parseInput(input)
	totalMarked := 0

	for {
		marked := removeAccesibleRolls(&grid)
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
		for j := range grid.rows[i] {
			if !checkIsSurrounded(grid.rows[i][j], grid) && grid.rows[i][j].state == Occupied {
				grid.rows[i][j].state = Marked
			}
		}
	}

	for i := range grid.rows {
		for j := range grid.rows[i] {
			if grid.rows[i][j].state == Marked {
				marked++
				grid.rows[i][j].state = Empty
			}
		}
	}

	return marked
}

func checkIsSurrounded(cell Cell, grid *Grid) bool {
	total := 0

	for _, offset := range offsets {
		nCol := cell.col + offset[0]
		nRow := cell.row + offset[1]
		if nCol < 0 || nRow < 0 || nCol >= grid.colLen || nRow >= grid.rowLen {
			continue
		}
		if grid.rows[nRow][nCol].state != Empty {
			total++
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
	rows   [][]Cell
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
