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
	mapShip, moves, location := findOxygenSystem(computer)

	fmt.Printf("What is the fewest number of movement commands required to move the repair droid from its starting position to the location at %v of the oxygen system? %v\n",
		location,
		len(moves))

	fmt.Printf("How many minutes will it take to fill with oxygen? %d\n", calculateMinutesTillFullOxygen(mapShip, location))
}

func findOxygenSystem(computer *intcode.V2Computer) (map[xy]int, []movement, xy) {
	shipMap := map[xy]int{}
	var backpath []movement
	var pathOxyRobot []movement
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
					pathOxyRobot = make([]movement, len(backpath))
					copy(pathOxyRobot, backpath)
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

	return shipMap, pathOxyRobot, exit
}

func calculateMinutesTillFullOxygen(shipMap map[xy]int, pos xy) int {
	queue := []xy{pos}
	ticksPerPos := map[xy]int{
		pos: 0,
	}

	visited := map[xy]bool{pos: true}
	for len(queue) > 0 {
		pos = queue[0]
		queue = queue[1:]

		for _, move := range movements {
			next := pos.move(move)
			if val, ok := shipMap[next]; ok && val != 0 && !visited[next] {
				ticksPerPos[next] = ticksPerPos[pos] + 1
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	//The last checked position has the answer.
	return ticksPerPos[pos]
}
