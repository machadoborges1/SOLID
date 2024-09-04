package main

func main() {
	sparrow := Sparrow{Bird{Name: "Sparrow"}}
	sparrow.Fly()

	penguin := Penguin{Bird{Name: "Penguin"}}
	penguin.Swim()
}
