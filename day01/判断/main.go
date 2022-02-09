package main

import "fmt"

func main(){
	score := 89
	if score > 90 {
		fmt.Println("优")
	} else if score == 89 {
		fmt.Println("B")		
	} else {
		fmt.Println("C")
	}
}