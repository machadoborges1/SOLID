package main

import "fmt"

// Classe base Bird
type Bird struct {
	Name string
}

// Interface Flyer para pássaros que voam
type Flyer interface {
	Fly()
}

// Interface Swimmer para pássaros que nadam
type Swimmer interface {
	Swim()
}

// Subclasse Sparrow (pardal) que implementa Flyer
type Sparrow struct {
	Bird
}

func (s Sparrow) Fly() {
	fmt.Printf("%s is flying!\n", s.Name)
}

// Subclasse Penguin (pinguim) que implementa Swimmer
type Penguin struct {
	Bird
}

func (p Penguin) Swim() {
	fmt.Printf("%s is swimming!\n", p.Name)
}
