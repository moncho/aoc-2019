package main

import (
	"testing"
)

func Test_meetCriteria(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"111111",
			args{
				111111,
			},
			true,
		},
		{
			"223450",
			args{
				223450,
			},
			false,
		},
		{
			"123789",
			args{
				123789,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := meetCriteria(tt.args.n); got != tt.want {
				t.Errorf("meetCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_meetNewCriteria(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"112233",
			args{
				112233,
			},
			true,
		},
		{
			"123444",
			args{
				123444,
			},
			false,
		},
		{
			"111122",
			args{
				111122,
			},
			true,
		},
		{
			"111233",
			args{
				111233,
			},
			true,
		},
		{
			"112333",
			args{
				112333,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := meetNewCriteria(tt.args.n); got != tt.want {
				t.Errorf("meetNewCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}
