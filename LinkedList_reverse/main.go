package main

import "fmt"

type LinkedList struct {
        head *Node
        size int
}

type Node struct {
        value int
        next  *Node
}

func (list *LinkedList) Reverse() {
        if list.head == nil || list.head.next == nil {
                return // Lista vazia ou com apenas um elemento, não precisa inverter
        }

        var prev *Node
        curr := list.head
        var next *Node

        for curr != nil {
                next = curr.next // Guarda o próximo nó
                curr.next = prev  // Inverte o ponteiro: o próximo do nó atual aponta para o anterior
                prev = curr       // O anterior agora é o nó atual
                curr = next       // O nó atual passa a ser o próximo (que foi guardado)
        }

        list.head = prev // A cabeça da lista agora é o último nó (que antes era o primeiro)
}

func (list *LinkedList) Insert(value int) {
        newNode := &Node{value: value}
        if list.head == nil {
                list.head = newNode
        } else {
                current := list.head
                for current.next != nil {
                        current = current.next
                }
                current.next = newNode
        }
        list.size++
}

func (list *LinkedList) Print() {
        current := list.head
        for current != nil {
                fmt.Printf("%d -> ", current.value)
                current = current.next
        }
        fmt.Println("nil")
}

func main() {
        list := &LinkedList{}
        list.Insert(1)
        list.Insert(2)
        list.Insert(3)
        list.Insert(4)

        fmt.Println("Lista original:")
        list.Print()

        list.Reverse()

        fmt.Println("Lista invertida:")
        list.Print()

    list2 := &LinkedList{}
    list2.Insert(1)
    fmt.Println("Lista 2 original:")
    list2.Print()
    list2.Reverse()
    fmt.Println("Lista 2 invertida:")
    list2.Print()

    list3 := &LinkedList{}
    fmt.Println("Lista 3 original (vazia):")
    list3.Print()
    list3.Reverse()
    fmt.Println("Lista 3 invertida (vazia):")
    list3.Print()


}