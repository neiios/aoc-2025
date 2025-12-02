package lib

import (
	"math"

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
func LengthInt(i int) int {
	if i == 0 {
		return 1
	}
	if i < 0 {
		i = -i
	}
	length := 0
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

// Checks whether all slice elements match a given predicate
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Remove duplicate ints from a given slice
func RemoveDuplicates(slice []int) []int {
	uniqueMap := make(map[int]bool)
	for _, val := range slice {
		uniqueMap[val] = true
	}

	result := make([]int, 0, len(uniqueMap))
	for key := range uniqueMap {
		result = append(result, key)
	}

	return result
}
