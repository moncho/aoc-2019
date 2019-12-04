package main

import (
	"fmt"
)

func main() {

	var passwords []int
	for candidate := 353096; candidate <= 843212; candidate++ {
		if meetCriteria(candidate) {
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
