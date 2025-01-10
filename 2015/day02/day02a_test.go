package day02

import (
	"fmt"
	"testing"
)

func TestCaseA1(t *testing.T) {
	if r := RunA("2x3x4"); r != 34 {
		fmt.Printf("expected: 34, obtained: %d\n", r)
		t.Fail()
	}
}

func TestCaseA2(t *testing.T) {
	if r := RunA("2x3x4\n1x1x10"); r != 48 {
		fmt.Printf("expected: 34, obtained: %d\n", r)
		t.Fail()
	}
}
