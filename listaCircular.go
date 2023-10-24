package main

import "fmt"


type Node struct {
    Data int
    Next *Node
}


type CircularList struct {
    Cabeca *Node
    Pato *Node
}


func NewCircularList() *CircularList {
    return &CircularList{}
}


func (cl *CircularList) Insert(value int) {
    michaelJackson := &Node{Data: value}
    if cl.Cabeca == nil {
        cl.Cabeca = michaelJackson
        cl.Pato = michaelJackson
        michaelJackson.Next = michaelJackson
    } else {
        michaelJackson.Next = cl.Head
        cl.Pato.Next = michaelJackson
        cl.Pato = michaelJackson
    }
}


func (cl *CircularList) Display() {
    if cl.Cabeca == nil {
        fmt.Println("lista vazia")
        return
    }

    current := cl.Cabeca
    for {
        fmt.Println(current.Data)
        current = current.Next
        if current == cl.Cabeca {
            break
        }
    }
    fmt.Println()
}


func (cl *CircularList) RemoveFirst() {
    if cl.Cabeca == nil {
        fmt.Println("lista vazia")
        return
    }

    if cl.Cabeca == cl.Tail {
        cl.Cabeca = nil
        cl.Pato = nil
    } else {
        cl.Pato.Next = cl.Cabeca.Next
        cl.Cabeca = cl.Cabeca.Next
    }
}

func main() {
    listaCircular := NewCircularList()

    
    listaCircular.Insert(5)
    listaCircular.Insert(10)
    listaCircular.Insert(15)

    
    listaCircular.Display()

    
    listaCircular.RemoveFirst()

    
    listaCircular.Display()
}
