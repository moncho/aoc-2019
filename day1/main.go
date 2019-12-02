package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalFuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		extraFuel := moduleFuel(mass)
		totalFuel += extraFuel
		for {
			extraFuel = moduleFuel(extraFuel)
			if extraFuel <= 0 {
				break
			}
			totalFuel += extraFuel
		}
	}

	fmt.Printf("Fuel requirements: %d\n", totalFuel)
}

func moduleFuel(mass int) int {
	return (mass / 3) - 2
}
