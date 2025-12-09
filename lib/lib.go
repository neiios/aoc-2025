package lib

import (
	"math"
	"strconv"

	"golang.org/x/exp/constraints"
)

// Inclusive integer range [a, b]
func RangeInt(a, b int) []int {
	if a > b {
		panic("RangeInt: a must be <= b")
	}
	if a == b {
		return []int{a}
	}
	n := b - a
	out := make([]int, n+1)
	for i := 0; i <= n; i++ {
		out[i] = a + i
	}
	return out
}

// Length of an integer
func LengthInt[T constraints.Integer](i T) T {
	if i == 0 {
		return 1
	}
	if i < 0 {
		i = -i
	}
	var length T
	for i > 0 {
		i /= 10
		length++
	}
	return length
}

// Go has no Sum builtin :kekw
func Sum[T constraints.Integer | constraints.Float](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Product[T constraints.Integer | constraints.Float](numbers []T) T {
	if len(numbers) == 0 {
		return 0
	}
	product := numbers[0]
	for _, num := range numbers[1:] {
		product *= num
	}
	return product
}

// Splits an integer into equally sized chunks
func ChunkInt(i int, size int) []int {
	length := LengthInt(i)
	if size < 1 || size >= length {
		return []int{i}
	}

	chunks := []int{}
	for length > 0 {
		chunkSize := min(length, size)
		divisor := int(math.Pow10(length - chunkSize))
		chunk := i / divisor
		chunks = append(chunks, chunk)
		i = i % divisor
		length -= chunkSize
	}

	return chunks
}

// Splits a string representation of a number into equally sized chunks and converts them to integers
func ChunkStrToInts(s string, size int) []int {
	if size < 1 || size >= len(s) {
		val, _ := strconv.Atoi(s)
		return []int{val}
	}

	chunks := []int{}
	for len(s) > 0 {
		chunkSize := min(len(s), size)
		chunkStr := s[:chunkSize]
		val, _ := strconv.Atoi(chunkStr)
		chunks = append(chunks, val)
		s = s[chunkSize:]
	}

	return chunks
}

// Checks whether all slice elements match a given predicate
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Remove duplicates from a given slice
func RemoveDuplicates[T comparable](slice []T) []T {
	uniqueMap := make(map[T]bool)
	for _, val := range slice {
		uniqueMap[val] = true
	}

	result := make([]T, 0, len(uniqueMap))
	for key := range uniqueMap {
		result = append(result, key)
	}

	return result
}

// Combines a list of integers into a single integer by concatenating them
func CombineInts[T constraints.Integer](numbers []T) T {
	var result T
	for _, num := range numbers {
		digits := LengthInt(num)
		result = result*T(math.Pow10(int(digits))) + num
	}
	return result
}

func TransposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}

func RemoveEveryNRow(rows []string, n int) []string {
	var result []string
	for i, row := range rows {
		if i%n == 0 {
			result = append(result, row)
		}
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
