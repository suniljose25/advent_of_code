package day03

import (
	"fmt"
	"testing"
)

func TestCaseA1(t *testing.T) {
	if r := RunA("^v"); r != 3 {
		fmt.Printf("expected: 3, obtained: %d\n", r)
		t.Fail()
	}
}

func TestCaseA2(t *testing.T) {
	if r := RunA("^>v<"); r != 3 {
		fmt.Printf("expected: 3, obtained: %d\n", r)
		t.Fail()
	}
}
