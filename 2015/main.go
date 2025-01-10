package main

import (
	"fmt"
	"os"

	"suniljose.in/aoc/2015/day01"
)

func main() {
	program := os.Args[1]

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error while reading input file in main: %+v\n", err)
		panic(err)
	}

	input := string(data)

	switch program {
	case "day01":
		fmt.Println(day01.Run(input))
	case "day01a":
		fmt.Println(day01.RunA(input))
	}

}
