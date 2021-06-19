package bfsDfs

import (
	"testing"
)

func Test_bfs(t *testing.T) {
	type args struct {
		q *queue
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfs(tt.args.q); got != tt.want {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfs(t *testing.T) {
	type args struct {
		n *node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.n); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
