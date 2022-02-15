package main

import "fmt"


func f1(x *int){
	*x = 100
	fmt.Println(*x)
}

func f2(){
	var a *int
	fmt.Println(a)
	a = new(int)  // 申请一个内存地址
	fmt.Println(a,*a)

	*a = 100
	fmt.Println(*a)

	/*
	<nil>
	0x140000140a8
	100
	*/

}
func main(){
	s := "哈哈"  // 变量
	p := &s  // 取s变量的内存地址
	fmt.Println(p)  //s变量的内存地址
	fmt.Println(*p)   // 指针取值，根据指针取内存地址取值
	// p的类型
	fmt.Printf("p: %T\n",p)

	x := 1
	f1(&x)  // Go 语言中传只都是值拷贝
	fmt.Println(x)

	f2()
}