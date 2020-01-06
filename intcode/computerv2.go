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
	memory       []int
	stackTrace   []int
	in           chan int
	out          chan int
	err          error
	ip           int
	relativeBase int
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
		out: make(chan int),
	}
}

// In inputs the given value to this computer.
func (c *V2Computer) In(i int) {
	c.in <- i
}

// Out returns this computer output.
func (c *V2Computer) Out() <-chan int {
	return c.out
}

func (c *V2Computer) writeMem(p parammode, offset int, value int) {

	var writeAt int
	switch p {
	case normal, position:
		writeAt = c.memory[c.ip+offset]
	case relative:
		writeAt = c.relativeBase + c.memory[c.ip+offset]
	}

	if writeAt >= len(c.memory) {
		newMem := make([]int, nextPow2(writeAt))
		copy(newMem, c.memory)
		c.memory = newMem
	}
	c.memory[writeAt] = value
}

func (c *V2Computer) readMem(p parammode, offset int) int {
	var readAt int
	switch p {
	case normal:
		readAt = c.memory[c.ip+offset]
	case position:
		readAt = c.ip + offset
	case relative:
		readAt = c.relativeBase + c.memory[c.ip+offset]
	}
	if readAt >= len(c.memory) {
		return 0
	}
	return c.memory[readAt]
}

// Run this V2Computer
func (c *V2Computer) Run() <-chan bool {
	done := make(chan bool)
	go func() {
		for c.memory[c.ip] != int(halt) {
			c.stackTrace = append(c.stackTrace, c.ip)
			op := c.memory[c.ip] % 100
			modes := modes(c.memory[c.ip])

			switch opcode(op) {
			case sum:
				param1 := c.readMem(modes[2], 1)
				param2 := c.readMem(modes[1], 2)
				c.writeMem(modes[0], 3, param1+param2)
				c.ip += 4
			case mul:
				param1 := c.readMem(modes[2], 1)
				param2 := c.readMem(modes[1], 2)
				c.writeMem(modes[0], 3, param1*param2)
				c.ip += 4
			case input:
				c.writeMem(modes[2], 1, <-c.in)
				c.ip += 2
			case output:
				out := c.readMem(modes[2], 1)
				c.out <- out
				c.ip += 2
			case jt, jnt:
				param1 := c.readMem(modes[2], 1)
				param2 := c.readMem(modes[1], 2)
				if (opcode(op) == jt && param1 != 0) || (opcode(op) == jnt && param1 == 0) {
					c.ip = param2
				} else {
					c.ip += 3
				}
			case lt, eq:
				param1 := c.readMem(modes[2], 1)
				param2 := c.readMem(modes[1], 2)

				val := 0
				if (opcode(op) == lt && param1 < param2) || (opcode(op) == eq && param1 == param2) {
					val = 1
				}
				c.writeMem(modes[0], 3, val)
				c.ip += 4
			case relOffset:
				c.relativeBase += c.readMem(modes[2], 1)
				c.ip += 2
			default:
				c.err = fmt.Errorf("unknown opcode %d at IP %d", c.memory[c.ip], c.ip)
			}
		}
		close(c.out)
		close(done)
	}()
	return done
}

// BSOD prints on the given writer the state of the computer
func (c *V2Computer) BSOD(out io.Writer) {
	fmt.Fprintf(out, "CodePath: %v\n", c.stackTrace)
	fmt.Fprintf(out, "Memory: %v\n", c.memory)

}

func nextPow2(n int) int {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}
