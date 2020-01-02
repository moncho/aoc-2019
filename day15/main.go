package main

import (
	"aoc-2019/intcode"
	"fmt"
	"os"
)

type movement int

const (
	north movement = iota + 1
	south
	west
	east
)

var movements = []movement{north,
	south,
	west,
	east}

func (m movement) direction() (int, int) {
	switch m {
	case north:
		return 0, 1
	case south:
		return 0, -1
	case west:
		return -1, 0
	default: //east
		return 1, 0
	}
}
func (m movement) backtrack() movement {
	switch m {
	case north:
		return south
	case south:
		return north
	case west:
		return east
	default: //east
		return west
	}
}

type xy struct {
	x int
	y int
}

func (a xy) move(m movement) xy {
	x, y := m.direction()
	return xy{
		x: a.x + x,
		y: a.y + y,
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	program := intcode.LoadProgram(file)

	computer := intcode.NewV2(program)

	computer.Run()
	moves, location := findOxygenSystem(computer)

	fmt.Printf("What is the fewest number of movement commands required to move the repair droid from its starting position to the location at %v of the oxygen system? %v\n",
		location,
		len(moves))
}

func findOxygenSystem(computer *intcode.V2Computer) ([]movement, xy) {
	shipMap := map[xy]int{}
	var backpath []movement
	var pos xy
	var exit xy

explore:
	for {
		for _, move := range movements {
			next := pos.move(move)
			if _, ok := shipMap[next]; !ok {
				computer.In(int(move))
				shipMap[next] = computer.Out()
				if shipMap[next] > 0 {
					pos = next
					backpath = append(backpath, move.backtrack())
				}
				if shipMap[next] == 2 {
					exit = next
					break explore
				}
				continue explore
			}
		}
		// Backtracking
		if len(backpath) < 1 {
			break
		}
		computer.In(int(backpath[len(backpath)-1]))
		computer.Out()
		pos = pos.move(movement(backpath[len(backpath)-1]))
		backpath = backpath[:len(backpath)-1]
	}

	return backpath, exit
}
