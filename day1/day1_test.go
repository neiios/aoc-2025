package day1

import (
	_ "embed"
	"testing"
)

//go:embed day1_test.txt
var input string

func TestSolvePart1(t *testing.T) {
	t.Run("solve part 1", func(t *testing.T) {
		expectedPassword := 3

		password, _ := SolvePart1(input)

		if password != expectedPassword {
			t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
		}
	})
}

func TestSolvePart2(t *testing.T) {
	t.Run("solve part 2", func(t *testing.T) {
		expectedPassword := 6

		password, _ := SolvePart2(input)

		if password != expectedPassword {
			t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
		}
	})

	t.Run("rotate2 right", func(t *testing.T) {
		var rotations = []Rotation{{Direction: "R", Distance: 1000}}
		var expectedPassword = 10

		var password = rotate2(rotations, 50)

		if password != expectedPassword {
			t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
		}
	})

	t.Run("rotate2 left", func(t *testing.T) {
		var rotations = []Rotation{{Direction: "L", Distance: 2}}
		var expectedPassword = 1

		var password = rotate2(rotations, 1)

		if password != expectedPassword {
			t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
		}
	})

	t.Run("rotate2 noop left", func(t *testing.T) {
		var rotations = []Rotation{{Direction: "L", Distance: 2}}
		var expectedPassword = 0

		var password = rotate2(rotations, 3)

		if password != expectedPassword {
			t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
		}
	})

	t.Run("rotate2 noop right", func(t *testing.T) {
		var rotations = []Rotation{{Direction: "R", Distance: 2}}
		var expectedPassword = 0

		var password = rotate2(rotations, 3)

		if password != expectedPassword {
			t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
		}
	})
}
