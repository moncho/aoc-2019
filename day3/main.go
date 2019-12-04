package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xy struct {
	x, y int
}

type line struct {
	a xy
	b xy
}

func (l line) vertical() bool {
	return l.a.x == l.b.x
}

func (l line) contains(xy xy) bool {
	minx, max := minmax(l.a.x, l.b.x)
	miny, may := minmax(l.a.y, l.b.y)

	return minx <= xy.x && xy.x <= max && miny <= xy.y && xy.y <= may
}

func (l line) intersection(other line) xy {

	if l.vertical() == other.vertical() {
		//parallel lines
		return xy{}
	}

	x := 0
	y := 0
	if l.vertical() {
		x = l.a.x
		y = other.b.y
	} else {
		x = other.a.x
		y = l.b.y
	}
	inter := xy{x, y}
	if !l.contains(inter) || !other.contains(inter) {
		return xy{}
	}

	return inter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wires []string
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}
	md := manhattanDistance(wires[0], wires[1])
	fmt.Printf("What is the Manhattan distance from the central port to the closest intersection? %d\n", md)
}

func manhattanDistance(wire1, wire2 string) int {
	path1 := path(wire1)
	path2 := path(wire2)

	candidate := 1000000
	for _, line1 := range path1 {
		for _, line2 := range path2 {
			intersect := line1.intersection(line2)
			if intersect.x != 0 && intersect.y != 0 {
				dto := distanceToOrigin(intersect)
				if dto < candidate {
					candidate = dto
				}
			}
		}
	}
	return candidate
}

func distanceToOrigin(xy xy) int {
	return abs(xy.x) + abs(xy.y)
}

func path(wire string) []line {
	turns := strings.Split(wire, ",")
	var wirePath []line
	var cur, prev xy

	for _, turn := range turns {
		dist, err := strconv.Atoi(turn[1:])
		if err != nil {
			panic(err)
		}
		switch turn[0] {
		case 'U':
			cur = xy{
				x: prev.x,
				y: prev.y - dist,
			}
		case 'D':
			cur = xy{
				x: prev.x,
				y: prev.y + dist,
			}
		case 'R':
			cur = xy{
				x: prev.x + dist,
				y: prev.y,
			}
		case 'L':
			cur = xy{
				x: prev.x - dist,
				y: prev.y,
			}
		}
		wirePath = append(wirePath, line{prev, cur})
		prev = cur
	}
	return wirePath
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minmax(a, b int) (int, int) {
	if a >= b {
		return b, a
	}
	return a, b
}
