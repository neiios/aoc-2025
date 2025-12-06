package day6

import (
	"aoc/lib"
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Operation int

const (
	Add Operation = iota
	Multiply
)

type Problem struct {
	numbers []int
	op      Operation
}

func parseInputPart1(input string) ([]Problem, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	grid := [][]int{}
	ops := []string{}
	problems := []Problem{}
	regex, _ := regexp.Compile(`\d+`)

	for scanner.Scan() {
		line := scanner.Text()
		nums := regex.FindAllString(line, -1)

		if len(nums) == 0 {
			regex, _ := regexp.Compile("[*+]")
			ops = regex.FindAllString(line, -1)
			continue
		}

		row := []int{}
		for _, num := range nums {
			numInt, _ := strconv.Atoi(num)
			row = append(row, numInt)
		}

		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	tGrid := lib.TransposeMatrix(grid)
	for i, r := range tGrid {
		var op Operation
		if ops[i] == "+" {
			op = Add
		}
		if ops[i] == "*" {
			op = Multiply
		}

		problems = append(problems, Problem{
			numbers: r,
			op:      op,
		})
	}

	return problems, nil
}

func parseInputPart2(input string) ([]Problem, error) {
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	cols := [][]string{}
	opRow := rows[len(rows)-1]
	colLengths := []int{}
	r, _ := regexp.Compile(`[*+]\s+`)
	ops := r.FindAllString(opRow, -1)
	problems := []Problem{}

	for _, op := range ops {
		colLengths = append(colLengths, len(op)-1)
	}
	colLengths[len(ops)-1] += 1

	for _, colLength := range colLengths {
		col := []string{}
		for i, row := range rows {
			col = append(col, row[:colLength])

			rows[i] = row[colLength:]
			if len(row[colLength:]) > 0 {
				rows[i] = rows[i][1:]
			}
		}
		cols = append(cols, col)
	}

	for _, col := range cols {
		var op Operation
		if col[len(col)-1][0] == '+' {
			op = Add
		}
		if col[len(col)-1][0] == '*' {
			op = Multiply
		}

		numbers := []int{}
		for i := len(col[len(col)-1]) - 1; i >= 0; i-- {
			var n string
			for _, num := range col[:len(col)-1] {
				n += string(num[i])
			}
			nInt, _ := strconv.Atoi(strings.ReplaceAll(n, " ", ""))
			numbers = append(numbers, nInt)
		}

		problems = append(problems, Problem{
			numbers: numbers,
			op:      op,
		})
	}

	return problems, nil
}

func SolvePart1(input string) (int, error) {
	problems, _ := parseInputPart1(input)

	solutions := []int{}
	for _, problem := range problems {
		solution := 0
		switch problem.op {
		case Add:
			solution = lib.Sum(problem.numbers)
		case Multiply:
			solution = lib.Product(problem.numbers)
		}
		solutions = append(solutions, solution)
	}

	return lib.Sum(solutions), nil
}

func SolvePart2(input string) (int, error) {
	problems, _ := parseInputPart2(input)

	solutions := []int{}
	for _, problem := range problems {
		solution := 0
		switch problem.op {
		case Add:
			solution = lib.Sum(problem.numbers)
		case Multiply:
			solution = lib.Product(problem.numbers)
		}
		solutions = append(solutions, solution)
	}

	return lib.Sum(solutions), nil
}
