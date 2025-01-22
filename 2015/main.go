package main

import (
	"fmt"
	"os"

	"suniljose.in/aoc/2015/day01"
	"suniljose.in/aoc/2015/day02"
	"suniljose.in/aoc/2015/day03"
	"suniljose.in/aoc/2015/day04"
	"suniljose.in/aoc/2015/day05"
)

type runner func(string) int

func main() {
	program := os.Args[1]

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error while reading input file in main: %+v\n", err)
		panic(err)
	}

	input := string(data)

	fncMap := make(map[string]func(string) int)
	fncMap["day01"] = day01.Run
	fncMap["day01a"] = day01.RunA
	fncMap["day02"] = day02.Run
	fncMap["day02a"] = day02.RunA
	fncMap["day03"] = day03.Run
	fncMap["day03a"] = day03.RunA
	fncMap["day04"] = day04.Run
	fncMap["day04a"] = day04.RunA
	fncMap["day05"] = day05.Run

	f, ok := fncMap[program]
	if !ok {
		fmt.Printf("invalid program name: %s", program)
		panic(program)
	}

	fmt.Println(f(input))
}
