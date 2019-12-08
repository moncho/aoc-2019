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
}
