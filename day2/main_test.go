package main

import (
	"reflect"
	"testing"
)

func Test_computer_run(t *testing.T) {
	type fields struct {
		program []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			"1,0,0,0,99",
			fields{
				program: []int{1, 0, 0, 0, 99},
			},
			[]int{2, 0, 0, 0, 99},
		},
		{
			"2,4,4,5,99,0",
			fields{
				program: []int{2, 4, 4, 5, 99, 0},
			},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			"1,1,1,4,99,5,6,0,99",
			fields{
				program: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := computer{
				memory: tt.fields.program,
			}
			if err := c.run(); err != nil {
				t.Errorf("computer.run() error = %v", err)
				return
			}

			if !reflect.DeepEqual(c.memory, tt.want) {
				t.Errorf("computer.run() unexpected computer state = %v, want = %v", c.memory, tt.want)

			}
		})
	}
}
