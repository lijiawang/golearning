package main

import (
	"fmt"
	"time"
)

// // defer 语句
// func f1() int {
// 	a := 1000
// 	fmt.Println("start")

// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// 	defer fmt.Println("3")


// 	fmt.Println("end")
// 	return a
// }

// func main() {
// 	a := f1()
// 	fmt.Println(a)
// }

// func calc(index string, a, b int) int {
// 	ret := a + b
// 	fmt.Println(index, a, b, ret)
// 	return ret
// }

func main() {
	// x := 1
	// y := 2
	// defer calc("AA", x, calc("A", x, y))
	// x = 10
	// defer calc("BB", x, calc("B", x, y))
	// y = 20

	now := time.Now() //获取当前时间
	fmt.Println( now)
	fmt.Println(now.Unix())
	fmt.Println(time.Hour)
}