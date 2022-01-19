package main

import "fmt"

type Pen struct {
	Name string
	age  int
}

func (p Pen) Sex(sex string) {
	fmt.Println(sex, p.age, p.Name)
}

func main() {
	var p1 Pen = Pen{
		Name: "ljw",
		age:  18,
	}
	p1.Sex("1")

}
