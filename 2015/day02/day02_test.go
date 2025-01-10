package day02

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	if r := Run("2x3x4"); r != 58 {
		fmt.Printf("expected: 58, obtained: %d\n", r)
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	if r := Run("2x3x4\n1x1x10"); r != 101 {
		fmt.Printf("expected: 58, obtained: %d\n", r)
		t.Fail()
	}
}

func TestConvertingInputDimensions(t *testing.T) {
	v1, v2, v3 := convertStringDimensionToInt("2x3x4")

	if v1 != 2 || v2 != 3 || v3 != 4 {
		t.Fail()
	}
}

func TestConvertingIntToString(t *testing.T) {
	if convertStringToInt("2") != 2 {
		t.Fail()
	}
}
