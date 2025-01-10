package main

import (
	"fmt"
	"os"

	"suniljose.in/aoc/2015/day01"
	"suniljose.in/aoc/2015/day02"
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

	f, ok := fncMap[program]
	if !ok {
		fmt.Printf("invalid program name: %s", program)
		panic(program)
	}

	fmt.Println(f(input))
}
