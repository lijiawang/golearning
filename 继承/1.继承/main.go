package main

import "fmt"

type Animal struct{
	Name string
	Sex string
}

func (a *Animal) Talk() {
	fmt.Printf("I talk is %s\n",a.Name)
}

type Dog struct{
	Feet string
	*Animal
}

func (d *Dog) Eat() {
	fmt.Printf("dog is eat\n")
}

func main(){
	var d *Dog = &Dog{
		Feet: "four feet",
		Animal: &Animal{
			Name: "dog",
			Sex: "1",
		},
		
	}
	// d.Name = "dog"
	// d.Sex = "1"
	d.Eat()
	fmt.Println(d.Name,d.Sex)
	d.Talk()
}