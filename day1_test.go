package main

import "testing"

func TestRunDay1Part1(t *testing.T) {
	filename := "./input/day1_test.txt"
	expectedPassword := 3

	password, _ := RunDay1Part1(filename)

	if password != expectedPassword {
		t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
	}
}

func TestRunDay1Part2(t *testing.T) {
	filename := "./input/day1_test.txt"
	expectedPassword := 6

	password, _ := RunDay1Part2(filename)

	if password != expectedPassword {
		t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
	}
}

func TestRotatePart2Right(t *testing.T) {
	var rotations = []Rotation{{Direction: "R", Distance: 1000}}
	var expectedPassword = 10

	var password = rotatePart2(rotations, 50)

	if password != expectedPassword {
		t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
	}
}

func TestRotatePart2Left(t *testing.T) {
	var rotations = []Rotation{{Direction: "L", Distance: 2}}
	var expectedPassword = 1

	var password = rotatePart2(rotations, 1)

	if password != expectedPassword {
		t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
	}
}

func TestRotatePart2NoopLeft(t *testing.T) {
	var rotations = []Rotation{{Direction: "L", Distance: 2}}
	var expectedPassword = 0

	var password = rotatePart2(rotations, 3)

	if password != expectedPassword {
		t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
	}
}

func TestRotatePart2NoopRight(t *testing.T) {
	var rotations = []Rotation{{Direction: "R", Distance: 2}}
	var expectedPassword = 0

	var password = rotatePart2(rotations, 3)

	if password != expectedPassword {
		t.Errorf("Expected = %d; actual: %d", expectedPassword, password)
	}
}
