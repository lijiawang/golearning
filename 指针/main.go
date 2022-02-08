package main

import "fmt"


// func swap(a, b *int)  {
//     *a, *b = *b, *a
//     fmt.Println(*a, *b)
// }

func main() {
	// a := 100
	// var ptr *int // 声明指针类型
	// ptr = &a     // 初始化指针类型值为变量 a
	// fmt.Println(&a)  // 0x14000124008
	// fmt.Println(ptr)   // 0x140000140a8
	// fmt.Println(*ptr) // 100

	// 通过指针给函数传值
	// a := 1  // 定义一个int类型并复值
    // b := 2
	// fmt.Println(a, b)
    // swap(&a, &b)
    // fmt.Println(a, b)
	/*
	1 2
	2 1
	2 1
	*/
	a := 100
	var ptr *int
	fmt.Println(ptr)
	fmt.Println(&a)
	ptr = &a 
	fmt.Println(ptr)
	fmt.Println(*ptr)   

}
