package day04

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
)

func Run(input string) int {
	salt := 1

	for {
		hash := getHash(input + strconv.Itoa(salt))

		if hash[:5] == "00000" {
			return salt
		}
		salt += 1
	}
}

func getHash(input string) string {
	h := md5.New()
	io.WriteString(h, input)
	return hex.EncodeToString(h.Sum(nil))
}
