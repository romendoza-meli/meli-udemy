package main

import "go-bases/02-interfaces/animal"

func main() {

	myDog := animal.Dog{Name: "Sebas"}
	myCat := animal.Cat{Name: "Michi"}

	animal.DoingSound(&myDog)
	animal.DoingSound(&myCat)
}
