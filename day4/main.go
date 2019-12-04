package main

import (
	"fmt"
)

func main() {

	var passwords []int
	for candidate := 353096; candidate <= 843212; candidate++ {
		if meetNewCriteria(candidate) {
			passwords = append(passwords, candidate)
		}
	}

	fmt.Printf("How many different passwords within the range given in your puzzle input meet these criteria? %d\n", len(passwords))

}

func meetCriteria(n int) bool {

	repeated := false
	prev := n % 10
	n /= 10
	for n > 0 {
		curr := n % 10
		if prev < curr {
			return false
		}
		if repeated || prev == curr {
			repeated = true
		}
		n /= 10
		prev = curr
	}
	return repeated
}

func meetNewCriteria(n int) bool {

	validBlock := false
	prev := n % 10
	repCount := 0
	for n := n / 10; n > 0; n = n / 10 {
		curr := n % 10
		if prev < curr {
			return false
		}
		// Ignore repetitions if a valid block has already been fouund
		if !validBlock {
			if prev == curr {
				repCount++
			} else {
				if repCount == 1 {
					validBlock = true
				}
				repCount = 0
			}
		}
		prev = curr
	}
	//Last repetition count might be valid
	return validBlock || repCount == 1
}
