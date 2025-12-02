package day2

import (
	"aoc/lib"
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type IdRange struct {
	from int
	to   int
}

func parseInput(input string) ([]IdRange, error) {
	inputs := strings.Split(input, ",")
	idRanges := []IdRange{}
	re := regexp.MustCompile(`(\d+)-(\d+)`)

	for _, part := range inputs {
		splits := re.FindStringSubmatch(part)

		from, err1 := strconv.Atoi(splits[1])
		to, err2 := strconv.Atoi(splits[2])
		if err1 != nil || err2 != nil {
			return nil, errors.Join(err1, err2)
		}

		idRanges = append(idRanges, IdRange{from: from, to: to})
	}

	return idRanges, nil
}

func SolvePart1(input string) (int, error) {
	invalidIds := []int{}
	idRanges, _ := parseInput(input)

	for _, idRange := range idRanges {
		for i := idRange.from; i <= idRange.to; i++ {
			length := lib.LengthInt(i)
			if length%2 != 0 {
				continue
			}

			divisor := int(math.Pow(10, float64(length/2)))
			lhs := i / divisor
			rhs := i % divisor
			if lhs == rhs {
				invalidIds = append(invalidIds, i)
			}
		}
	}

	return lib.Sum(invalidIds), nil
}

func SolvePart2(input string) (int, error) {
	invalidIds := []int{}
	idRanges, _ := parseInput(input)

	for _, idRange := range idRanges {
		for id := idRange.from; id <= idRange.to; id++ {
			length := lib.LengthInt(id)

			for chunkSize := 1; chunkSize <= length/2; chunkSize++ {
				if length%chunkSize != 0 {
					continue
				}

				chunks := lib.ChunkInt(id, chunkSize)
				head := chunks[0]
				if lib.All(chunks, func(n int) bool { return n == head }) {
					invalidIds = append(invalidIds, id)
				}
			}
		}
	}

	return lib.Sum(lib.RemoveDuplicates(invalidIds)), nil
}
