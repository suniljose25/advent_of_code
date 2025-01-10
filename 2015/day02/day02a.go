package day02

import (
	"sort"
	"strings"
)

func RunA(input string) int {
	measurements := strings.Split(input, "\n")

	total := 0

	for _, measurement := range measurements {
		s1, s2, s3 := convertStringDimensionToInt(measurement)

		volume := s1 * s2 * s3

		s := []int{s1, s2, s3}
		sort.Ints(s)

		smallestSidePerimenter := 2 * (s[0] + s[1])

		total += volume + smallestSidePerimenter
	}

	return total
}
