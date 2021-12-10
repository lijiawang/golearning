package main

import "fmt"

var name string
var a int64

func main() {
	a = 666
	name = "ljw"
	fmt.Println(name,a)
	s := "Hello 哈哈"
	for _,k := range s{
		fmt.Printf("%c\n",k)
	}
}
