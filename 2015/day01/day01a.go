package day01

// Return the position at which the lift goes to the basement
func RunA(input string) int {
	currentFloor := 0

	for i, r := range input {
		if r == '(' {
			currentFloor += 1
		}
		if r == ')' {
			currentFloor -= 1
		}
		if currentFloor == -1 {
			return i + 1
		}
	}
	return 0
}
