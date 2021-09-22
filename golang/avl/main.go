package main

func main() {
	avlTree := createNewTree()
	insertOnTree(avlTree, "leandro", 22)
	insertOnTree(avlTree, "juliana", 48)
	insertOnTree(avlTree, "roberto", 64)
	insertOnTree(avlTree, "adam", 13)
	insertOnTree(avlTree, "john doe", 1)
	insertOnTree(avlTree, "eric", 35)
	insertOnTree(avlTree, "jonas", 45)
	insertOnTree(avlTree, "leandro", 98)
	insertOnTree(avlTree, "quincas", 11)

	toString(avlTree.root)

}
