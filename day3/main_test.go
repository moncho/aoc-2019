package main

import "testing"

func Test_manhattanDistance(t *testing.T) {
	type args struct {
		wire1 string
		wire2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"R8,U5,L5,D3 - U7,R6,D4,L4",
			args{
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			},
			6,
		},
		{
			"(R75,D30,R83,U83,L12,D49,R71,U7,L72) (U62,R66,U55,R34,D71,R55,D58,R83)",
			args{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83"},
			159,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manhattanDistance(tt.args.wire1, tt.args.wire2); got != tt.want {
				t.Errorf("manhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
