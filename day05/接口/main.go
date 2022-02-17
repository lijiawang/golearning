package main

import "fmt"

type Mover interface{
	Move()
}

type Dog struct{}

func (d *Dog)Move(){
	fmt.Println("四条腿跑")
}

func main(){
	var m Mover
	d := &Dog{}
	m = d
	m.Move()

}