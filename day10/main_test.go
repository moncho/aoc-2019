package main

import (
	"strings"
	"testing"
)

func Test_asteriodsByLocation(t *testing.T) {
	type args struct {
		asteroids string
	}
	tests := []struct {
		name      string
		args      args
		want      xy
		wantCount int
	}{
		{
			"3,4 -> 8",
			args{`.#..#
.....
#####
....#
...##`},
			xy{3, 4},
			8,
		},
		{
			"5,8 -> 33",
			args{
				`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`},
			xy{5, 8},
			33,
		},
		{
			"1,2 -> 35",
			args{
				`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`},
			xy{1, 2},
			35,
		},
		{
			"6,3 -> 41",
			args{
				`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
			},
			xy{6, 3},
			41,
		},

		{
			"11,13 -> 210",
			args{
				`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`},
			xy{11, 13},
			210,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := asteroids(strings.NewReader(tt.args.asteroids))
			got, count := asteroidWithMoreVisibility(a)

			if got != tt.want {
				t.Errorf("asteriodsByLocation() = %v, want %v", got, tt.want)
			}
			if count != tt.wantCount {
				t.Errorf("asteriodsByLocation() = %v, want %v", count, tt.wantCount)
			}
		})
	}
}

func Test_angle_to_degrees(t *testing.T) {
	tests := []struct {
		name string
		a    xy
		b    xy
		want float64
	}{
		{
			"11, 13 -> 11, 12",
			xy{11, 13},
			xy{11, 12},
			270,
		},
		{
			"11, 13 -> 11, 12",
			xy{11, 13},
			xy{10, 12},
			315,
		},
		{
			"11, 13 -> 11, 14",
			xy{11, 13},
			xy{11, 14},
			90,
		},
		{
			"11, 13 -> 12, 13",
			xy{11, 13},
			xy{12, 13},
			180,
		},
		{
			"11, 13 -> 10, 13",
			xy{11, 13},
			xy{10, 13},
			0,
		},
		{
			"11, 13 -> 10, 14",
			xy{11, 13},
			xy{10, 14},
			45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.angleTo(tt.b).degrees(); got != tt.want {
				t.Errorf("degrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
