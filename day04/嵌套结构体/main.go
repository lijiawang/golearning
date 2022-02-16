package main

import "fmt"
// //Address 地址结构体
// type Address struct {
// 	Province string  	// 省
// 	City     string   // 城市
// }

// //User 用户结构体
// type User struct {
// 	Name    string
// 	Gender  string  // 性别
// 	Age     int8
// 	// Address Address  // 嵌套结构图
// 	Address  // 匿名结构体
// }

// func main(){
// 	user := User{
// 		Name: "李家旺",
// 		Gender: "男",
// 		Age:   18,
// 		Address: Address{
// 			Province: "河北",
// 			City: "沧州",
// 		},
// 	}
// 	fmt.Printf("user=%#v\n",user)
// 	fmt.Println(user.Address.Province)
// 	fmt.Println(user)

// 	user.Province="HEBI"
// 	fmt.Println(user)
	
// }



//////继承
type Animal struct{
	name  string
}

func (a Animal)move(){
	fmt.Printf("%s 会动...\n",a.name)
}

type Dog struct{
	let int
	Animal   // 继承了Animal所有行为
}

func (d Dog)wang(){
	fmt.Println("汪汪汪!")
}

func dome01(){
	d := Dog{
		let: 4,
		Animal: Animal{
			name: "dome01",
		},
	}
	d.move()
	d.wang()
}
func main(){
	dome01()
}
