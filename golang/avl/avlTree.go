package main

import "math"

type AVL struct {
	root *Node
}

func createNewTree() *AVL {
	return &AVL{
		root: nil,
	}
}

func search(root *Node, key string) *Node {
	if root == nil {
		return nil
	}
	if root.key == key {
		return root
	}
	if root.key > key {
		search(root.left, key)
	} else {
		search(root.right, key)
	}
	return nil
}

func checkHeight(root *Node) int {
	if root == nil {
		return 0
	}
	var left int = checkHeight(root.left)
	var right int = checkHeight(root.right)

	return int(math.Max(float64(left), float64(right))) + 1

}

func calcBalance(node *Node) int {
	if node == nil {
		return 0
	} else {
		return checkHeight(node.right) - checkHeight(node.left)
	}
}

func insertOnTree(tree *AVL, key string, value int) {
	if tree.root == nil {
		tree.root = &Node{key: key, value: value, right: nil, left: nil}
	} else {
		insertNode(tree.root, key, value)
	}
}

func insertNode(node *Node, key string, value int) {
	if node == nil {
		return
	}
	if node.key > key {
		if node.left == nil {
			node.left = &Node{key: key, value: value, right: nil, left: nil}
		} else {
			insertNode(node.left, key, value)
		}
		//After new node already inserted(or we just can't go deeper this way)
		//lets check for unbalances
		if calcBalance(node) < -1 {
			if calcBalance(node.left) > 0 {
				node.left = leftRotation(node.left)
			} else {
				node.left = rightRotation(node.left)
			}
		}
	}
	if node.key <= key {
		if node.right == nil {
			node.right = &Node{key: key, value: value, right: nil, left: nil}
		} else {
			insertNode(node.right, key, value)
		}
		//After new node already inserted(or we just can't go deeper this way)
		//lets check for unbalances
		if calcBalance(node) > 1 {
			if calcBalance(node.left) < 0 {
				node.right = rightRotation(node.right)
			} else {
				node.right = leftRotation(node.right)
			}
		}
	}

}

func leftRotation(node *Node) *Node {
	var rightNode *Node = node.right

	node.right = rightNode.left
	rightNode.left = node
	return rightNode
}

func rightRotation(node *Node) *Node {
	var leftNode *Node = node.left

	node.left = leftNode.right
	leftNode.right = node
	return leftNode
}

func toString(root *Node) {
	if root == nil {
		return
	} else {
		toString(root.left)
		println(root.key)
		toString(root.right)
	}
}
