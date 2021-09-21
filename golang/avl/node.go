package avl

type Node struct {
	left  *Node
	right *Node
	value interface{}
	key   string
	w     int
}

func createNewNode(key string, value interface{}, right *Node, left *Node) *Node {
	newP := new(Node)
	newP.key = key
	newP.value = value
	newP.right = right
	newP.left = left
	newP.w = 0
	return newP
}
