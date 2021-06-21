package bfsDfs

import (
	"testing"
)

type variables struct {
	node1   *node
	node2   *node
	node3   *node
	node4   *node
	node5   *node
	node6   *node
}

func values() variables {
	defaultValues := variables{
		node1: &node{name: "node1", end: false},
		node2: &node{name: "node2", end: false},
		node3: &node{name: "node3", end: false},
		node4: &node{name: "node4", end: true},
		node5: &node{name: "node5", end: true},
		node6: &node{name: "node6", end: true},
	}

	defaultValues.node1.next = queue{defaultValues.node2, defaultValues.node3}
	defaultValues.node2.next = queue{defaultValues.node4, defaultValues.node5}
	defaultValues.node3.next = queue{defaultValues.node6}
	defaultValues.node6.next = queue{defaultValues.node2}

	return defaultValues
}

func Test_bfs(t *testing.T) {
	type args struct {
		q *queue
	}
	tests := []struct {
		name string
		args args
		mapValues []string
		want bool
	}{
		{
			name: "EmptyTest",
			args: args{
				q: &(queue{}),
			},
			mapValues: []string{},
			want: true,
		},
		{
			name: "SingleElement",
			args: args{
				q: 	&queue{values().node1},
			},
			mapValues: []string{"node1", "node2", "node3", "node4", "node5", "node6"},
			want: false,
		},
	}

	for _, tt := range tests {
		//make a map
		nodes = make(map[string]node)
		t.Run(tt.name, func(t *testing.T) {
			if got := bfs(tt.args.q); got != tt.want {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
		for _, k := range tt.mapValues {
			if _, ok := nodes[k]; !ok {
				t.Errorf("Node %v missing", k)
			}
		}
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
		// {
		// 	name: "EmptyTest",
		// 	args: args{
		// 		n: nil,
		// 	},
		// 	want: true,
		// },
		{
			name: "SingleElement",
			args: args{
				n: 	values().node1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.n); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}