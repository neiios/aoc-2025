package day3

import (
	"aoc/lib"
	"bufio"
	"strings"
)

func parseInput(input string) ([]*Bank, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	banks := []*Bank{}
	for scanner.Scan() {
		line := scanner.Text()
		digits := lib.ChunkStrToInts(line, 1)

		batteries := make([]*Battery, len(digits))
		for i, digit := range digits {
			batteries[i] = &Battery{
				joltage: digit,
				on:      false,
				index:   i,
			}
		}

		banks = append(banks, &Bank{batteries: batteries})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return banks, nil
}

func SolvePart1(input string) (int, error) {
	numberOfBatteriesToTurnOn := 2
	banks, _ := parseInput(input)
	return findMaximumTotalOutputJoltage(banks, numberOfBatteriesToTurnOn)
}

func SolvePart2(input string) (int, error) {
	numberOfBatteriesToTurnOn := 12
	banks, _ := parseInput(input)
	return findMaximumTotalOutputJoltage(banks, numberOfBatteriesToTurnOn)
}

func findMaximumTotalOutputJoltage(banks []*Bank, numberOfBatteriesToTurnOn int) (int, error) {
	totalOutputJoltages := []int{}
	for _, bank := range banks {
		rightMostIndex := 0
		for i := numberOfBatteriesToTurnOn - 1; i >= 0; i-- {
			largestBattery := &Battery{joltage: 0, on: false}
			for j := rightMostIndex; j < len(bank.batteries)-i; j++ {
				battery := bank.batteries[j]
				batteriesLeft := len(bank.batteries) - (j + 1)
				if battery.joltage > largestBattery.joltage && !battery.on && batteriesLeft >= i && battery.index >= largestBattery.index {
					largestBattery = battery
				}
			}
			largestBattery.on = true
			rightMostIndex = largestBattery.index
		}

		joltages := []int{}
		for _, battery := range bank.batteries {
			if battery.on {
				joltages = append(joltages, battery.joltage)
			}
		}

		totalOutputJoltages = append(totalOutputJoltages, lib.CombineInts(joltages))
	}

	return lib.Sum(totalOutputJoltages), nil
}

type Bank struct {
	batteries []*Battery
}

type Battery struct {
	joltage int
	on      bool
	index   int
}
