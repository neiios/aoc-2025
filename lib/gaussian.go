package lib

import (
	"errors"
	"math"
)

const epsilon = 1e-9

func GaussianElimination(A [][]float64, b []float64) ([][]float64, []float64, error) {
	m := len(A)
	if m == 0 {
		return nil, nil, errors.New("matrix A is empty")
	}
	n := len(A[0])
	if len(b) != m {
		return nil, nil, errors.New("length of b must match rows of A")
	}

	augA := make([][]float64, m)
	augB := make([]float64, m)
	copy(augB, b)

	for i := range A {
		if len(A[i]) != n {
			return nil, nil, errors.New("matrix A must be rectangular (non-jagged)")
		}
		row := make([]float64, n)
		copy(row, A[i])
		augA[i] = row
	}

	pivotRow := 0

	for col := 0; col < n && pivotRow < m; col++ {
		maxVal := 0.0
		maxRow := pivotRow

		for i := pivotRow; i < m; i++ {
			absVal := math.Abs(augA[i][col])
			if absVal > maxVal {
				maxVal = absVal
				maxRow = i
			}
		}

		if maxVal < epsilon {
			continue
		}

		augA[pivotRow], augA[maxRow] = augA[maxRow], augA[pivotRow]
		augB[pivotRow], augB[maxRow] = augB[maxRow], augB[pivotRow]

		for i := pivotRow + 1; i < m; i++ {
			factor := augA[i][col] / augA[pivotRow][col]

			augA[i][col] = 0

			for j := col + 1; j < n; j++ {
				augA[i][j] -= factor * augA[pivotRow][j]
			}

			augB[i] -= factor * augB[pivotRow]
		}

		pivotRow++
	}

	return augA, augB, nil
}
