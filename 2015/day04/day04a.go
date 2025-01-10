package day04

import (
	"strconv"
)

func RunA(input string) int {
	salt := 1

	for {
		hash := getHash(input + strconv.Itoa(salt))

		if hash[:5] == "000000" {
			return salt
		}
		salt += 1
	}
}
