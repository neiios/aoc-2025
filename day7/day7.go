package day7

import (
	"aoc/lib"
	"strings"
)

func SolvePart1(input string) (int, error) {
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	rows = lib.RemoveEveryNRow(rows, 2)

	count := 0
	for i, row := range rows {
		if i == 0 {
			continue
		}

		for j, cell := range row {
			if cell == '^' {
				for k := i - 1; k >= 0; k-- {
					if rows[k][j] == 'S' || rows[k][j-1] == '^' || rows[k][j+1] == '^' {
						count++
						break
					} else if rows[k][j] == '^' {
						break
					}
				}
			}
		}
	}

	return count, nil
}

func SolvePart2(input string) (int, error) {
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	rows = lib.RemoveEveryNRow(rows, 2)

	timelineCount := 0
	type Key struct {
		i, j int
	}
	cache := map[Key]int{}

	lastRow := rows[len(rows)-1]
	for i := 0; i < len(lastRow); i++ {
		if lastRow[i] == '^' {
			cache[Key{len(rows) - 1, i}] = 2
		}
	}

	for i := len(rows) - 2; i >= 0; i-- {
		for j := len(rows[i]) - 1; j >= 0; j-- {
			if rows[i][j] == '^' {
				leftVal := 1
				for k := i + 1; k < len(rows); k++ {
					if rows[k][j-1] == '^' {
						leftVal = cache[Key{k, j - 1}]
						break
					}
				}

				rightVal := 1
				for k := i + 1; k < len(rows); k++ {
					if rows[k][j+1] == '^' {
						rightVal = cache[Key{k, j + 1}]
						break
					}
				}

				cache[Key{i, j}] = leftVal + rightVal
			}

			if rows[i][j] == 'S' {
				for k := i + 1; k < len(rows); k++ {
					if rows[k][j] == '^' {
						timelineCount = cache[Key{k, j}]
						break
					}
				}
			}
		}
	}

	return timelineCount, nil
}
