package day05

import (
	"slices"
	"strings"
)

var vowels []rune = []rune{'a', 'e', 'i', 'o', 'u'}

func Run(input string) int {
	lines := strings.Split(input, "\n")

	niceCount := 0
	for _, line := range lines {
		if isNiceString(line) {
			niceCount++
		}
	}

	return niceCount
}

func isNiceString(input string) bool {
	return !hasRestrictedStrings(input) && hasAtleastThreeVowels(input) && hasContigousRunes(input)
}

func hasAtleastThreeVowels(input string) bool {
	countMap := make(map[rune]int)
	for _, r := range vowels {
		countMap[r] = 0
	}

	for _, r := range input {
		if slices.Contains(vowels, r) {
			countMap[r] += 1
		}
	}

	total := 0
	for _, r := range vowels {
		total += countMap[r]
	}
	return total >= 3
}

func hasContigousRunes(input string) bool {
	var prev rune
	for i, r := range input {
		if i == 0 {
			prev = r
			continue
		}
		if prev == r {
			return true
		}
		prev = r
	}
	return false
}

func hasRestrictedStrings(input string) bool {
	return strings.Contains(input, "ab") || strings.Contains(input, "cd") || strings.Contains(input, "pq") || strings.Contains(input, "xy")
}
