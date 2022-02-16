package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	/*
		只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
		结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
	*/
	var p1 person
	p1.name = "ljw"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={ljw 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"ljw", city:"北京", age:18}
	fmt.Println(p1.name)

	var p2 = new(person)
	fmt.Printf("%T\n", p2) //*main.person
	fmt.Printf("p2=%#v\n", p2)
	fmt.Println(*&p2.age)
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2)
	fmt.Println(*&p2.age)


	p3 := &person{} // 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p3.name = "haha"
	p3.age = 30
	p3.city = "广东"
	fmt.Printf("p3=%#v\n", p3) // p2=&main.person{name:"haha", city:"广东", age:30}

	//p4 := person{}
	var p4 person // //p4 := person{}
	fmt.Printf("p4=%#v\n", p4)


	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)  // p5=main.person{name:"小王子", city:"北京", age:18}

	// 也可以对结构体指针进行键值对初始化，例如：
	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6)  // p6=&main.person{name:"小王子", city:"北京", age:18}
}
