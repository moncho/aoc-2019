package main

import (
	"aoc-2019/intcode"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type amplifiers []*intcode.V2Computer

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

	var maxOutput int
	var maxPhase []int
	for _, phase := range possiblePhases(5, 0) {
		amplifiers := new(program)
		output := amplifiers.output(phase)
		if output > maxOutput {
			maxOutput = output
			maxPhase = phase
		}
	}

	fmt.Printf("What is the highest signal that can be sent to the thrusters? %d. Phase: %v\n", maxOutput, maxPhase)

	for _, phase := range possiblePhases(5, 5) {
		amplifiers := new(program)
		output := amplifiers.feedbackLoop(phase)
		if output > maxOutput {
			maxOutput = output
			maxPhase = phase
		}
	}

	fmt.Printf("What is the highest signal that can be sent to the thrusters? %d. Phase: %v\n", maxOutput, maxPhase)

}
func possiblePhases(amplifiers int, basePower int) [][]int {
	base := make([]int, amplifiers)

	for i := 0; i < amplifiers; i++ {
		base[i] = i + basePower
	}
	return permutations(base)
}
func new(program []int) amplifiers {
	return []*intcode.V2Computer{
		intcode.NewV2(program),
		intcode.NewV2(program),
		intcode.NewV2(program),
		intcode.NewV2(program),
		intcode.NewV2(program),
	}
}

func (a amplifiers) output(phase []int) int {
	var output int
	for i, ampli := range a {
		ampli.Run()
		ampli.In(phase[i])
		ampli.In(output)
		output = <-ampli.Out()
	}
	return output
}

func (a amplifiers) feedbackLoop(phase []int) int {

	// done channel for the last amplifier
	var done <-chan bool
	for i, ampli := range a {
		done = ampli.Run()
		ampli.In(phase[i])
	}
	i := 0
	next := 0
	for {
		select {
		case <-done:
			return next
		default:
		}

		a[i].In(next)
		next = <-a[i].Out()
		if i == len(a)-1 {
			i = 0
		} else {
			i++
		}
	}
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
