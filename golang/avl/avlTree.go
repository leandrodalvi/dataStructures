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

//utils funcs:

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

func toString(root *Node) {
	if root == nil {
		return
	} else {
		toString(root.left)
		println(root.key)
		toString(root.right)
	}
}

//Aux operations

func leftRotation(node *Node, avl *AVL) *Node {
	var rightNode *Node = node.right

	node.right = rightNode.left
	rightNode.left = node
	rightNode.parent = node.parent
	node.parent = rightNode
	if node != avl.root {
		if rightNode.parent.left == node {
			rightNode.parent.left = rightNode
		} else {
			rightNode.parent.right = rightNode
		}
	} else {
		avl.root = rightNode
	}
	return rightNode
}

func rightRotation(node *Node, avl *AVL) *Node {
	var leftNode *Node = node.left

	node.left = leftNode.right
	leftNode.right = node
	leftNode.parent = node.parent
	node.parent = leftNode
	if node != avl.root {
		if leftNode.parent.left == node {
			leftNode.parent.left = leftNode
		} else {
			leftNode.parent.right = leftNode
		}
	} else {
		avl.root = leftNode
	}
	return leftNode
}

func rightLeftRotation(node *Node, avl *AVL) {
	newNode := rightRotation(node, avl)
	leftRotation(newNode, avl)
}

func leftRightRotation(node *Node, avl *AVL) {
	newNode := leftRotation(node, avl)
	rightRotation(newNode, avl)
}

func checkAndFixBalance(node *Node, avlTree *AVL) {
	if node == nil {
		return
	}
	if calcBalance(node) < -1 {
		if checkHeight(node.left.left) >= checkHeight(node.left.right) {
			rightRotation(node, avlTree)
		} else {
			rightLeftRotation(node, avlTree)
		}
	} else if calcBalance(node) > 1 {
		if checkHeight(node.right.right) >= checkHeight(node.right.left) {
			leftRotation(node, avlTree)
		} else {
			leftRightRotation(node, avlTree)
		}
	}
}

//tree operations

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

func insertOnTree(tree *AVL, key string, value int) {
	if tree.root == nil {
		tree.root = &Node{key: key, value: value, right: nil, left: nil, parent: nil}
	} else {
		insertNode(tree.root, key, value, tree)
	}
}

func insertNode(node *Node, key string, value int, avlTree *AVL) {
	if node == nil {
		return
	}
	if node.key > key {
		if node.left == nil {
			node.left = &Node{key: key, value: value, right: nil, left: nil, parent: node}
		} else {
			insertNode(node.left, key, value, avlTree)
		}
		//After new node already inserted(or we just can't go deeper this way)
		//lets check for unbalances
		checkAndFixBalance(node, avlTree)
	}
	if node.key <= key {
		if node.right == nil {
			node.right = &Node{key: key, value: value, right: nil, left: nil, parent: node}
		} else {
			insertNode(node.right, key, value, avlTree)
		}
		//After new node already inserted(or we just can't go deeper this way)
		//lets check for unbalances
		checkAndFixBalance(node, avlTree)
	}

}
