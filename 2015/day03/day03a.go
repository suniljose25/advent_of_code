package day03

import "strconv"

func RunA(input string) int {
	locations := make(map[string]bool)
	locations["0,0"] = true

	sX, sY := 0, 0
	rsX, rsY := 0, 0
	for i, r := range input {
		if i%2 == 0 {
			switch r {
			case '>':
				sX += 1
			case 'v':
				sY -= 1
			case '<':
				sX -= 1
			case '^':
				sY += 1
			}
			location := strconv.Itoa(sX) + "," + strconv.Itoa(sY)
			locations[location] = true

		} else {
			switch r {
			case '>':
				rsX += 1
			case 'v':
				rsY -= 1
			case '<':
				rsX -= 1
			case '^':
				rsY += 1
			}
			location := strconv.Itoa(rsX) + "," + strconv.Itoa(rsY)
			locations[location] = true
		}

	}

	return len(locations)
}
