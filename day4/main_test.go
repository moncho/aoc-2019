package main

import "testing"

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
