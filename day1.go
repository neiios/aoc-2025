package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Rotation struct {
	Direction string
	Distance  int
}

func parseLine(line string) (Rotation, error) {
	distance, err := strconv.Atoi(line[1:])
	if err != nil {
		return Rotation{}, err
	}

	return Rotation{
		Direction: line[:1],
		Distance:  distance,
	}, nil
}

func parseInput(filename string) ([]Rotation, error) {
	input, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var rotations = []Rotation{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var line = scanner.Text()
		rotation, err := parseLine(line)
		if err != nil {
			return []Rotation{}, err
		}
		rotations = append(rotations, rotation)
	}

	if err := scanner.Err(); err != nil {
		return []Rotation{}, err
	}

	return rotations, nil
}

func RunDay1Part1(filename string) (int, error) {
	rotations, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	return rotatePart1(rotations, 50), nil
}

func rotatePart1(rotations []Rotation, initialDial int) int {
	var dial = initialDial
	var password = 0
	for _, rotation := range rotations {
		switch rotation.Direction {
		case "L":
			dial = (dial - rotation.Distance) % 100
		case "R":
			dial = (dial + rotation.Distance) % 100
		}

		if dial == 0 {
			password++
		}
	}
	return password
}

func RunDay1Part2(filename string) (int, error) {
	rotations, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	return rotatePart2(rotations, 50), nil
}

func rotatePart2(rotations []Rotation, initialDial int) int {
	var dial = initialDial
	var password = 0
	for _, rotation := range rotations {
		var prevDial = dial
		switch rotation.Direction {
		case "L":
			dial = (((dial - rotation.Distance) % 100) + 100) % 100
		case "R":
			dial = (dial + rotation.Distance) % 100
		}

		password += rotation.Distance / 100

		if dial == 0 {
			password++
		} else if prevDial == 0 {
			continue
		} else if rotation.Direction == "L" && prevDial-rotation.Distance%100 < 0 {
			password++
		} else if rotation.Direction == "R" && prevDial+rotation.Distance%100 > 100 {
			password++
		}
	}
	return password
}
