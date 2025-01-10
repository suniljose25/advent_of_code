package day03

import "strconv"

func Run(input string) int {
	locations := make(map[string]bool)
	locations["0,0"] = true

	x, y := 0, 0
	for _, r := range input {
		switch r {
		case '>':
			x += 1
		case 'v':
			y -= 1
		case '<':
			x -= 1
		case '^':
			y += 1
		}

		location := strconv.Itoa(x) + "," + strconv.Itoa(y)
		locations[location] = true
	}

	return len(locations)
}
