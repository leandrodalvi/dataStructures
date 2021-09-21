package avl

type AVL struct {
	root *Node
}

func createNewTree() *AVL {
	newTree := new(AVL)
	newTree.root = nil
	return newTree
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

func insert(tree *AVL, key string, value interface{}) {
	if tree == nil {
		tree.root = createNewNode(key, value, nil, nil)
	}
	// else{

	// }
}

func checkHeight(root *Node) int {
	if root == nil {
		return 0
	} else {
		return 1 + checkHeight(root.left) + checkHeight(root.right)
	}
}
