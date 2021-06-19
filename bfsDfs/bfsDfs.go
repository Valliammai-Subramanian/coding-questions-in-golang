package bfsDfs

import (
	"fmt"
)

type node struct {
	name string
	next []*node
	end  bool
}

var nodes map[string]node

type queue []*node

func bfs(q *queue) bool {

	if len(*q) == 0 {
		return true
	}

	curNode := (*q)[0]
	_,ok := nodes[curNode.name]
	if (ok == true) {
		return false
	}

	nodes[curNode.name] = *curNode
	fmt.Println(curNode.name)
	*q = append((*q)[1:],curNode.next...)
	bfs(q)

	return false
}

func dfs(n *node) bool {

	curNode := n.name
	
	_,ok := nodes[curNode]
	if (ok == true) {
		return false
	}

	nodes[curNode] = *n
	fmt.Println(curNode)
	nextNode := n.next
	for _,v := range nextNode{
		dfs(v)
	}

	return true
}

func Main() {
	node1 := node{name: "node1", end: false}
	node2 := node{name: "node2", end: false}
	node3 := node{name: "node3", end: false}
	node4 := node{name: "node4", end: true}
	node5 := node{name: "node5", end: true}
	node6 := node{name: "node6", end: true}

	node1.next = []*node{&node2, &node3}
	node2.next = []*node{&node4, &node5}
	node3.next = []*node{&node6}
	node6.next = []*node{&node2}

	//make a map
	nodes = make(map[string]node)

	// Breadth first search
	// f := []*node{&node1}
	fmt.Println("Breadth First")
	f := queue{&node1}
	bfs(&f)

	//delete map values
	for k := range nodes {
		delete(nodes, k)
	}

	// Depth first search
	fmt.Println("Depth First")
	dfs(&node1)
}