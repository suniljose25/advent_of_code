package day01

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	if Run("((") != 2 {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	if Run("(()") != 1 {
		t.Fail()
	}
}

func TestCase3(t *testing.T) {
	if Run("(())") != 0 {
		t.Fail()
	}
}

func TestCaes4(t *testing.T) {
	if Run("()()") != 0 {
		t.Fail()
	}
}

func TestCaes5(t *testing.T) {
	if Run("(((") != 3 {
		t.Fail()
	}
}
func TestCaes6(t *testing.T) {
	if Run("(()(()(") != 3 {
		t.Fail()
	}
}
func TestCaes7(t *testing.T) {
	if Run("))(((((") != 3 {
		t.Fail()
	}
}
func TestCaes8(t *testing.T) {
	if Run("())") != -1 {
		t.Fail()
	}
}
func TestCaes9(t *testing.T) {
	if Run("))(") != -1 {
		t.Fail()
	}
}
func TestCaes10(t *testing.T) {
	if Run(")))") != -3 {
		t.Fail()
	}
}
func TestCaes11(t *testing.T) {
	fmt.Println("hello")
	if Run(")())())") != -3 {
		t.Fail()
	}
}
