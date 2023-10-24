package main

import "fmt"

// Defina a estrutura do nó
type Node struct {
    Value    int
    Previous *Node
    Next     *Node
}

// Defina a estrutura da lista duplamente encadeada
type DoublyLinkedList struct {
    Head    *Node
    Tail    *Node
    Size    int
}

// Função para exibir a lista duplamente encadeada
func DisplayList(dll *DoublyLinkedList) {
    node := dll.Head
    for node != nil {
        fmt.Println(node.Value)
        node = node.Next
    }
    fmt.Println()
}

// Função para buscar um elemento na lista duplamente encadeada
func SearchList(dll *DoublyLinkedList, key int) *Node {
    node := dll.Head
    for node != nil {
        if node.Value == key {
            return node
        }
        node = node.Next
    }
    return nil
}

// Função para inserir um elemento na lista duplamente encadeada
func InsertInList(dll *DoublyLinkedList, key int) {
    newNode := &Node{Value: key}

    if dll.Head == nil {
        dll.Head = newNode
        dll.Tail = newNode
    } else if key <= dll.Head.Value {
        newNode.Next = dll.Head
        dll.Head.Previous = newNode
        dll.Head = newNode
    } else if key >= dll.Tail.Value {
        newNode.Previous = dll.Tail
        dll.Tail.Next = newNode
        dll.Tail = newNode
    } else {
        node := dll.Head
        for node.Next != nil && node.Next.Value < key {
            node = node.Next
        }
        newNode.Next = node.Next
        newNode.Previous = node
        node.Next.Previous = newNode
        node.Next = newNode
    }
    dll.Size++
}

// Função para remover um elemento da lista duplamente encadeada
func RemoveFromList(dll *DoublyLinkedList, key int) {
    nodeToRemove := SearchList(dll, key)
    if nodeToRemove != nil {
        if nodeToRemove.Previous != nil {
            nodeToRemove.Previous.Next = nodeToRemove.Next
        } else {
            dll.Head = nodeToRemove.Next
        }
        if nodeToRemove.Next != nil {
            nodeToRemove.Next.Previous = nodeToRemove.Previous
        } else {
            dll.Tail = nodeToRemove.Previous
        }
        dll.Size--
    }
}

func main() {
    doublyLinkedList := DoublyLinkedList{}

    // Insira elementos na lista duplamente encadeada
    InsertInList(&doublyLinkedList, 3)
    InsertInList(&doublyLinkedList, 1)
    InsertInList(&doublyLinkedList, 2)
    InsertInList(&doublyLinkedList, 4)

    // Exiba a lista duplamente encadeada
    DisplayList(&doublyLinkedList)

    // Remova um elemento da lista
    RemoveFromList(&doublyLinkedList, 2)

    // Exiba a lista novamente após a remoção
    DisplayList(&doublyLinkedList)
}
