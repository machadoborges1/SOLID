package main

func main() {
	sparrow := Sparrow{Bird{Name: "Sparrow"}}
	sparrow.Fly() // Funciona como esperado

	penguin := Penguin{Bird{Name: "Penguin"}}
	penguin.Swim() // Agora, o pinguim faz algo que faz sentido
}
