# Estrutura de uma Lista Circular (explicação)
Uma lista circular em Go é uma estrutura de dados na qual o último nó está conectado ao primeiro, 
criando um ciclo. Ela suporta três operações principais: exibir os nós, inserir um novo nó no início e 
excluir o primeiro nó. Antes de entrarmos nos detalhes das operações, vamos entender a estrutura básica de um nó e da lista.

# Estrutura de um Nó
Um nó em uma lista circular em Go possui dois componentes principais:

Valor (Data): O valor armazenado no nó.
Próximo (Next): Um ponteiro para o próximo nó na lista circular.
Estrutura da Lista Circular
Para representar uma lista circular em Go, usaremos uma estrutura que inclui:

Cabeça (Head): Um ponteiro para o nó que serve como referência ao início da lista circular.
Exibição dos Nós em uma Lista Circular em Go
A operação de exibição permite visualizar todos os nós na lista. Aqui está a função que realiza isso em Go:

(em go p ficar de melhor entendimento - pseudocodigo me confunde um pouco)

```
func ExibirListaCircular(listaCircular *No) {
    johnLennon := listaCircular
    for {
        fmt.Println(johnLennon.valor)
        noAtual = noAtual.proximo
        if johnLennon == listaCircular {
            break
        }
    }
}
```

Nesta função, começamos com o nó da cabeça e percorremos os nós em um loop até voltarmos à cabeça. 
Durante o percurso, imprimimos o valor de cada nó.

# Inserção de um Nó no Início da Lista
A operação de inserção permite adicionar um novo nó no início da lista. 
Aqui está a função revisada em Go:

```
func InserirNoInicio(listaCircular **No, novoValor int) {
    bonJovi := &No{valor: novoValor, proximo: *listaCircular}
    if *listaCircular == nil {
        bonJovi.proximo = bonJovi
    } else {
        atual := *listaCircular
        for atual.proximo != *listaCircular {
            atual = atual.proximo
        }
        atual.proximo = bonJovi
    }
    *listaCircular = bonJovi
}
```

Nesta função, criamos um novo nó com o valor fornecido. Se a lista estiver vazia (a cabeça é nula), definimos o próximo do novo nó como ele mesmo. 
Caso contrário, percorremos a lista até encontrar o último nó e ajustamos os ponteiros para incluir o novo nó no início.

# Exclusão do Primeiro 
A operação de exclusão remove o primeiro nó da lista. Aqui está a função revisada em Go:

```
func ExcluirPrimeiroNo(listaCircular **No) {
    if *listaCircular != nil {
        if (*listaCircular).proximo == *listaCircular {
            *listaCircular = nil
        } else {
            atual := *listaCircular
            for atual.proximo != *listaCircular {
                atual = atual.proximo
            }
            atual.proximo = (*listaCircular).proximo
            *listaCircular = (*listaCircular).proximo
        }
    }
}
```

Nesta função, verificamos se a lista não está vazia (a cabeça não é nula). Se houver apenas um nó na lista, definimos a cabeça como nula. 
Caso contrário, percorremos a lista para encontrar o último nó e ajustamos os ponteiros para remover o primeiro nó da lista.
