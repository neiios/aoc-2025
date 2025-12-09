package day9

import (
	_ "embed"
	"testing"
)

//go:embed day9_test.txt
var input string

func TestSolvePart1(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		expectedResult := 50

		result, _ := SolvePart1(input)

		if result != expectedResult {
			t.Errorf("Expected = %d; actual: %d", expectedResult, result)
		}
	})
}

func TestSolvePart2(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		expectedResult := 24

		result, _ := SolvePart2(input)

		if result != expectedResult {
			t.Errorf("Expected = %d; actual: %d", expectedResult, result)
		}
	})
}
