package intcode

import (
	"fmt"
	"io"
)

type opcode int

const (
	sum opcode = iota + 1
	mul
	input
	output
	jt
	jnt
	lt
	eq
	relOffset
	halt opcode = 99
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
func (c *Computer) Run(in ...int) (int, error) {
	i := 0
	ip := 0
	relativeBase := 0
	var out int

	for c.memory[ip] != int(halt) {
		c.stackTrace = append(c.stackTrace, ip)
		op := c.memory[ip] % 100
		modes := modes(c.memory[ip])
		switch opcode(op) {
		case sum:
			param1 := readMem(modes[2], ip, relativeBase)(c.memory, 1)
			param2 := readMem(modes[1], ip, relativeBase)(c.memory, 2)

			c.memory[c.memory[ip+3]] = param1 + param2
			ip += 4
		case mul:
			param1 := readMem(modes[2], ip, relativeBase)(c.memory, 1)
			param2 := readMem(modes[1], ip, relativeBase)(c.memory, 2)

			c.memory[c.memory[ip+3]] = param1 * param2
			ip += 4
		case input:
			c.memory[c.memory[ip+1]] = in[i]
			i++
			ip += 2
		case output:
			out = readMem(modes[2], ip, relativeBase)(c.memory, 1)
			ip += 2
		case jt, jnt:
			param1 := readMem(modes[2], ip, relativeBase)(c.memory, 1)
			param2 := readMem(modes[1], ip, relativeBase)(c.memory, 2)

			if (opcode(op) == jt && param1 != 0) || (opcode(op) == jnt && param1 == 0) {
				ip = param2
			} else {
				ip += 3
			}

		case lt, eq:
			param1 := readMem(modes[2], ip, relativeBase)(c.memory, 1)
			param2 := readMem(modes[1], ip, relativeBase)(c.memory, 2)
			param3 := c.memory[ip+3]
			//param3 := readMem(modes[0], ip, relativeBase)(c.memory, 3)

			val := 0
			if (opcode(op) == lt && param1 < param2) || (opcode(op) == eq && param1 == param2) {
				val = 1
			}
			c.memory[param3] = val
			ip += 4
		case relOffset:
			relativeBase += readMem(modes[2], ip, relativeBase)(c.memory, 1)
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

type parammode int

const (
	normal parammode = iota
	position
	relative
)

func readMem(p parammode, ip, ip2 int) func([]int, int) int {
	switch p {
	case normal:
		return func(mem []int, offset int) int {
			return mem[mem[ip+offset]]
		}
	case position:
		return func(mem []int, offset int) int {
			return mem[ip+offset]
		}
	case relative:
		return func(mem []int, offset int) int {
			return mem[ip2+offset]
		}
	}
	panic("unreachable")
}

func modes(i int) []parammode {
	modes := make([]parammode, 3)
	i = i / 100
	pos := 2
	for i > 0 {
		modes[pos] = parammode(i % 10)
		pos--
		i = i / 10
	}
	return modes
}
