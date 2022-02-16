package main

/*
Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

方法的定义格式如下：
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
其中，

接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：具体格式与函数定义相同。

*/
import (
	"fmt"
)

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
//Dream Person做梦的方法
func (p Person) Dream(s string) {
	fmt.Printf("%s的梦想是%s\n", p.name,s)
}


// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}


// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
	fmt.Printf("当前函数的年龄%d\n",p.age)  // 只在本函数内有效
}
func main() {
	p1 := NewPerson("小王子", 25)
	fmt.Println(p1.age,p1.name)
	fmt.Printf("%v\n" , p1.Dream)  // 指针地址
	p1.Dream("学好go语言")

	p1.SetAge(30)   // 修改了内存地址
	fmt.Println(p1.age)

	p1.SetAge2(40)
	fmt.Println(p1.age)


}
