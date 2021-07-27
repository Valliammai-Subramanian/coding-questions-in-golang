package absDiff

import "testing"

func Test_diagDiff(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Matrix with even rows/cols",
			args: args{
				arr: [][]int{
					{1,2,3,4},
					{5,6,7,8},
					{9,10,11,12},
					{13,14,15,16},
				},
			},
			want: 0,
		},
		{
			name: "Matrix with odd rows/cols",
			args: args{
				arr: [][]int{
					{1,2,3,4,5},
					{6,7,8,9,10},
					{11,12,13,14,15},
					{16,17,18,19,20},
					{21,22,23,24,25},
				},
			},
			want: 0,
		},
		{
			name: "Matrix with negative elements",
			args: args{
				arr: [][]int{
					{1,2,3,4,5},
					{6,7,8,9,10},
					{0,0,0,0,0},
					{-1,-2,-3,-4,-5},
					{-6,-7,-8,-9,-10},					
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagDiff(tt.args.arr); got != tt.want {
				t.Errorf("diagDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}