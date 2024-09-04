// Temos uma classe base Bird e duas subclasses: Sparrow (um pássaro que voa) e Penguin (um pássaro que não voa).
// A função Fly() é definida na classe base, mas não faz sentido para a classe Penguin.

package main

import "fmt"

// Classe base Bird
type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Printf("%s is flying!\n", b.Name)
}

// Subclasse Sparrow (pardal)
type Sparrow struct {
	Bird
}

// Subclasse Penguin (pinguim)
type Penguin struct {
	Bird
}

func main() {
	sparrow := Sparrow{Bird{Name: "Sparrow"}}
	sparrow.Fly() // Funciona como esperado

	penguin := Penguin{Bird{Name: "Penguin"}}
	penguin.Fly() // Não faz sentido, pois pinguins não voam
}

// O método Fly() não deveria ser chamado para um Penguin, mas a classe Penguin herda esse método da classe base Bird.
// Isso viola o LSP, porque a substituição da classe base por uma subclasse (no caso, Penguin) leva a um comportamento incorreto ou inesperado.
