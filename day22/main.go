package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
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

	instructions := instructions(file)
	deck := newDeck(10007)
	deck.shuffle(instructions)

	fmt.Printf("After shuffling your factory order deck of 10007 cards, what is the position of card 2019? %d\n", deck.position(2019))

	cards := new(big.Int)
	cards.SetString("119315717514047", 10)

	iterations := new(big.Int)
	iterations.SetString("101741582076661", 10)
	card := cardInPos(cards, iterations, big.NewInt(2020), instructions)
	fmt.Printf("After shuffling your new, giant, factory order deck that many times, what number is on the card that ends up in position 2020? %d\n", card)

}

func newDeck(cards int) deck {
	deck := make([]int, cards)
	for i := 0; i < cards; i++ {
		deck[i] = i
	}
	return deck
}
func (d deck) shuffle(instructions []string) {
	for _, instruction := range instructions {
		parseTechnique(instruction)(d)
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
func instructions(r io.Reader) []string {
	var tt []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		tt = append(tt, strings.TrimSpace(scanner.Text()))
	}
	return tt
}
func parseTechnique(s string) technique {

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

func parse(s string) technique {

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

func cardInPos(cards, iterations, cardPos *big.Int, rules []string) *big.Int {
	// Build a linear formula to calculate a new card position from an initial position
	// and a set of shuffling techniques.
	// New cards positions is a*cardPos+b
	// Calculate a and b
	a, b := suffleFormula(cards, rules)
	//run "ax+b" iterations times
	a, b = modpow(a, b, iterations, cards)
	// New cards positions is a*cardPos+b%cards
	a = a.Mul(a, cardPos)
	a = a.Add(a, b)
	return a.Mod(a, cards)
}

// Defines a formula ax+b from the shuffling techniques. Techniques are read from last to
// first and, since it goes backward, the reversal of each technique has to be used.
// Inspired by https://www.reddit.com/r/adventofcode/comments/eeeixy/remember_the_challenges_arent_here_for_you_to/
func suffleFormula(l *big.Int, r []string) (*big.Int, *big.Int) {
	a := big.NewInt(1)
	b := big.NewInt(0)
	for i := len(r) - 1; i >= 0; i-- {
		s := strings.TrimSpace(r[i])
		if strings.HasPrefix(s, "cut") {
			n, ok := new(big.Int).SetString(s[len("cut "):], 10)
			if !ok {
				panic(s)
			}
			b = b.Add(b, n)
			b = b.Mod(b, l)
		} else if strings.HasPrefix(s, "deal with increment") {
			n, ok := new(big.Int).SetString(s[len("deal with increment "):], 10)
			if !ok {
				panic(s)
			}

			z := new(big.Int)
			z = z.ModInverse(n, l)

			a = a.Mul(a, z)
			a = a.Mod(a, l)

			b = b.Mul(b, z)
			b = b.Mod(b, l)

		} else if strings.HasPrefix(s, "deal into new stack") {
			a = a.Neg(a)
			b.Sub(l, b)
			b.Sub(b, big.NewInt(1))
		} else {
			panic(s)
		}
	}
	return a, b
}

// modpow the polynomial: (ax+b)^iterations % cards
func modpow(a, b, iterations, cards *big.Int) (*big.Int, *big.Int) {

	if iterations.Int64() == 0 {
		return big.NewInt(1), big.NewInt(0)
	}
	if iterations.Int64()%2 == 0 {
		abb := new(big.Int)
		abb = abb.Mul(a, b)
		abb = abb.Add(abb, b)
		abb = abb.Mod(abb, cards)

		aExp := new(big.Int)
		aExp = aExp.Exp(a, big.NewInt(2), cards)
		md := new(big.Int)
		md = md.Div(iterations, big.NewInt(2))
		return modpow(aExp, abb, md, cards)
	}

	newM := new(big.Int)
	newM = newM.Sub(iterations, big.NewInt(1))

	c, d := modpow(a, b, newM, cards)
	ac := new(big.Int)
	ac = ac.Mul(a, c)
	ac = ac.Mod(ac, cards)

	ad := new(big.Int)
	ad = ad.Mul(a, d)
	ad = ad.Add(ad, b)
	ad = ad.Mod(ad, cards)
	return ac, ad
}
