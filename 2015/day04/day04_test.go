package day04

import (
	"fmt"
	"strconv"
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

func BenchmarkGetHashes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getHash("abcdef123456")
	}
}

func Benchmark1_000_000GetHashes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1_000_000; i++ {
			getHash("abcdef123456" + strconv.Itoa(i))
		}
	}
}
