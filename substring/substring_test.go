package substring

import "testing"

func Test_subString(t *testing.T) {
	type args struct {
		A []string
		B []string
		P string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Single match", 
			args: args{
				A: []string {"pim","pom","podi"},
				B: []string {"999999999","777888999","888888888"},
				P: "88999",
			},
			want: "pom",
		},
		{
			name: "Multiple match", 
			args: args{
				A: []string {"sander","amy","ann","michael"},
				B: []string {"123456789","234567890","789123456","123123123"},
				P: "1",
			},
			want: "ann",
		},
		{
			name: "No match", 
			args: args{
				A: []string {"adam","eva","leo"},
				B: []string {"121212121","111111111","444555666"},
				P: "112",
			},
			want: "NO CONTACT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subString(tt.args.A, tt.args.B, tt.args.P); got != tt.want {
				t.Errorf("subString() = %v, want %v", got, tt.want)
			}
		})
	}
}