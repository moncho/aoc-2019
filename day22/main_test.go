package main

import (
	"math/big"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_deck_shuffle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		d      deck
		args   args
		expect deck
	}{
		{
			"",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{
				`deal with increment 7
				deal into new stack
				deal into new stack`,
			},
			[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			"",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{`cut 6
			deal with increment 7
			deal into new stack`},
			[]int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			"",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{
				`deal with increment 7
			deal with increment 9
			cut -2`,
			},
			[]int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			"",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{
				`deal into new stack
				cut -2
				deal with increment 7
				cut 8
				cut -4
				deal with increment 7
				cut 3
				deal with increment 9
				deal with increment 3
				cut -1`,
			},
			[]int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instructions := instructions(strings.NewReader(tt.args.s))
			tt.d.shuffle(instructions)

			if !reflect.DeepEqual(tt.d, tt.expect) {
				t.Errorf("After shuffling, expected: %v, got: %v", tt.expect, tt.d)
			}
		})
	}
}

func Test_deal(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		deck deck
		want deck
	}{
		{
			"deal with increment 3",
			args{
				3,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
		{
			"deal with increment 7",
			args{
				7,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.shuffle([]string{tt.name})
			if !reflect.DeepEqual(tt.deck, tt.want) {
				t.Errorf("deal() = %v, want %v", tt.deck, tt.want)
			}
		})
	}
}

func Test_cut(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		deck deck
		want deck
	}{
		{
			"cut 3",
			args{
				3,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		{
			"cut 2",
			args{
				2,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1},
		},
		{
			"cut -4",
			args{
				-4,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
		{
			"cut -2",
			args{
				-2,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{8, 9, 0, 1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.shuffle([]string{tt.name})
			if !reflect.DeepEqual(tt.deck, tt.want) {
				t.Errorf("cut() = %v, want %v", tt.deck, tt.want)
			}
		})
	}
}

func Test_dealNewStack(t *testing.T) {
	tests := []struct {
		name string
		deck deck
		want deck
	}{
		{
			"deal into new stack",
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.shuffle([]string{tt.name})
			if !reflect.DeepEqual(tt.deck, tt.want) {
				t.Errorf("deal() = %v, want %v", tt.deck, tt.want)
			}
		})
	}
}

func Test_cardInPos(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf(err.Error())
	}

	defer file.Close()

	instructions := instructions(file)
	type args struct {
		cards      *big.Int
		iterations *big.Int
		cardPos    *big.Int
		rules      []string
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			"10007 cards, 1 iteration",
			args{
				big.NewInt(10007),
				big.NewInt(1),
				big.NewInt(3324),
				instructions,
			},
			big.NewInt(2019),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cardInPos(tt.args.cards, tt.args.iterations, tt.args.cardPos, tt.args.rules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cardInPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
