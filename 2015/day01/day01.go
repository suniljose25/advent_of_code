package day01

func Run(input string) int {
	floor := 0
	for _, r := range input {
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		}
	}
	return floor
}
