package day11

import (
	_ "embed"
	"testing"
)

//go:embed day11_1_test.txt
var inputPart1 string

//go:embed day11_2_test.txt
var inputPart2 string

func TestSolvePart1(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		expectedResult := 5

		result, _ := SolvePart1(inputPart1)

		if result != expectedResult {
			t.Errorf("Expected = %d; actual: %d", expectedResult, result)
		}
	})
}

func TestSolvePart2(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		expectedResult := 2

		result, _ := SolvePart2(inputPart2)

		if result != expectedResult {
			t.Errorf("Expected = %d; actual: %d", expectedResult, result)
		}
	})
}
