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

	program := intcode.LoadProgram(file)
	computer := intcode.NewV2(program)
	computer.Run()
	computer.In(1)

	var last int
	for out := range computer.Out() {
		fmt.Printf("Output: %d\n", out)
		last = out
	}

	fmt.Printf("What BOOST keycode does it produce? %d\n", last)

	computer = intcode.NewV2(program)
	computer.Run()
	computer.In(2)

	fmt.Printf("What are the coordinates of the distress signal? %d\n", <-computer.Out())

}
