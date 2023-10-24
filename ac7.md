# Lista Duplamente Encadeada

Uma lista duplamente encadeada é uma estrutura de dados na qual cada nó contém um valor e dois ponteiros, um para o nó anterior na lista e outro para o próximo nó na lista, permitindo a navegação em ambas as direções.

## Estrutura do Nó

Cada nó na lista duplamente encadeada é definido pela seguinte estrutura:

go
type Node struct {
    Value    int      // Valor armazenado no nó
    Previous *Node    // Ponteiro para o nó anterior na lista
    Next     *Node    // Ponteiro para o próximo nó na lista
}



# Estrutura da Lista Dupla
A lista duplamente encadeada é mantida através da estrutura a seguir:

go
type DoublyLinkedList struct {
    Head    *Node  // Ponteiro para o primeiro nó da lista (cabeça)
    Tail    *Node  // Ponteiro para o último nó da lista (rabo)
    Size    int    // Tamanho da lista
}

## Operações Comuns

### Exibição dos Nós em uma Lista Duplamente Encadeada

Função para exibir os nós na lista duplamente encadeada:

go
func DisplayList(dll *DoublyLinkedList) {
    node := dll.Head
    for node != nil {
        fmt.Println(node.Value)
        node = node.Next
    }
    fmt.Println()
}


### Busca de um Nó em uma Lista Duplamente Encadeada

Função para buscar um nó com um valor específico na lista duplamente encadeada:

go
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

### Inserção de um Nó em uma Lista Duplamente Encadeada

Função para inserir um nó na lista duplamente encadeada:

go
func InsertInList(dll *DoublyLinkedList, key int) {
    newNode := &Node{Value: key}

    // Lógica de inserção com base no valor da chave
    // ...

    dll.Size++
}

### Remoção de um Nó em uma Lista Duplamente Encadeada

Função para remover um nó com um valor específico da lista duplamente encadeada:

go
func RemoveFromList(dll *DoublyLinkedList, key int) {
    // Lógica de remoção com base no valor da chave
    // ...

    dll.Size--
}

## Exemplo de Uso

O exemplo a seguir demonstra como criar uma lista duplamente encadeada, inserir elementos nela, exibir a lista e remover um elemento:

```go
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
