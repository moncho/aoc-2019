// Package intcode contains an implementation of an intcode computer.
package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// V2Computer is the version 2 of the intcode computer.
type V2Computer struct {
	memory     []int
	stackTrace []int
	in         chan int
	out        chan int
	err        error
}

// LoadProgram loads an Intcode program from the given reader.
func LoadProgram(r io.Reader) []int {
	var program []int
	scanner := bufio.NewScanner(r)
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
	return program
}

// NewV2 new v2 computer.
func NewV2(program []int) *V2Computer {
	programCopy := make([]int, len(program))
	copy(programCopy, program)
	return &V2Computer{memory: programCopy,
		in:  make(chan int),
		out: make(chan int, 16),
	}
}

func (c *V2Computer) In(i int) {
	c.in <- i
}

func (c *V2Computer) Out() int {
	return <-c.out
}

// Run this V2Computer
func (c *V2Computer) Run() <-chan bool {

	done := make(chan bool)
	go func() {
		ip := 0

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
				c.memory[c.memory[ip+1]] = <-c.in
				ip += 2
			case output:
				var out int
				if modes[2] == 0 {
					out = c.memory[c.memory[ip+1]]
				} else {
					out = c.memory[ip+1]
				}
				c.out <- out
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
				c.err = fmt.Errorf("unknown opcode %d at IP %d", c.memory[ip], ip)
			}
		}
		close(done)
	}()
	return done
}

// BSOD prints on the given writer the state of the computer
func (c *V2Computer) BSOD(out io.Writer) {
	fmt.Fprintf(out, "CodePath: %v\n", c.stackTrace)
	fmt.Fprintf(out, "Memory: %v\n", c.memory)

}
