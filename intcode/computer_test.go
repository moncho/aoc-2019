package intcode

import (
	"reflect"
	"testing"
)

func TestComputer_Run(t *testing.T) {
	type fields struct {
		memory []int
	}
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantMem []int
		want    int
		wantErr bool
	}{
		{
			"3,0,4,0,99",
			fields{
				[]int{3, 0, 4, 0, 99},
			},
			args{
				5,
			},
			[]int{5, 0, 4, 0, 99},
			5,
			false,
		},
		{
			"1002,4,3,4,33",
			fields{
				[]int{1002, 4, 3, 4, 33},
			},
			args{
				0,
			},
			[]int{1002, 4, 3, 4, 99},
			0,
			false,
		},
		{
			"1,0,0,0,99",
			fields{
				[]int{1, 0, 0, 0, 99},
			},
			args{
				0,
			},
			[]int{2, 0, 0, 0, 99},
			0,
			false,
		},
		{
			"1,1,1,4,99,5,6,0,99",
			fields{
				[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			args{
				0,
			},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			0,
			false,
		},
		{
			"101,-88,1,0,99",
			fields{
				[]int{101, -88, 1, 0, 99},
			},
			args{
				0,
			},
			[]int{-176, -88, 1, 0, 99},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Computer{
				memory: tt.fields.memory,
			}
			got, err := c.Run(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Computer.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(c.memory, tt.wantMem) {
				t.Errorf("Computer.Run() memory = %v, want %v", c.memory, tt.wantMem)
			}
			if got != tt.want {
				t.Errorf("Computer.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputer_IO(t *testing.T) {

	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		args    args
		out     int
		wantErr bool
	}{
		{
			"4",
			args{
				4,
			},
			999,
			false,
		},
		{
			"8",
			args{
				8,
			},
			1000,
			false,
		},
		{
			"9",
			args{
				17,
			},
			1001,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			programCopy := make([]int, len(program))
			copy(programCopy, program)
			c := Computer{
				memory: programCopy,
			}
			got, err := c.Run(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Computer.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.out {
				t.Errorf("Computer.Run() = %v, want %v", got, tt.out)
			}
		})
	}
}
