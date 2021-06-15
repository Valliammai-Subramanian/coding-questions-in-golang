package concatenate

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
			return false
	}
	for i, v := range a {
			if v != b[i] {
					return false
			}
	}
	return true
}

func TestConcatenation(t *testing.T) {
	type args struct {
		ary1   []int
		ary2   []int
		output *[]int
	}
	tests := []struct {
		name    string
		args    args
		desOutput []int
		wantErr bool
	}{
		{name: "Hi",
		args: args{
			ary1: []int{1,2,6},
			ary2: []int{4,5,7},
			output: &[]int{},				
		},
		desOutput: []int{1,2,4,5,6,7},
		wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Concatenation(tt.args.ary1, tt.args.ary2, tt.args.output)
			if !equal(*tt.args.output,tt.desOutput) || ((err != nil) != tt.wantErr) {
				t.Errorf("Concatenation() error = %v, wantErr %v", err, tt.wantErr)
			}			
		})
	}
}
