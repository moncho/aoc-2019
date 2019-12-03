package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type opcode int

const (
	sum  opcode = 1
	mul         = 2
	halt        = 99
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

	init := make([]int, len(program))
	copy(init, program)

	init[1] = 12
	init[2] = 2
	c := computer{
		init,
	}

	if err := c.run(); err != nil {
		panic(err)
	}
	fmt.Printf("What value is left at position 0 after the program halts? %d\n", c.memory[0])

	noun, verb := 0, 0
	found := false
	for i := 0; i <= 99 && !found; i++ {
		for j := 0; j <= 99; j++ {
			init := make([]int, len(program))
			copy(init, program)

			init[1] = i
			init[2] = j
			c := computer{
				init,
			}

			if err := c.run(); err != nil {
				panic(err)
			}

			if c.memory[0] == 19690720 {
				noun = i
				verb = j
				found = true
				break
			}
		}
	}
	if !found {
		panic("Noun and verb not found!!")
	}
	fmt.Printf("What is 100 * noun + verb? %d\n", 100*noun+verb)

}

type computer struct {
	memory []int
}

func (c computer) run() error {
	ip := 0
	for c.memory[ip] != halt {
		switch opcode(c.memory[ip]) {
		case sum:
			c.memory[c.memory[ip+3]] = c.memory[c.memory[ip+1]] + c.memory[c.memory[ip+2]]
		case mul:
			c.memory[c.memory[ip+3]] = c.memory[c.memory[ip+1]] * c.memory[c.memory[ip+2]]
		default:
			return fmt.Errorf("unexpected opcode %d at %d", c.memory[ip], ip)
		}
		ip += 4
	}
	return nil
}
