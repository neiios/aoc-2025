package day12

import (
	"strconv"
	"strings"
)

// type Present struct {
// 	shapes [][][]bool
// 	size int
// }

type Region struct {
	width      int
	length     int
	quantities []int
}

func SolvePart1(input string) (int, error) {
	parts := strings.Split(input, "\n\n")
	// presents := []Present{}
	regions := []Region{}

	// for _, part := range parts[:len(parts)-1] {
	// 	present := Present{}
	// 	for _, subpart := range strings.Split(part, "\n")[1:] {
	// 		for _, cell := range subpart {
	// 			if cell == '#' {
	// 				present.size++
	// 			}
	// 		}
	// 	}
	// 	presents = append(presents, present)
	// }

	for regionStr := range strings.SplitSeq(parts[len(parts)-1], "\n") {
		if regionStr == "" {
			continue
		}

		region := Region{}
		regionParts := strings.Split(regionStr, " ")
		sizes := strings.Split(regionParts[0][:len(regionParts[0])-1], "x")
		region.width, _ = strconv.Atoi(sizes[0])
		region.length, _ = strconv.Atoi(sizes[1])

		for _, part := range regionParts[:len(regionParts)-1] {
			quantity, _ := strconv.Atoi(part)
			region.quantities = append(region.quantities, quantity)
		}

		regions = append(regions, region)
	}

	count := 0
	for _, region := range regions {
		availableCapacity := region.width * region.length
		requiredCapacity := 0
		for j := range len(region.quantities) {
			requiredCapacity += 9 * region.quantities[j]
		}
		if requiredCapacity <= availableCapacity {
			count++
		}
	}

	return count, nil
}

func SolvePart2(input string) (int, error) {
	return 0, nil
}
