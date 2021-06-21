package reformat

import "testing"

func Test_reformat(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "One block with two digits",
			args: args{
				S: "00-44  48 5555 8361",
			},
			want: "004-448-555-583-61",
		},
		{
			name: "Two blocks with two digits",
			args: args{
				S: "0 - 22 1985--324",
			},
			want: "022-198-53-24",
		},
		{
			name: "No block with two digits",
			args: args{
				S: "555372654",
			},
			want: "555-372-654",
		},	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformat(tt.args.S); got != tt.want {
				t.Errorf("reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}