// Package intcode contains an implementation of an intcode computer.
package intcode

import (
	"reflect"
	"testing"
)

func TestV2Computer_Run(t *testing.T) {
	type fields struct {
		program []int
	}
	tests := []struct {
		name   string
		fields fields
		test   func(t *testing.T, computer *V2Computer)
	}{
		{
			"program produces a copy of itself",
			fields{program: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
			func(t *testing.T, computer *V2Computer) {
				expected := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
				computer.Run()
				output := make([]int, 16)
				i := 0
				for out := range computer.Out() {
					output[i] = out
					i++
				}

				if !reflect.DeepEqual(output, expected) {
					t.Errorf("V2Computer.Run() = %v, want %v", output, expected)
				}
			},
		},
		{
			"output 1125899906842624",
			fields{program: []int{104, 1125899906842624, 99}},
			func(t *testing.T, computer *V2Computer) {
				computer.Run()
				out := <-computer.Out()
				if out != 1125899906842624 {
					t.Errorf("V2Computer.Run() = %v, want %v", out, 1125899906842624)
				}
			},
		},
		{
			"output a 16-digit number",
			fields{program: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}},
			func(t *testing.T, computer *V2Computer) {
				computer.Run()
				out := <-computer.Out()
				if out != 1219070632396864 {
					t.Errorf("V2Computer.Run() = %v, want %v", out, 1219070632396864)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewV2(tt.fields.program)
			tt.test(t, c)
		})
	}
}
