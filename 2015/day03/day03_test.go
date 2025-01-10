package day03

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	if r := Run(">"); r != 2 {
		fmt.Printf("expected: 2, obtained: %d\n", r)
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	if r := Run("^>v<"); r != 4 {
		fmt.Printf("expected: 4, obtained: %d\n", r)
		t.Fail()
	}
}
