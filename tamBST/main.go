package main

import "fmt"

type BstNode struct {
    left  *BstNode
    value int
    right *BstNode
}

// Função Size que retorna o tamanho da árvore
func (node *BstNode) Size() int {
    tam := 1
	if node.left != nil{
		tam+= node.left.Size()
	}
	if node.right != nil{
		tam += node.right.Size()
	}
	return tam
}

// Função para criar um novo nó
func NewNode(value int) *BstNode {
    return &BstNode{value: value}
}

func main() {
    // Criando a árvore de exemplo
    root := NewNode(10)
    root.left = NewNode(5)
    root.right = NewNode(15)
    root.left.left = NewNode(3)
    root.left.right = NewNode(7)

    // Calculando o tamanho da árvore
    fmt.Println("Tamanho da árvore:", root.Size()) // Esperado: 5
}
