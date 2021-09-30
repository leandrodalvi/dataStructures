package main

func regularPopulate(avlTree *AVL) {
	insertOnTree(avlTree, "leandro", 22)
	insertOnTree(avlTree, "juliana", 48)
	insertOnTree(avlTree, "roberto", 64)
	insertOnTree(avlTree, "adam", 13)
	insertOnTree(avlTree, "john doe", 1)
	insertOnTree(avlTree, "eric", 35)
	insertOnTree(avlTree, "jonas", 45)
	insertOnTree(avlTree, "leandro", 98)
	insertOnTree(avlTree, "quincas", 11)
	insertOnTree(avlTree, "molango", 97)
	insertOnTree(avlTree, "apolo", 42)
	insertOnTree(avlTree, "lucas", 54)
	insertOnTree(avlTree, "rogerinho", 17)
	insertOnTree(avlTree, "evanildo", 26)
	insertOnTree(avlTree, "carol", 89)
	insertOnTree(avlTree, "pedro", 33)
	insertOnTree(avlTree, "winnie", 91)
	insertOnTree(avlTree, "giovanna", 2)
}

func main() {
	avlTree := createNewTree()
	regularPopulate(avlTree)

	println("Printing the tree...")
	toString(avlTree.root)

	println("Starting Benchmark checks on avl tree...\n")
	var avlTreeBalanceCoef int = calcBalance(avlTree.root)
	if avlTreeBalanceCoef <= 1 && avlTreeBalanceCoef >= -1 {
		println("The avl balance coeficient is ", avlTreeBalanceCoef, "\nThe status of the avl tree is balanced\n\n")
	} else {
		print("The avl balance coeficient is ", avlTreeBalanceCoef, "\nThis tree is unbalanced. Shame on you, Leo\n\n")
	}

	print("The tree height is ", checkHeight(avlTree.root), "\n\n")

}
