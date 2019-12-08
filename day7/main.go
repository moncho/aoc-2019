package main

import (
	"aoc-2019/intcode"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type amplifiers []intcode.Computer

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
	phases := phases(5)
	var maxOutput int
	var maxPhase []int
	for _, phase := range phases {
		amplifiers := new(program)
		output := amplifiers.output(phase)
		if output > maxOutput {
			maxOutput = output
			maxPhase = phase
		}
	}

	fmt.Printf("What is the highest signal that can be sent to the thrusters? %d. Phase: %v\n", maxOutput, maxPhase)

}
func phases(amplifiers int) [][]int {
	base := make([]int, amplifiers)

	for i := 0; i < amplifiers; i++ {
		base[i] = i
	}
	return permutations(base)
}
func new(program []int) amplifiers {
	return []intcode.Computer{
		intcode.New(program),
		intcode.New(program),
		intcode.New(program),
		intcode.New(program),
		intcode.New(program),
	}
}

func (a amplifiers) output(phase []int) int {
	var output int
	for i, ampli := range a {
		out, err := ampli.Run(phase[i], output)
		if err != nil {
			panic(err)
		}
		output = out
	}
	return output
}

func factorial(n int) int {
	var res int
	for ; n > 1; n-- {
		res *= n
	}
	return res
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
