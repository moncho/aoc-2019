package main

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instructions := shuffleInstructions(strings.NewReader(tt.args.s))
			tt.d.shuffle(instructions...)

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
			"",
			args{
				3,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := deal(tt.args.i)
			tt.deck.shuffle(d)
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
			"",
			args{
				3,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		{
			"",
			args{
				-4,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := cut(tt.args.i)
			tt.deck.shuffle(d)
			if !reflect.DeepEqual(tt.deck, tt.want) {
				t.Errorf("deal() = %v, want %v", tt.deck, tt.want)
			}
		})
	}
}
