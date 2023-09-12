package animal

import (
	"fmt"
)

type Animal interface{
	Sound()
}

type Dog struct{
	Name string
}

func (d *Dog) Sound(){
	fmt.Println(d.Name + " Woof woof!")
}
//Cat structure
type Cat struct{
	Name string
}

func (c *Cat) Sound(){
	fmt.Println(c.Name + " Meow meow!")
}

func DoingSound(animal Animal){
	animal.Sound()
}