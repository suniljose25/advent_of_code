package day04

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	if r := Run("abcdef"); r != 609043 {
		fmt.Printf("expected: 609043, obtained: %d\n", r)
		t.Fail()
	}
}

func TestGetHash(t *testing.T) {
	if r := getHash("abcdef609043"); r != "000001dbbfa3a5c83a2d506429c7b00e" {
		fmt.Printf("expected: 000001dbbfa3a5c83a2d506429c7b00e, obtained: %v\n", r)
		t.Fail()
	}
}
