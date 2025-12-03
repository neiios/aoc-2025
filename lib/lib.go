package lib

import (
	"math"
	"math/big"

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

// Splits an integer into equally sized chunks
func ChunkInt[T constraints.Integer](i T, size int) []T {
	length := LengthInt(i)
	if size < 1 || size >= int(length) {
		return []T{i}
	}

	chunks := []T{}
	for length > 0 {
		chunkSize := min(length, T(size))
		divisor := T(math.Pow10(int(length - chunkSize)))
		chunk := i / divisor
		chunks = append(chunks, chunk)
		i = i % divisor
		length -= chunkSize
	}

	return chunks
}

// Splits a big integer into equally sized chunks
func ChunkBigInt(num *big.Int, size int) []*big.Int {
	str := num.String()
	if size < 1 || size >= len(str) {
		return []*big.Int{new(big.Int).Set(num)}
	}

	chunks := []*big.Int{}
	for len(str) > 0 {
		chunkSize := min(len(str), size)
		chunkStr := str[:chunkSize]
		chunk := new(big.Int)
		chunk.SetString(chunkStr, 10)
		chunks = append(chunks, chunk)
		str = str[chunkSize:]
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
