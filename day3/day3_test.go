package day3

import (
	_ "embed"
	"testing"
)

//go:embed day3_test.txt
var input string

func TestSolvePart1(t *testing.T) {
	t.Run("find total joltage", func(t *testing.T) {
		expectedResult := 357

		result, _ := SolvePart1(input)

		if result != expectedResult {
			t.Errorf("Expected = %d; actual: %d", expectedResult, result)
		}
	})

	t.Run("find total joltage", func(t *testing.T) {
		expectedResult := 92

		result, _ := SolvePart1("818181911112111")

		if result != expectedResult {
			t.Errorf("Expected = %d; actual: %d", expectedResult, result)
		}
	})
}

func TestSolvePart2(t *testing.T) {
	t.Run("find total joltage", func(t *testing.T) {
		expectedResult := 3121910778619

		result, _ := SolvePart2(input)

		if result != expectedResult {
			t.Errorf("Expected = %d; actual: %d", expectedResult, result)
		}
	})
}
