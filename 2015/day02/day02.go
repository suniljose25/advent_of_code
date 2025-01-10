package day02

import (
	"strconv"
	"strings"
)

func Run(input string) int {
	measurements := strings.Split(input, "\n")

	total := 0

	for _, measurement := range measurements {

		s1, s2, s3 := convertStringDimensionToInt(measurement)

		a1 := s1 * s2
		a2 := s2 * s3
		a3 := s1 * s3

		excess := min(a1, a2, a3)

		total += (a1+a2+a3)*2 + excess
	}
	return total
}

// Given input "2x3x4", retuns (2,3,4)
func convertStringDimensionToInt(input string) (int, int, int) {
	dimensions := strings.Split(input, "x")
	s1 := convertStringToInt(dimensions[0])
	s2 := convertStringToInt(dimensions[1])
	s3 := convertStringToInt(dimensions[2])

	return s1, s2, s3
}

func convertStringToInt(input string) (value int) {
	value, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return
}
