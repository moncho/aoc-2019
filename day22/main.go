package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type technique func([]int)

type deck []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	instructions := shuffleInstructions(file)
	var deck deck = make([]int, 10007)
	for i := 0; i < len(deck); i++ {
		deck[i] = i
	}

	deck.shuffle(instructions...)

	fmt.Printf("After shuffling your factory order deck of 10007 cards, what is the position of card 2019? %d\n", deck.position(2019))
}

func (d deck) shuffle(tt ...technique) {
	for _, t := range tt {
		t(d)
	}
}
func (d deck) position(card int) int {
	for i, c := range d {
		if card == c {
			return i
		}
	}
	return -1
}
func shuffleInstructions(r io.Reader) []technique {
	var tt []technique
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		tt = append(tt, newTechnique(scanner.Text()))
	}
	return tt
}
func newTechnique(s string) technique {

	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "cut") {
		i, err := strconv.Atoi(s[len("cut "):])
		if err != nil {
			panic(err)
		}
		return cut(i)
	} else if strings.HasPrefix(s, "deal with increment") {
		i, err := strconv.Atoi(s[len("deal with increment "):])
		if err != nil {
			panic(err)
		}
		return deal(i)
	} else if strings.HasPrefix(s, "deal into new stack") {
		return dealNewStack()
	} else {
		panic(s)
	}

}

func cut(i int) technique {
	return func(a []int) {
		cut := i
		if cut < 0 {
			cut = len(a) + cut
		}

		cutted := append(a[cut:], a[:cut]...)
		copy(a, cutted)
	}
}
func deal(i int) technique {
	return func(a []int) {
		temp := make([]int, len(a))
		copy(temp, a)
		for j := 0; j < len(a); j++ {
			index := j * i
			if index >= len(a) {
				index %= len(a)
			}
			a[index] = temp[j]
		}
	}
}

func dealNewStack() technique {
	return func(a []int) {
		for i := len(a)/2 - 1; i >= 0; i-- {
			opp := len(a) - 1 - i
			a[i], a[opp] = a[opp], a[i]
		}
	}
}
