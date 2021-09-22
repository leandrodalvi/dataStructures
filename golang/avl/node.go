package main

type Node struct {
	left  *Node
	right *Node
	value int
	key   string
}

func createNewNode(key string, value int, right *Node, left *Node) *Node {
	return &Node{
		left:  left,
		right: right,
		value: value,
		key:   key,
	}
}
