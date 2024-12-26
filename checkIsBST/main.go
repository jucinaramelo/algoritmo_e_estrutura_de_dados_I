package main

import "fmt"

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

// Função IsBST que verifica se a árvore é uma BST
func (node *BstNode) IsBST() bool {
	left := true
	if node.left != nil {
		if node.value <= node.left.value {
			left = false
		} else {
			left = node.left.IsBST()
		}
	}

	right := true
	if node.right != nil {
		if node.value >= node.right.value {
			right = false
		} else {
			right = node.right.IsBST()
		}
	}

	return left && right
}

func main() {
	// Criando a árvore de exemplo
	root := &BstNode{value: 10}
	root.left = &BstNode{value: 5}
	root.right = &BstNode{value: 15}
	root.left.left = &BstNode{value: 3}
	root.left.right = &BstNode{value: 7}
	root.right.right = &BstNode{value: 20}

	// Teste com árvore válida
	fmt.Println("A árvore é BST?", root.IsBST()) 

}
