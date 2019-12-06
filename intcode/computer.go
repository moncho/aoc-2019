package intcode

import (
	"fmt"
)

type opcode int

const (
	sum    opcode = 1
	mul           = 2
	input         = 3
	output        = 4
	halt          = 99
)

// Computer is an intcode computer.
type Computer struct {
	memory []int
}

// New computer.
func New(program []int) Computer {
	return Computer{program}
}

// Run this computer
func (c Computer) Run(in int) (int, error) {
	ip := 0
	var jump int
	var out int
	for c.memory[ip] != halt {
		op := c.memory[ip] % 100
		modes := modes(c.memory[ip])
		switch opcode(op) {
		case sum:
			var param1, param2 int
			if modes[2] == 0 {
				param1 = c.memory[c.memory[ip+1]]
			} else {
				param1 = c.memory[ip+1]
			}
			if modes[1] == 0 {
				param2 = c.memory[c.memory[ip+2]]
			} else {
				param2 = c.memory[ip+2]
			}
			c.memory[c.memory[ip+3]] = param1 + param2
			jump = 4
		case mul:
			var param1, param2 int
			if modes[2] == 0 {
				param1 = c.memory[c.memory[ip+1]]
			} else {
				param1 = c.memory[ip+1]
			}
			if modes[1] == 0 {
				param2 = c.memory[c.memory[ip+2]]
			} else {
				param2 = c.memory[ip+2]
			}
			c.memory[c.memory[ip+3]] = param1 * param2
			jump = 4
		case input:
			c.memory[c.memory[ip+1]] = in
			jump = 2
		case output:
			out = c.memory[c.memory[ip+1]]
			jump = 2
			fmt.Printf("Output: %d\n", out)
		default:
			return 0, fmt.Errorf("unknown opcode %d at %d", c.memory[ip], ip)
		}
		ip += jump
	}
	return out, nil
}

func modes(i int) []int {
	modes := make([]int, 3)
	i = i / 100
	pos := 2
	for i > 0 {
		modes[pos] = i % 10
		pos--
		i = i / 10
	}
	return modes
}
