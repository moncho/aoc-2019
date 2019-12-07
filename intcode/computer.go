package intcode

import (
	"fmt"
	"io"
)

type opcode int

const (
	sum    opcode = 1
	mul           = 2
	input         = 3
	output        = 4
	jt            = 5
	jnt           = 6
	lt            = 7
	eq            = 8
	halt          = 99
)

// Computer is an intcode computer.
type Computer struct {
	memory     []int
	stackTrace []int
}

// New computer.
func New(program []int) Computer {
	programCopy := make([]int, len(program))
	copy(programCopy, program)
	return Computer{memory: programCopy}
}

// Run this computer
func (c *Computer) Run(in int) (int, error) {
	ip := 0
	var out int

	for c.memory[ip] != halt {
		c.stackTrace = append(c.stackTrace, ip)
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
			ip += 4
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
			ip += 4
		case input:

			c.memory[c.memory[ip+1]] = in
			ip += 2
		case output:
			if modes[2] == 0 {
				out = c.memory[c.memory[ip+1]]
			} else {
				out = c.memory[ip+1]
			}
			ip += 2
		case jt, jnt:
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

			if (opcode(op) == jt && param1 != 0) || (opcode(op) == jnt && param1 == 0) {
				ip = param2
			} else {
				ip += 3
			}

		case lt, eq:
			var param1, param2, param3 int
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

			param3 = c.memory[ip+3]
			val := 0
			if (opcode(op) == lt && param1 < param2) || (opcode(op) == eq && param1 == param2) {
				val = 1
			}
			c.memory[param3] = val
			ip += 4
		default:
			return 0, fmt.Errorf("unknown opcode %d at IP %d", c.memory[ip], ip)
		}
	}
	return out, nil
}

// BSOD prints on the given writer the state of the computer
func (c *Computer) BSOD(out io.Writer) {
	fmt.Fprintf(out, "CodePath: %v\n", c.stackTrace)
	fmt.Fprintf(out, "Memory: %v\n", c.memory)

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
