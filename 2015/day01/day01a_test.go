package day01

import (
	"testing"
)

func TestCaseA1(t *testing.T) {
	if RunA(")") != 1 {
		t.Fail()
	}
}

func TestCaseA2(t *testing.T) {
	if RunA("()())") != 5 {
		t.Fail()
	}
}
