package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc-2019/intcode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var program []int
	for scanner.Scan() {
		t := scanner.Text()
		for _, s := range strings.Split(t, ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			program = append(program, i)
		}
	}
	computer := intcode.New(program)
	out, err := computer.Run(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("After providing 1 to the only input instruction and passing all the tests, what diagnostic code does the program produce? %d\n", out)

	computer = intcode.New(program)
	out, err = computer.Run(5)
	if err != nil {
		panic(err)
	}

	fmt.Printf("What is the diagnostic code for system ID 5?? %d\n", out)
}
