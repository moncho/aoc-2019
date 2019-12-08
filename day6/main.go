package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	orbits := make(map[string]string)

	for scanner.Scan() {
		objects := strings.Split(scanner.Text(), ")")
		orbits[objects[1]] = objects[0]
	}

	count := 0

	for _, v := range orbits {
		count++
		for k, ok := orbits[v]; ok; k, ok = orbits[k] {
			count++
		}
	}
	fmt.Printf("What is the total number of direct and indirect orbits in your map data? %d\n", count)

	myPath := orbitalPath(orbits, "YOU")
	santaPath := orbitalPath(orbits, "SAN")

	transfers := 0

found:
	for i, me := range myPath {
		for j, santa := range santaPath {
			if santa == me {
				transfers = i + j
				break found
			}
		}

	}

	fmt.Printf("What is the minimum number of orbital transfers required to move from the object YOU are orbiting to the object SAN is orbiting? %d\n", transfers)
}

func orbitalPath(orbits map[string]string, s string) []string {
	var path []string
	for k, ok := orbits[s]; ok; k, ok = orbits[k] {
		path = append(path, k)
	}
	return path
}
