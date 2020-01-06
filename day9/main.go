package main

import (
	"aoc-2019/intcode"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	computer := intcode.NewV2(intcode.LoadProgram(file))
	computer.Run()
	computer.In(1)

	var last int
	for out := range computer.Out() {
		fmt.Printf("Output: %d\n", out)
		last = out
	}

	fmt.Printf("What BOOST keycode does it produce? %d\n", last)
}
