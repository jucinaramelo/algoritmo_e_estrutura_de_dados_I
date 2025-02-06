package main

import "fmt"

type Node2P struct {
	prev  *Node2P
	value int
	next  *Node2P
}

type DoublyLinkedList struct {
	head *Node2P
	tail *Node2P
	size int
}

func (list *DoublyLinkedList) Reverse() {
	current := list.head
	var temp *Node2P

	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	if temp != nil {
		list.head = temp.prev
	}
}

// Função auxiliar para adicionar elementos à lista
func (list *DoublyLinkedList) Append(value int) {
	newNode := &Node2P{value: value}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.size++
}

// Função auxiliar para imprimir a lista
func (list *DoublyLinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Print(current.value, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := &DoublyLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	fmt.Println("Lista original:")
	list.Print()

	list.Reverse()

	fmt.Println("Lista invertida:")
	list.Print()
}
